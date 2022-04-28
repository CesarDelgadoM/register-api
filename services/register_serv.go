package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/CesarDelgadoM/register-API/dto"
	"github.com/CesarDelgadoM/register-API/models"
	"github.com/CesarDelgadoM/register-API/repository"
	"github.com/CesarDelgadoM/register-API/response"
	"github.com/CesarDelgadoM/register-API/utils"
)

type RegisterService struct {
	repo *repository.RegisterRepo
}

func NewRegisterService() *RegisterService {
	return &RegisterService{
		repo: repository.NewRegisterRepo(),
	}
}

func (service *RegisterService) SaveRegister(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(rw, http.StatusUnprocessableEntity, err)
		return
	}
	reg := &models.Register{}
	reg.RegCheckIn = time.Now()
	err = json.Unmarshal(body, reg)
	if err != nil {
		response.Error(rw, http.StatusUnprocessableEntity, err)
		return
	}
	reg, err = service.repo.SaveOne(reg)
	if err != nil {
		response.Error(rw, http.StatusInternalServerError, err)
		return
	}
	response.Json(rw, http.StatusCreated, reg)
}

func (service *RegisterService) GetRegister(rw http.ResponseWriter, r *http.Request) {
	repoObjects := repository.NewRepository(models.Object{})
	id := utils.GetIdUrl("id", r)
	reg, err := service.repo.GetById("reg_id", id)
	if err != nil {
		response.Error(rw, http.StatusBadRequest, err)
		return
	}
	obj, err := repoObjects.GetAllById("obj_reg_id", id)
	if err != nil {
		response.Error(rw, http.StatusUnprocessableEntity, err)
		return
	}
	reg.Objects = obj
	regdto := dto.ModeltoDto(reg)
	response.Json(rw, http.StatusFound, regdto)
}

func (service *RegisterService) GetAllRegisters(rw http.ResponseWriter, r *http.Request) {
	repoObjects := repository.NewRepository(models.Object{})
	regs, err := service.repo.GetAll()
	if err != nil {
		response.Error(rw, http.StatusBadRequest, err)
		return
	}
	objs, err := repoObjects.GetAll()
	if err != nil {
		response.Error(rw, http.StatusInternalServerError, err)
		return
	}
	regsdto := service.buildResponseModels(regs, objs)
	response.Json(rw, http.StatusOK, regsdto)
}

func (service *RegisterService) buildResponseModels(regs *[]models.Register, objs *[]models.Object) *[]dto.RegisterDTO {
	for i, reg := range *regs {
		for _, obj := range *objs {
			if reg.RegId == obj.ObjRegId {
				(*regs)[i].Objects = append((*regs)[i].Objects, obj)
			}
		}
	}
	regsdto := make([]dto.RegisterDTO, 0)
	for _, reg := range *regs {
		regsdto = append(regsdto, *dto.ModeltoDto(&reg))

	}
	return &regsdto
}

func (service *RegisterService) GetRegisterJoin(rw http.ResponseWriter, r *http.Request) {
	id := utils.GetIdUrl("id", r)
	reg, err := service.repo.GetByIdJoin(uint32(id))
	if err != nil {
		response.Error(rw, http.StatusBadRequest, err)
		return
	}
	response.Json(rw, http.StatusFound, reg)
}

func (service *RegisterService) UpdateRegister(rw http.ResponseWriter, r *http.Request) {
	id := utils.GetIdUrl("id", r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(rw, http.StatusBadRequest, err)
		return
	}
	regup := &models.Register{}
	err = json.Unmarshal(body, regup)
	if err != nil {
		response.Error(rw, http.StatusUnprocessableEntity, err)
		return
	}
	if regup.Objects != nil {
		repoObjects := repository.NewRepository(models.Object{})
		repoObjects.Delete("obj_reg_id", id)
		for i, obj := range regup.Objects {
			obj.ObjRegId = id
			regup.Objects[i] = obj
		}
		repoObjects.SaveAll(&regup.Objects)
	}
	regup, err = service.repo.Update("reg_id", id, map[string]interface{}{
		"reg_name":    regup.RegName,
		"reg_company": regup.RegCompany,
	})
	if err != nil {
		response.Error(rw, http.StatusInternalServerError, err)
	}
	response.Json(rw, http.StatusOK, regup)
}

func (service *RegisterService) UpdateCheckOut(rw http.ResponseWriter, r *http.Request) {
	id := utils.GetIdUrl("id", r)
	regup, err := service.repo.Update("reg_id", id, map[string]interface{}{
		"reg_check_out": time.Now(),
	})
	if err != nil {
		response.Error(rw, http.StatusInternalServerError, err)
		return
	}
	response.Json(rw, http.StatusOK, regup)
}

func (service *RegisterService) DeleteRegister(rw http.ResponseWriter, r *http.Request) {
	id := utils.GetIdUrl("id", r)
	rows, err := service.repo.Delete("reg_id", id)
	if err != nil {
		response.Error(rw, http.StatusUnprocessableEntity, err)
		return
	}
	response.Json(rw, http.StatusOK, ("deleted rows: " + strconv.Itoa(int(rows))))
}
