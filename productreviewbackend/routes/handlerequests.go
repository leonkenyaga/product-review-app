package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Handlepost(){
	route := mux.NewRouter()
s := route.PathPrefix("/api").Subrouter() //Base Path
route.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))//For serving static files 

//ROUTES
s.HandleFunc("/uploadFile", createProfile).Methods("POST")  //upload           //File handler
s.HandleFunc("/getAllUsers", getAllUsers).Methods("GET") //display //All list of users

log.Print("Server Connected 🚀 ")
log.Fatal(http.ListenAndServe(":8000", route)) //run server
}