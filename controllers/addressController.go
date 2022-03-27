package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rabbice/ecommerce/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "user id is empty"})
			c.Abort()
			return
		}
		address, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}
		var addresses models.Address
		addresses.ID = primitive.NewObjectID()
		if err = c.BindJSON(&addresses); err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		matchStage := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: address}}}}
		unwindStage := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$address"}}}}
		groupStage := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$address_id"}, {Key: "count", Value: bson.D{primitive.E{Key: "$sum", Value: 1}}}}}}

		cursor, err := userCollection.Aggregate(ctx, mongo.Pipeline{matchStage, unwindStage, groupStage})
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}

		var addressInfo []bson.M
		if err = cursor.All(ctx, &addressInfo); err != nil {
			panic(err)
		}

		var size int32
		for _, address_no := range addressInfo {
			count := address_no["count"]
			size = count.(int32)
		}
		if size < 2 {
			filter := bson.D{primitive.E{Key: "_id", Value: address}}
			update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}
			_, err := userCollection.UpdateOne(ctx, filter, update)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			c.IndentedJSON(400, "not allowed")
		}

		c.IndentedJSON(200, address)

	}
}

func EditAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		var address models.Address
		user_id := c.Query("id")
		if user_id == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid search index"})
			c.Abort()
			return
		}
		uaid, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key: "_id", Value: uaid}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "house_number", Value: address.HouseNumber}}}}
		_, err = userCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
			return
		}
		defer cancel()
		ctx.Done()
		c.IndentedJSON(200, "Successfully updated address")
	}
}

func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")
		if user_id == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid search index"})
			c.Abort()
			return
		}

		addresses := make([]models.Address, 0)
		uaid, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key: "_id", Value: uaid}}
		update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}
		_, err = userCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(400, "Bad request")
			return
		}
		ctx.Done()
		c.IndentedJSON(200, "Successfully deleted address")
	}

}
