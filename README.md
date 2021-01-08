# Hexagolang
A project example of Hexagonal architecture.

In this project we apply architect design patterns like ```CQRS:Command-Query-responsibility-segregation``` and ``Event Sourcing``

We also implement ```idempotent``` for **POST** Operations


### Commands

#### Start Order
````
curl  http://localhost:1981/order/
````

#### Create Order
````
curl -H "transactionId:the_transactionId" http://localhost:1981/order/create/
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
