package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/go-redis/redis"
	"github.com/rabbice/ecommerce/src/backend/cache"
	"github.com/rabbice/ecommerce/src/backend/database"
	"github.com/rabbice/ecommerce/src/backend/helpers"
	"github.com/rabbice/ecommerce/src/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = database.OpenCollection(database.Client, "product")
var validate = validator.New()
var red = cache.GetRedisConnection()

func AddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "SELLER"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}

		shop_id := c.Query("id")
		if shop_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "shop id is empty"})
			c.Abort()
			return
		}

		product, err := primitive.ObjectIDFromHex(shop_id)
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}

		var products models.Product
		products.ID = primitive.NewObjectID()
		products.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		products.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if err := c.BindJSON(&products); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		matchStage := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: product}}}}
		unwindStage := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$product"}}}}

		cursor, err := shopCollection.Aggregate(ctx, mongo.Pipeline{matchStage, unwindStage})
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}

		var productInfo []bson.M
		if err = cursor.All(ctx, &productInfo); err != nil {
			panic(err)
		}

		filter := bson.D{primitive.E{Key: "_id", Value: product}}
		update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "products", Value: products}}}}
		_, err = shopCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			fmt.Println(err)
		}
		log.Println("Remove data from Redis")
		red.Del("product")
		c.IndentedJSON(200, gin.H{"message": "Product successfully added"})
	}
}

func GetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		val, err := red.Get("product").Result()
		if err == redis.Nil {
			log.Printf("Request to MongoDB")
			cursor, err := productCollection.Find(ctx, bson.M{})

			if err != nil {
				c.IndentedJSON(500, "Internal Server Error")
			}
			defer cursor.Close(ctx)

			products := make([]models.Product, 0)
			for cursor.Next(ctx) {
				var product models.Product
				cursor.Decode(&product)
				products = append(products, product)
			}

			data, _ := json.Marshal(products)
			red.Set("product", string(data), 0)
			c.JSON(http.StatusOK, products)
		} else if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		} else {
			log.Printf("Request to Redis")
			products := make([]models.Product, 0)
			json.Unmarshal([]byte(val), &products)

			c.JSON(http.StatusOK, products)

		}
		defer cancel()

	}
}

func GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		productId := c.Param("product_id")
		var product models.Product
		objectId, _ := primitive.ObjectIDFromHex(productId)
		filter := bson.M{"_id": objectId}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		err := productCollection.FindOne(ctx, filter).Decode(&product)
		defer cancel()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": "error while fetching the product"})
			return
		}
		c.IndentedJSON(http.StatusOK, product)
	}
}

func DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "SELLER"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		productId := c.Param("product_id")
		objectId, _ := primitive.ObjectIDFromHex(productId)
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		_, err := productCollection.DeleteOne(ctx, bson.M{"_id": objectId})
		defer cancel()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "Product successfully deleted"})
	}
}

func UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "SELLER"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		productId := c.Param("product_id")
		var product models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		objectId, _ := primitive.ObjectIDFromHex(productId)
		ctx := context.Background()
		_, err := productCollection.UpdateOne(ctx, bson.D{primitive.E{
			Key: "_id", Value: objectId}},
			bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "name", Value: product.Name},
				primitive.E{Key: "description", Value: product.Description},
				primitive.E{Key: "price", Value: product.Price},
				primitive.E{Key: "image", Value: product.Image},
			}}})
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error()})
			return
		}
		log.Println("Remove data from Redis")
		red.Del("product")
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "Product successfully updated"})
	}
}

func SearchProductByQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		var searchProducts []models.Product
		query := c.Query("tag")
		if query == "" {
			log.Println("Query is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid search params"})
			c.Abort()
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		search, err := productCollection.Find(ctx, bson.M{"tags": bson.M{"$regex": query}})
		if err != nil {
			c.IndentedJSON(404, "Cannot fetch the data")
			return
		}
		err = search.All(ctx, &searchProducts)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(400, "invalid")
			return
		}
		defer search.Close(ctx)

		if err := search.Err(); err != nil {
			log.Println(err)
			c.IndentedJSON(400, "invalid request")
			return
		}

		defer cancel()
		c.IndentedJSON(200, searchProducts)

	}
}
