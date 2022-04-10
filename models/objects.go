package models

type Objects struct {
	ObjId     uint32 `gorm:"primary_key;auto_increment" json:"obj_id"`
	ObjIdUser uint32 `json:"obj_id_user"`
	ObjType   string `gorm:"size:15" json:"obj_type"`
}
