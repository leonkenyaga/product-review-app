package db

import (
	"context"
	"fmt"
	//"os"
	"reflect"

	"os"
	//"reflect"
	//"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//enstablish a connection with mongodb

/*func Connect() (context.Context,*mongo.Collection){

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("ClientOptopm TYPE:", reflect.TypeOf(clientOptions))
	client, err := mongo.Connect(context.TODO(),clientOptions)
	if err!=nil{
		fmt.Println("mongo.connect() ERROR: ",err)
		os.Exit(1)
    }
    ctx,_:=context.WithTimeout(context.Background(),15*time.Second)
	col := client.Database("Second_Database").Collection("third COllection")
	fmt.Println("Collection Type:",reflect.TypeOf(col))
	return ctx,col

}*/


func InitiateMongoClient() *mongo.Client {
   /* var err error
    var client *mongo.Client
    uri := "mongodb://localhost:27017"
    opts := options.Client()
    opts.ApplyURI(uri)
    opts.SetMaxPoolSize(5)
    if client, err = mongo.Connect(context.Background(), opts); err != nil {
		fmt.Println("initiation error")
        fmt.Println(err.Error())
    }*/
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("ClientOptopm TYPE:", reflect.TypeOf(clientOptions))
	client, err := mongo.Connect(context.TODO(),clientOptions)
	if err!=nil{
		fmt.Println("mongo.connect() ERROR: ",err)
		os.Exit(1)
    }
	
    return client
}

