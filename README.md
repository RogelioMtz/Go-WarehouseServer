# Go-WarehouseServer

Go-WarehouseServer is a simple HTTP server written in Go for managing warehouse items. It uses the Gorilla Mux router for handling HTTP requests.

## Features
- Add new warehouse items
- List all warehouse items
- Remove warehouse items by ID

## Requirements
- Go 1.23.2 or higher

## Usage
1. Run the server:
    ```sh
    go run main.go
    ```
2. Use the provided `warehouseRegistration.rest` file to interact with the server using REST API requests.

## API Endpoints
- `GET /warehouseItems` - List all warehouse items
- `POST /warehouseItems` - Add a new warehouse item
- `DELETE /warehouseItems/{id}` - Remove a warehouse item by ID

## License
This project is licensed under the Mozilla Public License Version 2.0
