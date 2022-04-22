package dto

import (
	"time"

	"github.com/CesarDelgadoM/register-API/models"
)

type RegisterDTO struct {
	RegName     string      `json:"reg_name"`
	RegCompany  string      `json:"reg_company"`
	RegCheckIn  time.Time   `json:"reg_check_in"`
	RegCheckOut time.Time   `json:"reg_check_out"`
	Objects     []ObjectDTO `json:"reg_objects"`
}

func ModeltoDto(model *models.Register) *RegisterDTO {
	objdto := make([]ObjectDTO, len(model.Objects))
	for i, obj := range model.Objects {
		objdto[i] = ObjectDTO{
			ObjType:  obj.ObjType,
			ObjModel: obj.ObjModel,
		}
	}
	return &RegisterDTO{
		RegName:     model.RegName,
		RegCompany:  model.RegCompany,
		RegCheckIn:  model.RegCheckIn,
		RegCheckOut: model.RegCheckOut,
		Objects:     objdto,
	}
}
