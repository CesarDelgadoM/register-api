package models

import "time"

type Register struct {
	RegId       uint32    `gorm:"primary_key;auto_increment" json:"reg_id"`
	RegNames    string    `gorm:"size:25" json:"reg_names"`
	RegCheckIn  time.Time `json:"reg_check_in"`
	RegCheckOut time.Time `json:"reg_check_out"`
	Objects     []Objects `json:"foreignKey:ObjIdUser;references:RegId"`
}
