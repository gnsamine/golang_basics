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
	//This line reads the content of the response body into the bodyBytes variable using io.ReadAll().
	//It reads all the data until the end of the response body or an error occurs.
	//The underscore _ is used to discard the error value returned by io.ReadAll()
	//since we are not handling errors in this specific case.

	var products []Product
	json.Unmarshal(bodyBytes, &products)
	fmt.Println(products)
}
func AddProduct() {
	product := Product{Id: 4, ProductName: "Telephone", CategoryId: 1, UnitPrice: 5000.0}
	jsonProduct, _ := json.Marshal(product) //convert to json

	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))
	//in this code is responsible for making an HTTP POST request to the server at the specified URL with the provided data.
	//The first argument is the URL to which the POST request will be sent.
	//The second argument is the content type of the data being sent in the request body. It specifies that the data is in
	//JSON format with the character encoding set to UTF-8.
	//The bytes.Buffer implements the io.Reader interface, allowing you to pass it as the request body.

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)

	var productResponse Product

	json.Unmarshal(bodyBytes, &productResponse)
	//assing bodybytes to productResponse
	fmt.Println("added", productResponse)

	//if you write 4 your Id it cannot be saved bc there is already
	//product that has Id:4

}

// refactoring
func GetAllProducts2() ([]Product, error) {
	response, err := http.Get("http://localhost:3000/products")

	if err != nil {
		return nil, err
	}
	//if there is an error, products will be empty and errror will be written

	defer response.Body.Close()
	//to read response
	bodyBytes, _ := io.ReadAll(response.Body)

	var products []Product
	json.Unmarshal(bodyBytes, &products)

	return products, nil
	//if there is not an error, error will be empty and product will be written
}
func AddProduct2() (Product, error) {
	product := Product{ProductName: "microphone", CategoryId: 1, UnitPrice: 5000.0}
	jsonProduct, _ := json.Marshal(product)

	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))
	//in this code is responsible for making an HTTP POST request to the server at the specified URL with the provided data.
	//The first argument is the URL to which the POST request will be sent.
	//The second argument is the content type of the data being sent in the request body. It specifies that the data is in
	//JSON format with the character encoding set to UTF-8.
	//The bytes.Buffer implements the io.Reader interface, allowing you to pass it as the request body.

	if err != nil {
		return Product{}, err
	}
	//we cannnot return thing that have not pointer value as a nil
	//so we write return Product{}, err

	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)

	var productResponse Product

	json.Unmarshal(bodyBytes, &productResponse)
	//assing bodybytes to productResponse
	return productResponse, nil

	//if you write 4 your Id it cannot be saved bc there is already
	//product that has Id:4

}
