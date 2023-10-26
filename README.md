# payment-apps-backend

This API built with Golang programming language. The purpose of this API is to facilitate transactions between customers and merchants


## Package

This API using Go package:
-   [gin-gonic/gin](https://github.com/gin-gonic/gin)
-   [golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt)
-   [google/uuid](https://github.com/google/uuid)
-   [joho/godotenv](https://github.com/joho/godotenv)
-   [sirupsen/logrus](https://github.com/sirupsen/logrus)
-   [golang.org/x/crypto/bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt])
-   [gorm.io/driver/postgresql](https://pkg.go.dev/gorm.io/driver/postgresql)
-   [gorm.io/gorm](https://pkg.go.dev/gorm.io/gorm)

## Deployment

 - Make sure your machine have Golang installed
 - Clone this repository to your machine.
 - Install all the package that mentioned ablove.
 - Create database and table based on `db.sql` file.
 - Go to application's directory
 - Run the application with terminal using `go run main.go`
