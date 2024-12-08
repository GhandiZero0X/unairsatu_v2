package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Template represents a single module template for a user type.
type Template struct {
	IDModul primitive.ObjectID `bson:"id_modul" json:"id_modul"`
}

// JenisUser represents a type of user with associated templates.
type JenisUser struct {
	ID			primitive.ObjectID  `bson:"_id" json:"id"`
	NmJenisUser string    			`bson:"nm_jenis_user" json:"nm_jenis_user"`
	Templates   []Template 			`bson:"templates" json:"templates"`
}
