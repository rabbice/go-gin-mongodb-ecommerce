package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rabbice/ecommerce/database"
	"github.com/rabbice/ecommerce/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = database.OpenCollection(database.Client, "product")
var validate = validator.New()

func AddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var shop models.Shop
		var product models.Product

		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		validationErr := validate.Struct(product)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		}
		err := shopCollection.FindOne(ctx, bson.M{"shop_id": product.Shop_ID}).Decode(&shop)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Shop was not found"})
			return
		}
		product.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		product.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		product.ID = primitive.NewObjectID()
		product.Product_ID = product.ID.Hex()

		_, insertErr := productCollection.InsertOne(ctx, product)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Product was not created"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, product)
	}
}

func GetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}

		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}

		startIndex := (page - 1) * recordPerPage
		startIndex, _ = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{primitive.E{Key: "$match", Value: bson.D{{}}}}
		groupStage := bson.D{primitive.E{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: bson.D{primitive.E{Key: "_id", Value: "null"}}}, {Key: "total_count", Value: bson.D{primitive.E{Key: "$sum", Value: 1}}}, {Key: "data", Value: bson.D{primitive.E{Key: "$push", Value: "$$ROOT"}}}}}}
		projectStage := bson.D{
			primitive.E{
				Key: "$project", Value: bson.D{
					{Key: "_id", Value: 0},
					{Key: "total_count", Value: 1},
					{Key: "product_items", Value: bson.D{primitive.E{Key: "$slice", Value: []interface{}{"$data", startIndex, recordPerPage}}}},
				}}}

		result, err := productCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing food items"})
		}
		var allProducts []bson.M
		if err = result.All(ctx, &allProducts); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allProducts[0])
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
