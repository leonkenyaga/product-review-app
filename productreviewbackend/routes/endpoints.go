package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"log"

	"github.com/leon/review-module/models"
	"github.com/leon/review-module/models/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func createProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseMultipartForm(10 << 20)
	file:= r.MultipartForm.File["file"]//initialize an array type file
    it:=[]string{}//array variable to store images
	var tempFile *os.File
	var err error
	//loop through file to store images into array "it"
	for i:=0;i<len(file);i++{
	tempFile, err = ioutil.TempFile("static", "upload-*.png")
	if err != nil {
	fmt.Println(err)
	}
		f,_:=file[i].Open()
		fileBytes, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err)
			}
			tempFile.Write(fileBytes)
			fmt.Println(filepath.Base(tempFile.Name()))
			fmt.Println(tempFile.Name())
			image:="/" + filepath.Base(tempFile.Name())
			 
			 it=append(it,image)

	}
	defer tempFile.Close()

	person := models.Product{
	Images:     it,     //assign "it" to images array                        
	Tag:        r.FormValue("tag"),
	Reviews:    r.PostForm["reviews"],
	Likes:      r.PostForm["likes"],
	}
	ctx,col:=db.Connect()
	insertResult, insertErr := col.InsertOne(ctx,person)
	if insertErr != nil {
		fmt.Println("insertONE Error:",insertErr)
		os.Exit(1)
	}else {
	fmt.Println("BINGO!")
	}
	json.NewEncoder(w).Encode(insertResult.InsertedID) // return the //mongodb ID of generated document
	}
	func getAllUsers(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var results []primitive.M    //slice for multiple documents
		ctx,col:=db.Connect()
		cur, err := col.Find(ctx, bson.D{{}}) //returns a *mongo.Cursor
		if err != nil {
	
			fmt.Println(err)
	
		}
		for cur.Next(context.TODO()) { //Next() gets the next document for corresponding cursor
	
			var elem primitive.M
			err := cur.Decode(&elem)
			if err != nil {
			log.Fatal(err)
			}
	
			results = append(results, elem) // appending document pointed by Next()
		}
		cur.Close(context.TODO()) // close the cursor once stream of documents has exhausted
		json.NewEncoder(w).Encode(results)
	}
	


		