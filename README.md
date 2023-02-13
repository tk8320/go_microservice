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

`git clone https://github.com/tk8320/go_microservice app/`
`git checkout master`
`cd app`
`docker build -t go_microservice .`
`docker run -d -p 8080:8080 go_microservice`
   
  **Testing:**
  Automated test are already ran at the time of testing.
  If need to run the test again then open docker terminal and run go inside `/app` directory and run `go test -v .`
  The test output will be printed on screen
