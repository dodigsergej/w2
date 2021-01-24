package main

import (
	"fmt"

	"./db"
	"./mux"
)

func init() {
	fmt.Println("http://localost:8002/devicedata/")
	fmt.Println("*************************")
}

func main() {
	db.ConnectToDB()
	mux.ReceiveData()
}
