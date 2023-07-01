package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Todo struct {
	//make pair  json and your code
	//apideki neyi maplediğimizi görürüz
	UserId    int    `json:"userId" `
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Demo1() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	//we can think http package like a browser

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	//we should use defer at this point
	//close when the process finish
	bodyBytes, _ := io.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	//string convertion
	fmt.Println(bodyString)

	var todo Todo
	json.Unmarshal(bodyBytes, &todo)
	//to convert json to string
	//assign bodyBytes to todo
	fmt.Println(todo)

}

func Demo2() {
	todo := Todo{1, 2, "go to shopping", false}
	//converting todo to json in order to put it in http.Post
	jsonTodo, err := json.Marshal(todo)
	// todo to json

	if err != nil {
		fmt.Println(err)
	}
	response, err := http.Post("https://jasonplaceholder.typicode.com/todos", "application/jason;charset=utf-8", bytes.NewBuffer(jsonTodo))

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	//string convertion
	fmt.Println(bodyString)

	var todoResponse Todo
	json.Unmarshal(bodyBytes, &todoResponse)
	//to convert json to string
	//assign bodyBytes to todo
	fmt.Println(todoResponse)

}
