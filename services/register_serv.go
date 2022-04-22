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
	"github.com/gorilla/mux"
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
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response.Error(rw, http.StatusBadRequest, err)
		return
	}
	reg, err := service.repo.GetById("reg_id", uint32(id))
	if err != nil {
		response.Error(rw, http.StatusBadRequest, err)
		return
	}
	obj, err := repoObjects.GetAllById("obj_reg_id", uint32(id))
	if err != nil {
		response.Error(rw, http.StatusUnprocessableEntity, err)
		return
	}
	reg.Objects = obj
	regdto := dto.ModeltoDto(reg)
	response.Json(rw, http.StatusFound, regdto)
}

func (service *RegisterService) GetRegisterJoin(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response.Error(rw, http.StatusBadRequest, err)
		return
	}
	reg, err := service.repo.GetByIdJoin(uint32(id))
	if err != nil {
		response.Error(rw, http.StatusBadRequest, err)
		return
	}
	response.Json(rw, http.StatusFound, reg)
}

func (service *RegisterService) UpdateRegister(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response.Error(rw, http.StatusBadRequest, err)
		return
	}
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
		repoObjects.Delete("obj_reg_id", uint32(id))
		for i, obj := range regup.Objects {
			obj.ObjRegId = uint32(id)
			regup.Objects[i] = obj
		}
		repoObjects.SaveAll(&regup.Objects)
	}
	regup, err = service.repo.Update("reg_id", uint32(id), map[string]interface{}{
		"reg_name":    regup.RegName,
		"reg_company": regup.RegCompany,
	})
	if err != nil {
		response.Error(rw, http.StatusInternalServerError, err)
	}
	response.Json(rw, http.StatusOK, regup)
}

func (service *RegisterService) UpdateCheckOut(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response.Error(rw, http.StatusBadRequest, err)
		return
	}
	regup, err := service.repo.Update("reg_id", uint32(id), map[string]interface{}{
		"reg_check_out": time.Now(),
	})
	if err != nil {
		response.Error(rw, http.StatusInternalServerError, err)
		return
	}
	response.Json(rw, http.StatusOK, regup)
}

func (service *RegisterService) DeleteRegister(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response.Error(rw, http.StatusBadRequest, err)
		return
	}
	rows, err := service.repo.Delete("reg_id", uint32(id))
	if err != nil {
		response.Error(rw, http.StatusUnprocessableEntity, err)
		return
	}
	response.Json(rw, http.StatusOK, ("deleted rows: " + strconv.Itoa(int(rows))))
}
