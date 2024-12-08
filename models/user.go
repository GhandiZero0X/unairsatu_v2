package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Moduls struct {
	ModulID  primitive.ObjectID `json:"modul_id" bson:"modul_id"`
	NmModul  string             `json:"nm_modul" bson:"nm_modul"`
	KetModul string             `json:"ket_modul" bson:"ket_modul"`
	Alamat   string             `json:"alamat" bson:"alamat"`
	GbrIcon  string             `json:"gbr_icon" bson:"gbr_icon"`
}

type User struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username      string             `json:"username" bson:"username"`
	Nm_user       string             `json:"nm_user" bson:"nm_user"`
	Pass          string             `json:"pass" bson:"pass"`
	Email         string             `json:"email" bson:"email"`
	Role_aktif    int 				 `json:"role_aktif" bson:"role_aktif"`
	Created_at    primitive.DateTime `json:"created_at" bson:"created_at"`
	Updated_at    primitive.DateTime `json:"updated_at" bson:"updated_at"`
	Created_by    primitive.ObjectID `json:"created_by" bson:"created_by"`
	Updated_by    primitive.ObjectID `json:"updated_by" bson:"updated_by"`
	AuthKey       string             `json:"auth_key" bson:"auth_key"`
	Jenis_kelamin int                `json:"jenis_kelamin" bson:"jenis_kelamin"`
	Photo         string             `json:"photo" bson:"photo"`
	Phone         string             `json:"phone" bson:"phone"`
	Token         string             `json:"token" bson:"token"`
	Id_jenis_user primitive.ObjectID `json:"id_jenis_user" bson:"id_jenis_user"`
	Pass_2        string             `json:"pass_2" bson:"pass_2"`
	Moduls        []Moduls           `json:"moduls" bson:"moduls"`
}
