package db

import (
	"database/sql"
	"fmt"

	//MySQL
	_ "github.com/go-sql-driver/mysql"
)

//Tag struktura
type Tag struct {
	ID  int
	CIP string
}

var dba sql.DB

//ConnectToDB -- konekcija na bazu podataka
func ConnectToDB() {
	var host string = "127.0.0.1"
	var username string = "root"
	var password string = "sd1906971"
	var port string = "3306"
	var database string = "IMP_DB"

	dba, err := sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+database)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("DB stats:")
	fmt.Println(dba.Stats())
	//defer dba.Close()
}

//StoreData -- snimi detalje u bazu
func StoreData() bool {
	var RetVal bool = false

	return RetVal
}
