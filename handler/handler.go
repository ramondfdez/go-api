package handler

import (
    "context"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

// Define your Todo struct similar to your Rust schema
type Todo struct {
    ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Title     string             `json:"title,omitempty" bson:"title,omitempty"`
    Content   string             `json:"content,omitempty" bson:"content,omitempty"`
    Completed bool               `json:"completed,omitempty" bson:"completed,omitempty"`
    CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
    UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

// CreateTodoHandler creates a new Todo
func CreateTodoHandler(c *gin.Context, db *mongo.Database) {
    var todo Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Check if the todo already exists
    collection := db.Collection("todos")
    filter := bson.M{"title": todo.Title}
    var existingTodo Todo
    err := collection.FindOne(context.TODO(), filter).Decode(&existingTodo)
    if err == nil {
        c.JSON(http.StatusConflict, gin.H{"error": "Todo with this title already exists"})
        return
    }

    // Insert new todo
    todo.ID = primitive.NewObjectID()
    todo.CreatedAt = time.Now()
    todo.UpdatedAt = time.Now()
    _, err = collection.InsertOne(context.TODO(), todo)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, todo)
}

// TodosListHandler retrieves the list of todos
func TodosListHandler(c *gin.Context, db *mongo.Database) {
    collection := db.Collection("todos")

    // Find all todos
    cursor, err := collection.Find(context.TODO(), bson.D{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer cursor.Close(context.TODO())

    var todos []Todo
    for cursor.Next(context.TODO()) {
        var todo Todo
        cursor.Decode(&todo)
        todos = append(todos, todo)
    }

    c.JSON(http.StatusOK, todos)
}

// GetTodoHandler retrieves a single todo by ID
func GetTodoHandler(c *gin.Context, db *mongo.Database) {
    id := c.Param("id") // Use Gin to get URL params
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var todo Todo
    collection := db.Collection("todos")
    err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&todo)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }

    c.JSON(http.StatusOK, todo)
}

// EditTodoHandler updates an existing todo by ID
func EditTodoHandler(c *gin.Context, db *mongo.Database) {
    id := c.Param("id") // Use Gin to get URL params
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var updateData Todo
    if err := c.ShouldBindJSON(&updateData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    updateData.UpdatedAt = time.Now()

    collection := db.Collection("todos")
    filter := bson.M{"_id": objID}
    update := bson.M{
        "$set": bson.M{
            "title":     updateData.Title,
            "content":   updateData.Content,
            "completed": updateData.Completed,
            "updatedAt": updateData.UpdatedAt,
        },
    }

    _, err = collection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully"})
}

// DeleteTodoHandler deletes a todo by ID
func DeleteTodoHandler(c *gin.Context, db *mongo.Database) {
    id := c.Param("id") // Use Gin to get URL params
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    collection := db.Collection("todos")
    _, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusNoContent, nil)
}

// HealthCheckerHandler provides a basic health check
func HealthCheckerHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "success", "message": "API is running"})
}
