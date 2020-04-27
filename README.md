# FriendManagementAPI_Golang

**1. Summary**
-	Programming Language: Golang
-	Packages: 
    + database/sql
    + github.com/go-sql-driver/mysql
    + net/http
    + github.com/gorilla/mux
    + github.com/swaggo/http-swagger
    + github.com/stretchr/testify/assert
    + github.com/stretchr/testify/mock
    + testing
-	Database: MySQL
-	Deployment: Linux, Docker
-	Tools: Goland, Git

**2. Deployment**
- This project was deployed by Docker to Linux server at: http://localhost:8081/
- Optionally, I have supported OpenAPI Specification by the Swagger framework, you can see at:
http://localhost:8081/swagger/index.html#/


**# Prerequisites:**
-	Docker
-	Docker Compose
-	MySQL 5.7 for docker

**# Quick Run Project**

```
      git clone https://github.com/toannx95/FriendManagementAPI_Golang
      
      cd src
      docker-compose up -d
```

**# Run test**

```
      cd src
      go test ./...
```

**# Run test with coverage**

```
      cd src
      go test -cover ./...
```

