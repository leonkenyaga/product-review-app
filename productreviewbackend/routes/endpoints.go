package routes

import (
	//"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	//"path/filepath"
	"log"

	"github.com/leon/review-module/models"
	"github.com/leon/review-module/models/db"
	//"go.mongodb.org/mongo-driver/bson"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*func createProfile(w http.ResponseWriter, r *http.Request) {
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
	}*/
	
	
	
	func UploadFile(w http.ResponseWriter, r *http.Request) {
		//file, filename string
        filename:="first_file"

	     fmt.Println("reading")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			
			log.Fatal(err)
		    
		}

//initiate connection
fmt.Println("connecting")
		conn:= db.InitiateMongoClient()
    bucket, err := gridfs.NewBucket(
        conn.Database("myfiles"),
    )
    if err != nil {
		
        log.Fatal(err)
        os.Exit(1)
    }
		//initiate stream
		fmt.Println("streaming")
		uploadStream, err := bucket.OpenUploadStream(
			filename, // this is the name of the file which will be saved in the database
	)
	if err != nil {
		fmt.Println("streaming error")
			fmt.Println(err)
			os.Exit(1)
		}
		defer uploadStream.Close()
	
		fileSize, err := uploadStream.Write(data)
		if err != nil {
			fmt.Println("streaming error2")
			log.Fatal(err)
			os.Exit(1)
		}
		//log.Printf("Write file to DB was successful. File size: %d \n", fileSize)
		json.NewEncoder(w).Encode(fileSize)
	}

	
	
	func ServeGridFSFile(w http.ResponseWriter, r *http.Request) {
		// Create a new MongoDB client and connect to the server
		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			fmt.Println(err)
			// handle error
		}
	
		// Open the GridFS bucket and get a handle to the file
		bucket, err := gridfs.NewBucket(
			client.Database("myfiles"),
			options.GridFSBucket().SetName("fs"),
		)
		if err != nil {
			// handle error
			fmt.Println(err)
		}

		var user models.Product
		json.NewDecoder(r.Body).Decode(&user)

		file, err := bucket.OpenDownloadStreamByName(user.Tag, nil)
		if err != nil {
			// handle error
			fmt.Println(err)
		}
		defer file.Close()
	
		// Set the content type and length headers
		

		// Write the file contents to the response
		if _, err := io.Copy(w, file); err != nil {
			// handle error
			fmt.Println(err)
		}
	}
	




	
	
	


		