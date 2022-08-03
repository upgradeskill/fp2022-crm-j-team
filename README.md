# Golang mini-pos - Majoo Bootcamp

## Setup project
 - Clone project
 - Run command
    ```
    go mod tidy
    ```

## Run project

 - Go to root folder
 - Run command
    ```
    go run ./cmd/main.go
    ```
- Server will run in `localhost:8000`

## API List
 - User
    - POST api/v1/login
    - GET api/v1/users
    - GET api/v1/users/:id
    - PUT api/v1/users/:id
    - DELETE api/v1/users/:ID
 - Outlet
 - Category
    - POST api/v1/category
    - GET api/v1/category
    - GET api/v1/category/:id
    - PUT api/v1/category/:id
    - DELETE api/v1/category/:id
 - Product
    - POST api/v1/product
    - GET api/v1/product
    - GET api/v1/product/:id
    - PUT api/v1/product/:id
    - DELETE api/v1/product/:id
 - Transaction
