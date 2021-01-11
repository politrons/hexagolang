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
	"log"
	"net/http"
	"strings"
)

var orderDAO dao.OrderDAO = nil
var orderService OrderService = nil
var orderHandler handler.OrderHandler = nil
var productDAO dao.ProductDAO = nil
var productService ProductService = nil
var productHandler handler.ProductHandler = nil

func main() {

	orderDAO = dao.OrderDAOImpl{}.Create()
	orderService = OrderServiceImpl{OrderDAO: orderDAO}
	orderHandler = handler.OrderHandlerImpl{OrderDAO: orderDAO}
	productDAO = dao.ProductDAOImpl{}
	productService = ProductServiceImpl{ProductDAO: productDAO}
	productHandler = handler.ProductHandlerImpl{OrderDAO: orderDAO}

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
func createOrderId(writer http.ResponseWriter, _ *http.Request) {
	orderId := []byte(uuid.New().String())
	renderResponse(writer, orderId)
}

/**
In order to have an idempotent endpoint,We check if the Order already exist,
if it does not, we create one, otherwise we just return the OrderId.
*/
func createOrder(writer http.ResponseWriter, request *http.Request) {
	transactionId := request.Header.Get("transactionId")
	orderResponse := <-orderService.GetOrder(transactionId)
	if !orderResponse.Exist {
		log.Printf("Creating order for transactionId :%s......", transactionId)
		orderHandler.CreateOrder(transactionId, command.CreateOrderCommand{Id: transactionId})
	}
	renderResponse(writer, []byte(transactionId))
}

/**
Using the orderId, we're able to rehydrate the Order model from all the events persisted,
so using event sourcing, we can recreate the current status of the Order
*/
func findOrder(writer http.ResponseWriter, request *http.Request) {
	orderId := getArgumentAtIndex(request, 3)
	orderResponse := <-orderService.GetOrder(orderId)
	jsonResponse, err := json.Marshal(orderResponse.Order)
	if err != nil {
		panic(err)
	}
	renderResponse(writer, jsonResponse)

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

/**
We extract from headers the transactionId to ensure that the event has not been sent and process into the service twice
implementing then idempotent
*/
func addProduct(writer http.ResponseWriter, request *http.Request) {
	transactionId := request.Header.Get("transactionId")
	log.Printf("Adding product for trasnsactionId %s!", transactionId)
	decoder := json.NewDecoder(request.Body)
	addProductCommand := command.AddProductCommand{}
	err := decoder.Decode(&addProductCommand)
	if err != nil {
		writeErrorResponse(writer, err)
	}
	productHandler.AddProduct(transactionId, addProductCommand)
	renderResponse(writer, []byte(""))
}

func removeProduct(writer http.ResponseWriter, request *http.Request) {
	transactionId := request.Header.Get("transactionId")
	log.Printf("Removing product for trasnsactionId %s!", transactionId)
	decoder := json.NewDecoder(request.Body)
	removeProductCommand := command.RemoveProductCommand{}
	err := decoder.Decode(&removeProductCommand)
	if err != nil {
		writeErrorResponse(writer, err)
	}
	productHandler.RemoveProduct(transactionId, removeProductCommand)
	renderResponse(writer, []byte(""))
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
