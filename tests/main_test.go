package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetProducts(t *testing.T) {
	r := SetupRouter()
	r.GET("/products", GetProducts)
	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var products []Product
	json.Unmarshal([]byte(w.Body.Bytes()), &products)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 1, len(products))
}

func TestAddProduct(t *testing.T) {
	r := SetupRouter()
	r.POST("/product", AddProduct)

	product := Product{
		Name: "Rugged Jeans",
	}
	jsonValue, _ := json.Marshal(product)
	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateProduct(t *testing.T) {
	r := SetupRouter()
	r.PUT("/product/:id", UpdateProduct)

	product := Product{
		ID:   "c0283p3d0cvuglq85log",
		Name: "Red Tie",
		Description: "For all formal occasions",
	}
	jsonValue, _ := json.Marshal(product)
	reqFound, _ := http.NewRequest("PUT", "/product/"+product.ID, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)

	assert.Equal(t, http.StatusOK, w.Code)

	reqNotFound, _ := http.NewRequest("PUT", "/product/1", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqNotFound)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
