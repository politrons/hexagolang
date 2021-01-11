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
curl http://localhost:1981/order/find/{order_id}/
````

#### Find Product
````
curl http://localhost:1981/product/
````

#### Add Product 
````
curl -H "transactionId:the_transactionId" -X POST -d "{\"OrderId\": \"OrderId\",\"Id\": \"Id\",\"Price\": 10,\"Description\": \"Description\"}" http://localhost:1981/product/add/
````

#### Remove Product 
````
curl -H "transactionId:the_transactionId" -X POST -d "{\"OrderId\": \"OrderId\",\"Id\": \"Id\"}" http://localhost:1981/product/remove/
````
