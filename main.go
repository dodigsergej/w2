package main

import (
	"fmt"

	"./db"
	"./mux"
)

func init() {
	fmt.Println("http://localost:8002/devicedata/")
	fmt.Println("*************************")
	fmt.Println("Server startovan")
	fmt.Println("*************************")
}

func main() {
	db.ConnectToDB()
	mux.ReceiveData()
}
