package controllers

import (
	"context"
	"log"
    
	  "net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive" // Add this import for primitive.ObjectID

	"github.com/fatih/color" 
	"time" 
	"encoding/json"
	"certificate-management-system/models"
)

func CreateCertificate(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
    // Ensure the request method is POST
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Read the request body
    var requestBody map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Log the request body with color
    requestBodyJSON, err := json.Marshal(requestBody)
    if err != nil {
        log.Println(color.RedString("Error encoding request body:"), err)
    } else {
        log.Println(color.GreenString("Request Body:"), string(requestBodyJSON))
    }

    // Decode the request body into the certificate model
    var cert models.Certificate
    if err := json.Unmarshal(requestBodyJSON, &cert); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Convert expiry_date string to time.Time
    expiryDateStr, ok := requestBody["expiry_date"].(string)
    if !ok {
        http.Error(w, "expiry_date is not a string", http.StatusBadRequest)
        return
    }
    expiryDate, err := time.Parse("2006-01-02", expiryDateStr)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    cert.ExpiryDate = expiryDate.Format("2006-01-02") // Format expiryDate as a string

    // Check if certificate with the same name already exists
    collection := db.Collection("certificates")
    existingCert := models.Certificate{}
    err = collection.FindOne(context.Background(), bson.M{"name": cert.Name}).Decode(&existingCert)
    if err == nil {
        http.Error(w, "Certificate with the same name already exists", http.StatusConflict)
        return
    } else if err != mongo.ErrNoDocuments {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Insert the certificate into the collection
    if _, err := collection.InsertOne(context.Background(), cert); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with success message and status code
    response := map[string]interface{}{
        "status":  http.StatusCreated,
        "message": "Certificate created successfully",
    }
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write(jsonResponse)
}


func GetAllCertificates(db *mongo.Database, w http.ResponseWriter) {
	// Find all certificates in the "certificates" collection
	collection := db.Collection("certificates")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		// Error finding certificates
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	// Slice to store certificates
	var certificates []models.Certificate

	// Iterate over the cursor and decode each document into a Certificate struct
	for cursor.Next(context.Background()) {
		var cert models.Certificate
		if err := cursor.Decode(&cert); err != nil {
			// Error decoding certificate
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		certificates = append(certificates, cert)
	}

	// Check if cursor encountered any errors during iteration
	if err := cursor.Err(); err != nil {
		// Error iterating over cursor
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Respond with the found certificates
	if len(certificates) == 0 {
		http.Error(w, "No certificates found", http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"status":     http.StatusOK,
		"message":    "Certificates found successfully",
		"certificates": certificates,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func UpdateCertificate(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a Certificate struct
	var updatedCert models.Certificate
	if err := json.NewDecoder(r.Body).Decode(&updatedCert); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Define filter to find the certificate by ID
	filter := bson.M{"_id": updatedCert.ID}

	// Initialize the update document with only the fields present in the request body
	update := bson.M{"$set": bson.M{}}
	if updatedCert.Name != "" {
		update["$set"].(bson.M)["name"] = updatedCert.Name
	}
	if updatedCert.Description != "" {
		update["$set"].(bson.M)["description"] = updatedCert.Description
	}
	if updatedCert.Issuer != "" {
		update["$set"].(bson.M)["issuer"] = updatedCert.Issuer
	}
	if updatedCert.ExpiryDate != "" {
		update["$set"].(bson.M)["expiryDate"] = updatedCert.ExpiryDate
	}

	// Access the certificates collection
	collection := db.Collection("certificates")

	// Perform the update operation
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result.ModifiedCount == 0 {
		// No document was modified
		http.Error(w, "No certificate found with the provided ID", http.StatusNotFound)
		return
	}

	// Respond with success message and status code
	response := struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		Status:  http.StatusOK,
		Message: "Certificate updated successfully",
	}
	w.Header().Set("Content-Type", "application/json")
		
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func DeleteCertificate(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a map
	var requestBody map[string]string
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extract the ID from the request body
	idString, ok := requestBody["_id"]
	if !ok || idString == "" {
		http.Error(w, "Certificate ID is required", http.StatusBadRequest)
		return
	}

	// Convert the ID string to a primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Define filter to find the certificate by ID
	filter := bson.M{"_id": objectID}

	// Access the certificates collection
	collection := db.Collection("certificates")

	// Perform the delete operation
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		// No document was deleted
		response := struct {
			Status  int    `json:"status"`
			Message string `json:"message"`
		}{
			Status:  http.StatusNotFound,
			Message: "No certificate found with the provided ID",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Respond with success message and status code in JSON format
	response := struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		Status:  http.StatusOK,
		Message: "Certificate deleted successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GetCertificateByID(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a map
	var requestBody map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extract the ID from the request body
	idString, idOK := requestBody["_id"].(string)
	if !idOK {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	// Convert the ID string to a primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Define filter to find the certificate by ID
	filter := bson.M{"_id": objectID}

	// Find the certificate in the "certificates" collection
	var foundCert models.Certificate
	collection := db.Collection("certificates")
	err = collection.FindOne(context.Background(), filter).Decode(&foundCert)
	if err == mongo.ErrNoDocuments {
		// No documents found
		http.Error(w, "No certificate found", http.StatusNotFound)
		return
	} else if err != nil {
		// Other errors
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Respond with the found certificate
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foundCert)
}