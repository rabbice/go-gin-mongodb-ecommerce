package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rabbice/ecommerce/src/backend/database"
	"github.com/rabbice/ecommerce/src/backend/helpers"
	"github.com/rabbice/ecommerce/src/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = database.OpenCollection(database.Client, "product")
var validate = validator.New()

func AddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "SELLER"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var product models.Product

		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		defer cancel()

		validationErr := validate.Struct(product)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		}
		product.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		product.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		product.ID = primitive.NewObjectID()

		_, insertErr := productCollection.InsertOne(ctx, product)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Product was not created"})
			fmt.Println(insertErr)
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, product)
	}
}

func GetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var products []bson.M

		cursor, err := productCollection.Find(ctx, bson.M{})

		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}
		defer cancel()
		if err = cursor.All(ctx, &products); err != nil {
			c.IndentedJSON(500, "Internal Server Error")
			return
		}
		c.IndentedJSON(200, products)
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
				primitive.E{Key:"price", Value: product.Price},
				primitive.E{Key:"image", Value: product.Image},
			}}})
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error()})
			return
		}
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
