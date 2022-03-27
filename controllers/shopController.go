package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rabbice/ecommerce/database"
	"github.com/rabbice/ecommerce/helpers"
	"github.com/rabbice/ecommerce/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var shopCollection *mongo.Collection = database.OpenCollection(database.Client, "shop")
var ctx = context.Background()

func AddShop() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "SELLER"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		var shop models.Shop

		if err := c.BindJSON(&shop); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(shop)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		shop.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		shop.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		shop.ID = primitive.NewObjectID()

		result, insertErr := shopCollection.InsertOne(ctx, shop)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Shop was not created"})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func GetShop() gin.HandlerFunc {
	return func(c *gin.Context) {
		shopId := c.Param("shop_id")
		var shop models.Shop

		err := shopCollection.FindOne(ctx, bson.M{"shop_id": shopId}).Decode(&shop)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the shop"})
		}

		c.JSON(http.StatusOK, shop)
	}
}

func DeleteShop() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "SELLER"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		shopId := c.Param("shop_id")
		objectId, _ := primitive.ObjectIDFromHex(shopId)
		_, err := shopCollection.DeleteOne(ctx, bson.M{"_id": objectId})
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "Shop successfully deleted"})
	}
}

func GetShops() gin.HandlerFunc {
	return func(c *gin.Context) {
		var shops []bson.M
		cur, err := shopCollection.Find(ctx, bson.M{})
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error()})
		}
		if err = cur.All(ctx, &shops); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		c.IndentedJSON(http.StatusOK, shops)
	}
}

func UpdateShop() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "SELLER"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		shopId := c.Param("shop_id")
		var shop models.Shop
		if err := c.ShouldBindJSON(&shop); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		objectId, _ := primitive.ObjectIDFromHex(shopId)
		_, err := shopCollection.UpdateOne(ctx, bson.D{primitive.E{
			Key: "_id", Value: objectId}},
			bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "name", Value: shop.Name},
				primitive.E{Key: "description", Value: shop.Description},
			}}})
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "Shop successfully updated"})
	}
}
