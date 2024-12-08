package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

// Role represents the structure of the 'role' collection in MongoDB
type Role struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    NmRole    string             `bson:"nm_role" json:"nm_role"`
    CreatedAt time.Time          `bson:"created_at" json:"created_at"`
    CreatedBy primitive.ObjectID `bson:"created_by" json:"created_by"`
    UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
    UpdatedBy primitive.ObjectID `bson:"updated_by" json:"updated_by"`
}
