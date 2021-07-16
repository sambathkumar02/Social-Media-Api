package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func Follow(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(response, "Method Not Allowed", http.StatusMethodNotAllowed)
	} else {
		var UpdateData FollowStruct
		data_body, _ := ioutil.ReadAll(request.Body)
		err = json.Unmarshal(data_body, &UpdateData)
		filter := bson.M{"username": UpdateData.Username}
		inc := bson.M{"followingcount": 1}
		IncrementFollower := bson.M{"$inc": inc}
		collection := client.Database("SocialMedia").Collection("profiles")
		result, _ := collection.UpdateOne(context.TODO(), filter, IncrementFollower)
		fmt.Print(result)
		addfollowing := bson.M{"following": UpdateData.FollowedAccount}
		UpdateFollowerList := bson.M{"$push": addfollowing}
		fmt.Print(UpdateFollowerList)
		result1, err := collection.UpdateOne(context.TODO(), filter, UpdateFollowerList)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result1)

	}
}

func UnFollow(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(response, "Method Not Allowed", http.StatusMethodNotAllowed)
	} else {
		var UpdateData FollowStruct
		data_body, _ := ioutil.ReadAll(request.Body)
		err = json.Unmarshal(data_body, &UpdateData)
		filter := bson.M{"username": UpdateData.Username}
		dec := bson.M{"followingcount": -1}
		DecrementFollower := bson.M{"$inc": dec}
		collection := client.Database("SocialMedia").Collection("profiles")
		result, _ := collection.UpdateOne(context.TODO(), filter, DecrementFollower)
		fmt.Print(result)
		addfollowing := bson.M{"following": UpdateData.FollowedAccount}
		UpdateFollowerList := bson.M{"$pull": addfollowing}
		fmt.Print(UpdateFollowerList)
		result1, err := collection.UpdateOne(context.TODO(), filter, UpdateFollowerList)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result1)

	}
}

//Need to fix the non atribute injection
func UpdateProfile(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(response, "Method Not Allowed", http.StatusMethodNotAllowed)
	} else {
		//_ = request.ParseForm()
		//for key, value := range request.PostForm {
		//}
		var test UserUpdate
		//var result_map map[string]interface{} //for unsyruc tured json
		data_bytes, _ := ioutil.ReadAll(request.Body)
		//json.Unmarshal(data_bytes, &result_map)
		//fmt.Print(result_map)
		json.Unmarshal(data_bytes, &test)
		//err = bson.UnmarshalExtJSON(data_bytes, true, &test)
		fmt.Print(test)

		//The Unstructured data is stored in map object

		filter := bson.M{"username": "ManiMaran_b"}

		//It Automatically get value from map //Append the Update Query
		update_query := bson.M{"$set": test}

		collection := client.Database("SocialMedia").Collection("profiles")

		_, err := collection.UpdateOne(context.TODO(), filter, update_query)
		if err != nil {
			http.Error(response, "Update Failed", http.StatusBadRequest)
		} else {
			http.Error(response, "Updated", http.StatusAccepted)
			js, _ := json.Marshal(test)
			response.Write(js)

		}
		//Technique works change map to bson and perform operation--study structured and unstructured json formats
	}
}

func UsersList(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var result User
	collection := client.Database("SocialMedia").Collection("profiles")
	err = collection.FindOne(context.TODO(), bson.M{"username": "skumar123"}).Decode(&result)
	result.Followerscount = result.Followerscount + 1
	_, _ = collection.InsertOne(context.TODO(), result)

	js, _ := json.Marshal(result)
	response.Write(js)

}

func CreateUserProfile(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(response, "Method Not Allowed", http.StatusMethodNotAllowed)
	} else {
		//decoder := json.NewDecoder(request.Body)    //create new json decoder and pass the request body
		var data User
		//err = decoder.Decode(&data)  //Decode the Jsom data to the struct
		//fmt.Print(data)
		request_data, _ := ioutil.ReadAll(request.Body)        //Another Method
		err = bson.UnmarshalExtJSON(request_data, true, &data) //Converting json bytes  to bson directly
		collection := client.Database("SocialMedia").Collection("profiles")
		_, _ = collection.InsertOne(context.TODO(), data)
		js, _ := json.Marshal(data)
		response.Write(js)

	}

}
