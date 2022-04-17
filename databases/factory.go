package databases

import (
	"log"

	"github.com/CesarDelgadoM/register-API/databases/database"
	"github.com/CesarDelgadoM/register-API/databases/postgres"
)

var psql *postgres.Postgres

func SelectDB(namedb string) database.IDataBase {
	switch namedb {
	case "postgres":
		return psql
	default:
		log.Println("name database not exist: ", namedb)
		return nil
	}
}

func ConnectDB(namedb string) {
	switch namedb {
	case "postgres":
		psql = postgres.New()
		psql.Connect()
	default:
		log.Println("name database not exist: ", namedb)
	}
}
