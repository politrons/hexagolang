package main

import (
	"app/command"
	"app/handler"
	. "app/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"infra"
	"log"
	"net/http"
	"strings"
)

var orderDAO infra.OrderDAO = infra.OrderDAOImpl{}
var orderService OrderService = OrderServiceImpl{OrderDAO: orderDAO}
var orderHandler handler.OrderHandler = handler.OrderHandlerImpl{OrderDAO: orderDAO}

func main() {
	port := "1981"
	log.Printf("Running hexagonal server on port:%s......", port)
	server := http.NewServeMux()
	server.HandleFunc("/order/", createOrderId)
	server.HandleFunc("/order/find/{id}", findOrder)
	server.HandleFunc("/order/create/", createOrder)
	server.HandleFunc("/product/{id}/", findProduct)
	server.HandleFunc("/product/add/", addProduct)
	server.HandleFunc("/product/remove/", removeProduct)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), server))
}

/**
In order to have an idempotent endpoint in creation of Order, we generate a orderId
and the client must use this id in the another endpoint to perform the creation of the Order
*/
func createOrderId(writer http.ResponseWriter, request *http.Request) {
	orderId := []byte(uuid.New().String())
	renderResponse(writer, orderId)
}

/**
In order to have an idempotent endpoint,We check if the Order already exist,
if it does not, we create one, otherwise we just return the OrderId.
*/
func createOrder(writer http.ResponseWriter, request *http.Request) {
	orderId := getOrderId(request)
	exist, _ := orderService.GetOrder(orderId)
	if !exist {
		orderHandler.CreateOrder(command.CreateOrderCommand{Id: orderId})
	}
	renderResponse(writer, []byte(orderId))
}

/**
Using the orderId, we're able to rehydrate the Order model from all the events persisted,
so using event sourcing, we can recreate the current status of the Order
*/
func findOrder(writer http.ResponseWriter, request *http.Request) {
	orderId := getOrderId(request)
	_, order := orderService.GetOrder(orderId)
	response, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}
	renderResponse(writer, response)
}

/**
Having a productId we can search for a product to obtain information, and then we create a transactionId for the product
in order to be idempotent when client want to add a product in the basket
*/
func findProduct(writer http.ResponseWriter, request *http.Request) {
	productId := []byte(uuid.New().String())

	response := "Remove"
	renderResponse(writer, []byte(response))
}

func removeProduct(writer http.ResponseWriter, request *http.Request) {
	response := "Remove"
	renderResponse(writer, []byte(response))
}

func addProduct(writer http.ResponseWriter, request *http.Request) {
	response := "Add"
	renderResponse(writer, []byte(response))
}

func renderResponse(writer http.ResponseWriter, response []byte) {
	code, err := writer.Write(response)
	if err != nil {
		log.Println("Error rendering response. Caused by ")
	} else {
		log.Printf("Success in response with code %d", code)
	}
}

func getOrderId(request *http.Request) string {
	return strings.Split(request.URL.Path, "/")[2]
}
