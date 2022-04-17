package models

type Object struct {
	ObjId    uint32 `gorm:"primary_key;auto_increment" json:"obj_id"`
	ObjType  string `gorm:"size:30" json:"obj_type"`
	ObjModel string `gorm:"size:30" json:"obj_model"`
	ObjRegId uint32 `json:"obj_reg_id"`
}
