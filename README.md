![Imersão Full Stack && Full Cycle](https://events-fullcycle.s3.amazonaws.com/events-fullcycle/static/site/img/grupo_4417.png)

# Imersao Full Cycle Challenge 01

Given the following gRPC protofile:

```proto
syntax = "proto3";

package github.com.codeedu.codepix;

option go_package = "protofiles/pb";

service ProductService {
    rpc CreateProduct (CreateProductRequest) returns (CreateProductResponse) {};
    rpc FindProducts(FindProductsRequest) returns (FindProductsResponse) {};
}


message CreateProductRequest {
    string name = 1;
    string description = 2;
    float price = 3;
}

message Product {
    int32 id = 1;
    string name = 2;
    string description = 3;
    float price = 4;
}

message CreateProductResponse {
    Product product = 1;
}

message FindProductsRequest{

}

message FindProductsResponse{
   repeated Product products = 1;
}
```

Create a gRPC server in Golang that performs the following two operations:

Create products
Query products
The application should use SQLite as the database and GORM as the ORM. Use the AutoMigrate mode to generate tables automatically.

You should provide a main.go script that, when run using go run main.go, will start the server and make port 50051 accessible via localhost.

## How to run

1. Spin up the Docker container
    ```bash
    docker-compose up -d
    ```
2. Access the container
    ```bash
    docker-compose exec app bash
    ```
3. Run the gRPC server
    ```bash
    go run main.go grpc
    ```
4. Run evans CLI
    ```bash
    evans -r repl
    ```
5. Call the gRPC methods
    ```bash
    call CreateProduct
    call FindProducts
    ```