package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// Kategori represents the structure of the 'kategori' collection in MongoDB
type Kategori struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    NmKategori string            `bson:"nm_kategori" json:"nm_kategori"`
}