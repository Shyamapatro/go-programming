// main.go
package main

import (
	"context" 
    "log"
    "net/http"

    "github.com/gorilla/mux"
   
    "certificate-management-system/config"
    "certificate-management-system/routes"
)

func main() {
    // Connect to MongoDB
    client, err := config.ConnectMongoDB()
    if err != nil {
        log.Fatal("Failed to connect to MongoDB:", err)
    }
    defer client.Disconnect(context.Background()) // Disconnect when the application exits

    // Access database
    db := client.Database("certificateDB")

    // Initialize router
    router := mux.NewRouter()

    // Setup routes
    routes.SetupRoutes(router, db)

    // Start server
    log.Println("Server started on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
