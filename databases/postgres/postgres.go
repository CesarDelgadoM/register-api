package postgres

import "github.com/CesarDelgadoM/register-API/databases/database"

type Postgres struct {
	database.DataBase
}

var psql *Postgres

func Connect() {

}

func GetDB() *Postgres {
	return psql
}
