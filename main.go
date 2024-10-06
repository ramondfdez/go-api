package main

import (
    "context"
    "fmt"
    "log"
    "net/http"

    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

    "github.com/gin-gonic/gin"
	"go-api/handler"
)

func main() {
    // Initialize MongoDB
    db, err := initDB()
    if err != nil {
        log.Fatal("Failed to connect to MongoDB:", err)
    }

    // Create a new Gin router
    router := gin.Default()

    // Define your routes
    router.POST("/api/todos", func(c *gin.Context) {
        handler.CreateTodoHandler(c, db)
    })

    router.GET("/api/todos", func(c *gin.Context) {
        handler.TodosListHandler(c, db)
    })

    router.GET("/api/todos/:id", func(c *gin.Context) {
        handler.GetTodoHandler(c, db)
    })

    router.PATCH("/api/todos/:id", func(c *gin.Context) {
        handler.EditTodoHandler(c, db)
    })

    router.DELETE("/api/todos/:id", func(c *gin.Context) {
        handler.DeleteTodoHandler(c, db)
    })

    router.GET("/api/healthchecker", handler.HealthCheckerHandler)

    // Start the server
    fmt.Println("🚀 Server started successfully on port 8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}

// Initialize MongoDB connection
func initDB() (*mongo.Database, error) {
    clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        return nil, err
    }

    err = client.Ping(context.TODO(), nil)
    if err != nil {
        return nil, err
    }

    fmt.Println("Connected to MongoDB!")
    return client.Database("todo_db"), nil
}
