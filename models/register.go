package models

import "time"

type Register struct {
	RegId       uint64    `gorm:"primary_key;auto_increment" json:"reg_id"`
	RegName     string    `gorm:"size:30" json:"reg_name"`
	RegCompany  string    `gorm:"size:30" json:"reg_company"`
	RegCheckIn  time.Time `json:"reg_check_in"`
	RegCheckOut time.Time `json:"reg_check_out"`
	Objects     []Object  `gorm:"foreignKey:ObjRegId;references:RegId;constraint:OnDelete:CASCADE" json:"reg_objects_id"`
}
