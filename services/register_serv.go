package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/CesarDelgadoM/register-API/models"
	"github.com/CesarDelgadoM/register-API/repository"
	"github.com/CesarDelgadoM/register-API/response"
	"github.com/gorilla/mux"
)

type RegisterService struct {
	repo *repository.Repository[models.Register]
}

func New() *RegisterService {
	return &RegisterService{
		repo: repository.New(models.Register{}),
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
	reg, err = service.repo.Save(reg)
	if err != nil {
		response.Error(rw, http.StatusInternalServerError, err)
		return
	}
	response.Json(rw, http.StatusCreated, reg)
}

func (service *RegisterService) GetRegister(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response.Error(rw, http.StatusBadRequest, err)
		return
	}
	reg, err := service.repo.GetById("obj_id", uint32(id))
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
	// objs := []models.Object{}
	// for i, obj := range regup.Objects {
	// 	objs[i] = obj
	// }
	fmt.Println(regup.Objects)
	regup, err = service.repo.Update("reg_id", uint32(id), regup, map[string]interface{}{
		"reg_name":       regup.RegName,
		"reg_company":    regup.RegCompany,
		"reg_objects_id": regup.Objects,
	})
	if err != nil {
		response.Error(rw, http.StatusInternalServerError, err)
	}
	response.Json(rw, http.StatusOK, regup)
}
