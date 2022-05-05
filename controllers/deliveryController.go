package controllers

import (
	"context"
	"errors"
	"fmt"
	"log"
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

var deliveryCollection *mongo.Collection = database.OpenCollection(database.Client, "orders")

func CreateDelivery() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, true); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		sellerid := c.Query("id")
		if sellerid == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "seller id is empty"})
			c.Abort()
			return
		}
		orderid := c.Query("orderID")
		if orderid == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "order id is empty"})
			c.Abort()
			return
		}
		delivery, err := primitive.ObjectIDFromHex(sellerid)
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}
		var deliveries models.Delivery
		deliveries.ID = primitive.NewObjectID()
		if err = c.BindJSON(&deliveries); err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		matchStage := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: delivery}}}}
		unwindStage := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$deliveries"}}}}
		groupStage := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$delivery_id"}}}}
		cursor, err := userCollection.Aggregate(ctx, mongo.Pipeline{matchStage, unwindStage, groupStage})
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}

		var deliveryInfo []bson.M
		if err = cursor.All(ctx, &deliveryInfo); err != nil {
			panic(err)
		}
		filter := bson.D{primitive.E{Key: "_id", Value: delivery}}
		update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "deliveries", Value: deliveries}}}}
		_, err = userCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			fmt.Println(err)
		}
		_, insertErr := deliveryCollection.InsertOne(ctx, deliveries)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Delivery was not created"})
			return
		}
		c.IndentedJSON(200, gin.H{"message": "Order successfully shipped"})

	}

}

func ConfirmDelivery() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, false); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		deliveryQueryID := c.Query("id")
		if deliveryQueryID == "" {
			log.Println("delivery id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("delivery id is empty"))
			return
		}
		var delivery models.Delivery
		if err := c.ShouldBindJSON(&delivery); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		deliveryID, err := primitive.ObjectIDFromHex(deliveryQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key: "_id", Value: deliveryID}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "status", Value: delivery.Status}}}}
		_, err = deliveryCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
			return
		}
		defer cancel()
		ctx.Done()
		c.IndentedJSON(200, "Product successfully delivered")
	}
}
