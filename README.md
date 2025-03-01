# Simple API Project

This project is a simple API built using Go. It provides various endpoints to manage products, users, and transactions. The project also includes simple CORS handling.

## Features

### API Endpoints

- **User Endpoints**
  - `POST /user/add`: Add a new user.
  - `GET /users`: Retrieve all users.
  - `GET /user/{id}`: Retrieve a user by ID.
  - `PATCH /user/update/{id}`: Update a user by ID.
  - `DELETE /user/delete/{id}`: Delete a user by ID.

- **Product Endpoints**
  - `POST /product/add`: Add a new product.
  - `GET /products`: Retrieve all products.
  - `GET /product/{id}`: Retrieve a product by ID.
  - `PATCH /product/update/{id}`: Update a product by ID.
  - `DELETE /product/delete/{id}`: Delete a product by ID.

- **Transaction Endpoints**
  - `POST /transaction/add`: Add a new transaction.
  - `GET /transactions`: Retrieve all transactions.
  - `GET /transaction/{code}`: Retrieve a transaction by code.

### Simple CORS

The project includes simple CORS handling to allow cross-origin requests. This is implemented using middleware that filters requests based on IP addresses.

### Database

The project uses a MySQL database to store data for users, products, and transactions. The database connection is managed using a custom `db` package. The driver for mysql is [go-sql-driver/mysql](github.com/go-sql-driver/mysql)

### Error Handling

The project includes comprehensive error handling to ensure that appropriate error messages and status codes are returned for various error scenarios.

## Getting Started

### Prerequisites

- Go 1.16 or later
- MySQL database

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/simple-api.git
   cd simple-api
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Set up the database:
   - Create a MySQL database.
   - Update the database connection details in the `db` package.

4. Run the application:
   ```sh
   go run main.go
   ```

### Usage

You can use tools like `curl` or Postman to interact with the API endpoints. Example requests are provided in the `generated-requests.http` file.

### Database schema looks like

![image](https://github.com/user-attachments/assets/56b56972-442c-4b5d-896b-6c0050b955c5)

