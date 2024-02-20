// models/certificate.go
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Certificate struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Issuer      string             `bson:"issuer"`
	ExpiryDate  string             `bson:"expiry_date"`
}
