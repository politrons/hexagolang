package main

import (
	"app/handler"
	. "app/service"
	"encoding/json"
	"fmt"
	"infra"
	"log"
	"net/http"
)

var orderDAO infra.OrderDAO = infra.OrderDAOImpl{}
var orderService OrderService = OrderServiceImpl{OrderDAO: orderDAO}
var orderHandler handler.OrderHandler = handler.OrderHandlerImpl{OrderDAO: orderDAO}

func main() {
	port := "1981"
	log.Printf("Running hexagonal server on port:%s......", port)
	server := http.NewServeMux()
	server.HandleFunc("/order/", findOrderHandle)
	server.HandleFunc("/order/create/", createOrderHandle)
	server.HandleFunc("/order/add/", addProductHandle)
	server.HandleFunc("/order/remove/", removeProductHandle)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), server))
}

func findOrderHandle(writer http.ResponseWriter, request *http.Request) {
	order := orderService.GetOrder(198)
	response, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}
	renderResponse(writer, response)
}

func createOrderHandle(writer http.ResponseWriter, request *http.Request) {
	response := "Create"
	renderResponse(writer, []byte(response))
}

func removeProductHandle(writer http.ResponseWriter, request *http.Request) {
	response := "Remove"
	renderResponse(writer, []byte(response))
}

func addProductHandle(writer http.ResponseWriter, request *http.Request) {
	response := "Add"
	renderResponse(writer, []byte(response))
}

func renderResponse(writer http.ResponseWriter, response []byte) {
	code, error := writer.Write(response)
	if error != nil {
		log.Println("Error rendering response. Caused by ")
	} else {
		log.Printf("Success in response with code %d", code)
	}
}
