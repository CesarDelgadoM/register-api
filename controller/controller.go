package controller

import (
	"github.com/CesarDelgadoM/register-API/middleware"
	"github.com/CesarDelgadoM/register-API/services"
	"github.com/gorilla/mux"
)

type Controller struct {
	service *services.RegisterService
	router  *mux.Router
}

func New() *Controller {
	return &Controller{
		service: services.New(),
		router:  mux.NewRouter(),
	}
}

func (controller *Controller) GetRouter() *mux.Router {
	return controller.router
}

func (controller *Controller) RoutesRegister() {
	controller.router.HandleFunc("/register/save",
		middleware.SetMiddlewareJson(controller.service.SaveRegister)).Methods("POST")
	controller.router.HandleFunc("/register/update/{id}",
		middleware.SetMiddlewareJson(controller.service.UpdateRegister)).Methods("PUT")
}
