package databases

import (
	"log"

	"github.com/CesarDelgadoM/register-API/databases/database"
	"github.com/CesarDelgadoM/register-API/databases/postgres"
)

func SelectDB(namedb string) database.IDataBase {
	var idb database.IDataBase
	switch namedb {
	case "postgres":
		idb = postgres.GetDB()
	default:
		log.Println("name database not exist: ", namedb)
	}
	return idb
}

func ConnectDB(namedb string) {

}
