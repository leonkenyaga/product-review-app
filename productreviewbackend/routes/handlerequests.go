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
s.HandleFunc("/uploadFile", UploadFile).Methods("POST") 
s.HandleFunc("/DownloadFile", ServeGridFSFile).Methods("GET") 
//s.HandleFunc("/uploadFiles", createProfile).Methods("POST")  //upload           //File handler
//s.HandleFunc("/getAllUsers", getAllUsers).Methods("GET") //display //All list of users

log.Print("Server Connected 🚀 ")
log.Fatal(http.ListenAndServe(":4000", route)) //run server
}