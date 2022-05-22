package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

var products []Product

type Product struct {
	ID          string    `bson:"_id"`
	Name        string    `json:"name" validate:"required"`
	Description string    `bson:"description" json:"description"`
	Price       float64   `bson:"price" json:"price" validate:"required"`
	Image       string    `bson:"image" json:"image"`
	Tags        []string  `bson:"tags" json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
}

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func AddProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	product.ID = xid.New().String()
	product.CreatedAt = time.Now()

	products = append(products, product)

	c.IndentedJSON(200, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(id)

	index := -1
	for i := 0; i < len(products); i++ {
		if products[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	products[index] = product

	c.JSON(http.StatusOK, product)
}

func init() {
	products = make([]Product, 0)
	file, _ := ioutil.ReadFile("products.json")
	_ = json.Unmarshal([]byte(file), &products)
}

func main() {
	router := gin.Default()
	router.GET("/products", GetProducts)
	router.POST("/product", AddProduct)
	router.PUT("/product/:id", UpdateProduct)
	router.Run()
}
