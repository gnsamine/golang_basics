package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() {
	response, err := http.Get("http://localhost:3000/products")

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	//to read response
	bodyBytes, _ := io.ReadAll(response.Body)

	var products []Product
	json.Unmarshal(bodyBytes, &products)
	fmt.Println(products)
}
func AddProduct() {
	product := Product{Id: 4, ProductName: "Telephone", CategoryId: 1, UnitPrice: 5000.0}
	jsonProduct, _ := json.Marshal(product)

	response, err := http.Post("http://localhost:3000/products", "application/jason;charset=utf-8", bytes.NewBuffer(jsonProduct))

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

}
