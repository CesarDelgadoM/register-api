package controller

import (
	"github.com/CesarDelgadoM/register-API/middleware"
	"github.com/CesarDelgadoM/register-API/services"
	"github.com/gorilla/mux"
)

type Controller struct {
	registerServ *services.RegisterService
	router       *mux.Router
}

func NewController() *Controller {
	return &Controller{
		registerServ: services.NewRegisterService(),
		router:       mux.NewRouter(),
	}
}

func (controller *Controller) GetRouter() *mux.Router {
	return controller.router
}

func (controller *Controller) RoutesRegister() {
	controller.router.HandleFunc("/register/getall",
		middleware.SetMiddlewareJson(controller.registerServ.GetAllRegisters)).Methods("GET")
	controller.router.HandleFunc("/register/get/{id}",
		middleware.SetMiddlewareJson(controller.registerServ.GetRegister)).Methods("GET")
	controller.router.HandleFunc("/register/save",
		middleware.SetMiddlewareJson(controller.registerServ.SaveRegister)).Methods("POST")
	controller.router.HandleFunc("/register/update/{id}",
		middleware.SetMiddlewareJson(controller.registerServ.UpdateRegister)).Methods("PUT")
	controller.router.HandleFunc("/register/delete/{id}",
		middleware.SetMiddlewareJson(controller.registerServ.DeleteRegister)).Methods("DELETE")
	controller.router.HandleFunc("/register/checkout/{id}",
		middleware.SetMiddlewareJson(controller.registerServ.UpdateCheckOut)).Methods("PUT")
}
