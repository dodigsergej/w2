package mux

import (
	"log"
	"net/http"

	"../parserdata"

	"github.com/gorilla/mux"
)

//ReceiveData func
func ReceiveData() {
	var router = mux.NewRouter()
	router.HandleFunc("/devicedata", processfnc).Methods("POST")

	log.Fatal(http.ListenAndServe(":8002", router))
}

func processfnc(w http.ResponseWriter, r *http.Request) {
	parserdata.ParseData(w, r)
}
