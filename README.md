# Simple REST API Golang

## Requirement ##
#

```go
go get github.com/dgrijalva/jwt-go
go get github.com/globalsign/mgo
go get github.com/gorilla/mux
```

## Installation & Run ##
#

```go
go get github.com/scys12/simple-api-go
```

Before you run the application, you should configure your MongoDB configuration in [db.go](https://github.com/scys12/simple-api-go/blob/master/models/db.go)

```go
const (
	host   = "127.0.0.1:27017"
	source = "admin"
	user   = "user"
	pass   = "123456"
)

```
After that you could build and run the application
```go
go build
./simple-api-go
```
The application will run on http://127.0.0.1:8080

## API ##
#
/books
* ```GET``` : Get All Books
* ```POST``` : Post a Book
* ```PUT``` : Update a Book

/books/{id}
* ```GET``` : Get a book
* ```DELETE``` : Delete a book

/user/login
* ```POST``` : Login to an account

/user/register
* ```POST``` : Register an account

## Authentication ##
#
 Once you login, you will get the token authentication. Put in Authorization header
 ```go 
 Authorization : Bearer [token_authentication]
 ```