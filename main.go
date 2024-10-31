package main

import (
	"calculator/calculate"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()
	calRouter := router.PathPrefix("/calculator").Subrouter()
	router.HandleFunc("/health",health).Methods("GET")
	calRouter.HandleFunc("/calculate",calculate.Calculator).Methods("GET")
	calRouter.HandleFunc("/calculate",calculate.Calculate).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000",router))
}

func health(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","Application-json")
	w.WriteHeader(http.StatusFound)
	w.Write([]byte("Running on port 3000"))
}