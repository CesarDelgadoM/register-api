package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/CesarDelgadoM/register-API/databases/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	database.DataBase
}

func New() *Postgres {
	return &Postgres{}
}

func (psql *Postgres) Connect() {
	dbhost := os.Getenv("DB_HOST")
	dbuser := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")
	dns := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		dbhost, dbuser, dbpassword, dbname, dbport)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("connection to postgres failed:", err)
	} else {
		log.Printf("connection successful")
	}
	psql.DataBase.DB = db
}
