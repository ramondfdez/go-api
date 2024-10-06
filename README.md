# Go Todo API

![Go](https://img.shields.io/badge/Go-1.23.1-blue.svg)
![Status](https://img.shields.io/badge/status-active-brightgreen.svg)

## Overview

This is a simple Todo API built with Go using the [Gin](https://github.com/gin-gonic/gin) web framework and MongoDB as the database. The API allows users to create, read, update, and delete todo items.

## Features

- **RESTful API**: Standard RESTful endpoints for todo management.
- **MongoDB Integration**: Persistent data storage using MongoDB.
- **Health Check**: Basic health check endpoint to verify the API is running.

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.23.1 or later)
- [Docker](https://www.docker.com/products/docker-desktop) (optional, for running MongoDB in a container)
- MongoDB (can be run locally or in a Docker container)

### Clone the Repository

```bash
git clone https://github.com/your-username/go-api.git
cd go-api
```

###  Setup MongoDB

You can run MongoDB locally or use Docker. To run MongoDB using Docker, execute the following command:

```bash
docker run --name mongodb -d -p 27017:27017 mongo:latest
```
### Install Dependencies

Run the following command to install the required Go modules:

```bash
go mod tidy
```
### Running the Application

To start the API server, run:

```bash
go run main.go
```
The API will start at http://localhost:8000.
### API Endpoints

    POST /api/todos - Create a new todo
    GET /api/todos - Retrieve all todos
    GET /api/todos/{id} - Retrieve a specific todo by ID
    PATCH /api/todos/{id} - Update an existing todo
    DELETE /api/todos/{id} - Delete a todo by ID
    GET /api/healthchecker - Check if the API is running

### Example Requests

Create a new Todo:

    curl -X POST http://localhost:8000/api/todos -H "Content-Type: application/json" -d '{"title": "Sample Todo", "content": "This is a sample todo item", "completed": false}'

Retrieve all Todos:

    curl -X GET http://localhost:8000/api/todos

Retrieve a specific Todo by ID:


    curl -X GET http://localhost:8000/api/todos/{id}

Update a Todo:


    curl -X PATCH http://localhost:8000/api/todos/{id} -H "Content-Type: application/json" -d '{"completed": true}'

Delete a Todo:


    curl -X DELETE http://localhost:8000/api/todos/{id}

Health Check:


    curl -X GET http://localhost:8000/api/healthchecker

### Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any enhancements or bug fixes.

### License

This project is licensed under the MIT License - see the LICENSE file for details.