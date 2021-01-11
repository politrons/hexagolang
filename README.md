# HexaGolang

An Order service project example implemented with Hexagonal architecture.

# ![alt text](img/ddd.png) 

In this project we apply architect design patterns like ```CQRS:Command-Query-responsibility-segregation``` and ``Event Sourcing``

We also implement ```idempotent``` for **POST** Operations


### Commands

#### Start Order
````
curl  http://localhost:1981/order/
````

#### Create Order
````
curl -H "transactionId:8052c985-d1d8-480a-9e6c-ef5d0ed27126" http://localhost:1981/order/create/
````

#### Find Order
````
curl http://localhost:1981/order/find/8052c985-d1d8-480a-9e6c-ef5d0ed27126/
````

#### Find Product
````
curl http://localhost:1981/product/
````

#### Add Product 
````
curl -H "transactionId:9f990929-74a7-44b2-a3dc-0a9abc3f3f61" -X POST -d "{\"OrderId\": \"8052c985-d1d8-480a-9e6c-ef5d0ed27126\",\"Id\": \"1234\",\"Price\": 10,\"Description\": \"Coca-Cole\"}" http://localhost:1981/product/add/
````

#### Remove Product 
````
curl -H "transactionId:9f990929-74a7-44b2-a3dc-0a9abc3f3f62" -X POST -d "{\"OrderId\": \"8052c985-d1d8-480a-9e6c-ef5d0ed27126\",\"Id\": \"1234\"}" http://localhost:1981/product/remove/
````
