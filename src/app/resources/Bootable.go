package main

import (
	"app/command"
	"app/handler"
	"app/render"
	. "app/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"infra/dao"
	"infra/response"
	"log"
	"net/http"
	"strings"
	"time"
)

var orderDAO dao.OrderDAO = dao.OrderDAOImpl{}
var orderService OrderService = OrderServiceImpl{OrderDAO: orderDAO}
var orderHandler handler.OrderHandler = handler.OrderHandlerImpl{OrderDAO: orderDAO}
var productDAO dao.ProductDAO = dao.ProductDAOImpl{}
var productService ProductService = ProductServiceImpl{ProductDAO: productDAO}
var productHandler handler.ProductHandler = handler.ProductHandlerImpl{OrderDAO: orderDAO}

func main() {
	port := "1981"
	log.Printf("Running hexagonal server on port:%s......", port)
	server := http.NewServeMux()
	server.HandleFunc("/order/", createOrderId)
	server.HandleFunc("/order/create/", createOrder)
	server.HandleFunc("/order/find/", findOrder)
	server.HandleFunc("/product/", findProduct)
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
	orderId := getArgumentAtIndex(request, 3)
	awaitOrderResponseChannel(orderService.GetOrder(orderId), func(orderResponse response.OrderResponse) {
		if !orderResponse.Exist {
			log.Printf("Creating order for transaction Id t:%s......", orderId)
			orderHandler.CreateOrder(command.CreateOrderCommand{Id: orderId})
		}
		renderResponse(writer, []byte(orderId))
	})
}

/**
Using the orderId, we're able to rehydrate the Order model from all the events persisted,
so using event sourcing, we can recreate the current status of the Order
*/
func findOrder(writer http.ResponseWriter, request *http.Request) {
	orderId := getArgumentAtIndex(request, 3)
	awaitOrderResponseChannel(orderService.GetOrder(orderId), func(orderResponse response.OrderResponse) {
		jsonResponse, err := json.Marshal(orderResponse.Order)
		if err != nil {
			panic(err)
		}
		renderResponse(writer, jsonResponse)
	})

}

/**
Function that receive a channel and a function to apply with the result of the channel once it ends.
We keep in a loop waiting for that channel to ends. In case ir does not end in 50ms we break the loop
and we consider the request wrong
*/
func awaitOrderResponseChannel(channel chan response.OrderResponse,
	action func(orderResponse response.OrderResponse)) {
loop:
	for {
		select {
		case orderResponse := <-channel:
			action(orderResponse)
		case <-time.After(50 * time.Millisecond):
			break loop
		}
	}
}

/**
Having a productId we can search for a product to obtain information, and then we create a transactionId for the product
in order to be idempotent when client want to add a product in the basket
*/
func findProduct(writer http.ResponseWriter, request *http.Request) {
	transactionId := uuid.New().String()
	products := productService.GetAllProduct()
	writeResponse(writer, render.ProductsResponse{TransactionId: transactionId, Products: products})
}

func addProduct(writer http.ResponseWriter, request *http.Request) {
	transactionId := request.Header.Get("transactionId")
	log.Printf("Add product for trasnsactionId %s!", transactionId)
	decoder := json.NewDecoder(request.Body)
	addProductCommand := command.AddProductCommand{}
	err := decoder.Decode(&addProductCommand)
	if err != nil {
		writeErrorResponse(writer, err)
	}
	productHandler.AddProduct(addProductCommand)
	renderResponse(writer, []byte(""))
}

func removeProduct(writer http.ResponseWriter, request *http.Request) {
	response := "Remove"
	renderResponse(writer, []byte(response))
}

func getArgumentAtIndex(request *http.Request, index int) string {
	return strings.Split(request.URL.Path, "/")[index]
}

func writeResponse(response http.ResponseWriter, t interface{}) {
	jsonResponse, err := json.Marshal(t)
	if err != nil {
		writeErrorResponse(response, err)
	} else {
		writeSuccessResponse(response, jsonResponse)
	}
}

func writeSuccessResponse(response http.ResponseWriter, jsonResponse []byte) {
	response.WriteHeader(http.StatusOK)
	_, _ = response.Write(jsonResponse)
}

func writeErrorResponse(response http.ResponseWriter, err error) {
	response.WriteHeader(http.StatusServiceUnavailable)
	errorResponse, _ := json.Marshal("Error in request since " + err.Error())
	_, _ = response.Write(errorResponse)
}

func renderResponse(writer http.ResponseWriter, response []byte) {
	code, err := writer.Write(response)
	if err != nil {
		log.Println("Error rendering response. Caused by ")
	} else {
		log.Printf("Success in response with code %d", code)
	}
}
