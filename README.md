# GoTestServer

GoTestServer is a simple HTTP server written in Go for managing warehouse items. It uses the Gorilla Mux router for handling HTTP requests.

## Features
- Add new warehouse items
- List all warehouse items
- Remove warehouse items by ID

## Requirements
- Go 1.23.2 or higher

## Installation
1. Clone the repository:
    ```sh
    git clone https://github.com/RogelioMtz/server.git
    ```
2. Navigate to the project directory:
    ```sh
    cd server
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```

## Files

### main.go
The `main.go` file is the main entry point for the GoTestServer application. It sets up an HTTP server using the Gorilla Mux router and listens on localhost at port 8080.

### go.mod
The `go.mod` file defines the module and its dependencies.

### server.go
The `server.go` file contains the implementation of the server, including the router setup and handlers for the API endpoints.

### warehouseRegistration.rest
The `warehouseRegistration.rest` file contains HTTP requests for interacting with the GoTestServer API. You can use this file with a REST client like VS Code's REST Client extension to test the API endpoints.

## Usage
1. Run the server:
    ```sh
    go run main.go
    ```
2. Use the `warehouseRegistration.rest` file to interact with the server using REST API requests.

## API Endpoints
- `GET /warehouseItems` - List all warehouse items
- `POST /warehouseItems` - Add a new warehouse item
- `DELETE /warehouseItems/{id}` - Remove a warehouse item by ID

## License
This project is licensed under the Mozilla Public License 2.0.
