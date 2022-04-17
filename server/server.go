package server

import (
	"log"
	"net/http"
	"os"

	"github.com/CesarDelgadoM/register-API/controller"
	"github.com/CesarDelgadoM/register-API/databases"
	"github.com/CesarDelgadoM/register-API/models"
)

type Server struct {
	driver     string
	controller *controller.Controller
}

func New() *Server {
	driver := os.Getenv("DB_DRIVER")
	databases.ConnectDB(driver)
	return &Server{
		driver:     driver,
		controller: controller.New(),
	}
}

func (server *Server) Initialize() {
	server.initializeDataBases()
	server.initializeRoutes()
}

func (server *Server) Run(port string) {
	log.Println("Listening to port", port)
	log.Fatal(http.ListenAndServe(":"+port, server.controller.GetRouter()))
}

func (server *Server) initializeDataBases() {
	conn := databases.SelectDB(server.driver).GetConnection()
	conn.AutoMigrate(&models.Register{}, &models.Object{})
}

func (server *Server) initializeRoutes() {
	server.controller.RoutesRegister()
}
