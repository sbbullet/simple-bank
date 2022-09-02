# Simple Bank
A simple bank API using Golang, Postgres and gRPC

# Instructions
1.  Install `Docker` on your machine.
2.  Clone this project and navigate to the project directory.
3.  Run the following command.
    ```sh
    docker compose up
    ```
4. Navigate to `localhost:8080/swagger` for API documentation. This documentation has only two endpoints that are implemented using gRPC gateway. For all the APIs of Simple Bank project comment out gRPC gateway server initialization code from `main.go` and uncomment Gin server initialization code as below.
```go
	// gRPC Gateway server
	// This server has only two endpoints viz creating user and logging in user
	//go runGatewayServer(config, store)
	//runGrpcServer(config, store)

	// only Gin server
	// This server has all APIs of Simple Bank
	runGinServer(config, store)
```
