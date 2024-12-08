package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Modul represents the structure of a module document in MongoDB
type Modul struct {
	ID      	 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	IDKategori   primitive.ObjectID `bson:"id_kategori" json:"id_kategori"`
	NmModul      string             `bson:"nm_modul" json:"nm_modul"`
	KetModul     string             `bson:"ket_modul" json:"ket_modul"`
	IsAktif      string             `bson:"is_aktif" json:"is_aktif"`
	Alamat       string             `bson:"alamat" json:"alamat"`
	Urutan       int                `bson:"urutan" json:"urutan"`
	GbrIcon      string             `bson:"gbr_icon" json:"gbr_icon"`
	CreatedAt    primitive.DateTime `bson:"created_at" json:"created_at"`
	CreatedBy    primitive.ObjectID `bson:"created_by" json:"created_by"`
	UpdatedAt    primitive.DateTime `bson:"updated_at" json:"updated_at"`
	UpdatedBy    primitive.ObjectID `bson:"updated_by" json:"updated_by"`
	Icon         string             `bson:"icon" json:"icon"`
}
