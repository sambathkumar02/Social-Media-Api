package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//set for connection URI
var client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

//set the timeout limit
var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

func main() {

	if err != nil {
		fmt.Print("Error Connecting Database!")

	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//Automatically close connection When exits
	defer client.Disconnect(ctx)

	//Get the List of all Databases
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	http.HandleFunc("/users/", UsersList)
	http.HandleFunc("/createprofile", CreateUserProfile)
	http.HandleFunc("/update", UpdateProfile)
	http.HandleFunc("/follow", Follow)
	http.HandleFunc("/unfollow", UnFollow)
	http.HandleFunc("/newpost", NewPost)
	http.HandleFunc("/addlike", AddLike)
	http.HandleFunc("/dislike", DisLike)
	http.HandleFunc("/addcomment", AddComment)
	http.HandleFunc("/deletecomment", DeleteComment)
	//http.HandleFunc("/updatebio", UpdateUserBio)

	log.Fatal(http.ListenAndServe(":8088", nil))

}
