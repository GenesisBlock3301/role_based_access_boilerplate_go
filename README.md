# Role based access boilerplate Go
# Step-1:
### Setup Project Structure:
```
role_based_access_boilerplate_go/
├── cmd/
│   └── main.go
├── internal/
│   ├── database/
│   │   └── models.go
│   └── handlers/
│       └── http.go
│   └── service/
│       └── my_service.go
├── test/
│   └── my_service_test.go
└── README.md
```
## Step-2:
- Open project folder
- Command: `go mod init`

## Step-3:
### Install swagger:
1. Follow below link:
```
https://github.com/swaggo/
https://github.com/swaggo/gin-swagger
```
Use bellow command & it will generate DOCs folder:
`swag init`

