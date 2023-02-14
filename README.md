# GO REST MICROSERVICE

This is an Order Management microservice built using Golang and MySQL.

The Directory structure is as follows

```
app
│
├── src
|	├── routes/
|	│   	└── routes.go
|	└── models
| 			└── models.go
├── pkg
├── bin
├── README.md
├── microservice_test.go
├── Dockerfile
└── main.go
```
**Files overview:**
- route.go    --> This file contains all of the routes function
- models.go --> This file has the database models 
- main.go  	 --> This file is an  executable and has server initiation logic
- microservice_test.go --> Contains unit test for microservice
-  Dockerfile -> Docker build specification
 
**Packages Required**

-  [Gorilla Mux] --> github.com/gorilla/mux
- [MySQL Connector] --> github.com/go-sql-driver/mysql

Prerequisite to run on docker:
- Having Docker installed in system

**Commands to run the service :**

```
git clone https://github.com/tk8320/go_microservice app/
cd app
docker build -t go_microservice .
docker run -d -p 8080:8080 go_microservice
```
   
  **Testing:**
  Automated test are already ran at the time of testing.
  If need to run the test again then open docker terminal and run go inside `/app` directory and run `go test -v .`
  The test output will be printed on screen

**API ENDPOINTS**
1.  "/order" : 
		Allowed Methods : ["GET"]
		Description : View all orders 

2. "/order/id : int " :
		Allowed Methods : ["GET", "PUT", "DELETE"]
		Description : Get order info by order id 
	
3. "order/create" :
		Allowed Methods : ["POST"]
		Description : Create order 
		Payload: `{"status":  "ORDER_CREATED","items":  [{"id":  12,"Description":  "Fruit Slushie","price":  29.99,"quantity":  1}],"total":  29.99,"currencyUnit":  "USD"}`

4. "order/search":
		Allowed Methods : ["POST"]
		Description : Search orders and sort
		Payload:
		`{"limit":5,"orderAsc":"id","currencyUnit":"USD","status":"ORDER_PLACED"}`
		Search Parameters : ["status", "currencyUnit"]
		Sort Parameters : 
		limit : limit the results
		orderAsc : order by fieldname in ascending order
		orderDesc : order by fieldname in Descending order
