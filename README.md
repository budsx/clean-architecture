# Clean Architecture Golang

## Installation

```bash
go get -u github.com/go-playground/validator/v10
go get -u github.com/julienschmidt/httprouter
go get -u github.com/lib/pq
```

## Folder Structure

```
+-- Controllers
|   +--- category_controller.go
+-- Helper
|   +--- error.go
|   +--- response.go
|   +--- transaction.go
+-- Models
|   +--- Domain
|   |   +--- category.go
|   +--- Web
|   |   +--- category_create_request.go
|   |   +--- category_response.go
|   |   +--- category_update_request.go
+-- Repository
|   +--- category_repository.go
+-- Services
|   +--- category_service.go
|   go.mod
|   go.sum
|   README.md 
```

## Logic
```
Domain -> Repository -> Service -> Controller -> Web
```

```
Domain -> Respresent Table on Database
DTO ->  Request and Response Model
Repository -> Handle Query Database
Service -> Handle Business Logic
Controller -> Handle Request From Web
```


