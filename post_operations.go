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

func NewPost(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(response, "Method Not Allowed", http.StatusMethodNotAllowed)
	} else {
		var PostData Post
		data_body, _ := ioutil.ReadAll(request.Body)
		err = json.Unmarshal(data_body, &PostData)
		username := request.URL.Query().Get("username")
		filter := bson.M{"username": username}
		collection := client.Database("SocialMedia").Collection("profiles")
		PostBson := bson.M{"posts": PostData}
		UpdatePost := bson.M{"$push": PostBson}
		fmt.Print(UpdatePost)
		result1, err := collection.UpdateOne(context.TODO(), filter, UpdatePost)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result1)

	}

}

func AddLike(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(response, "Method Not Allowed", http.StatusMethodNotAllowed)
	} else {

		//query db.profiles.
		var data AddLikeStruct
		request_body, _ := ioutil.ReadAll(request.Body)
		err = json.Unmarshal(request_body, &data)
		filter := bson.M{"posts.postid": data.Postid}
		incrementlike := bson.M{"posts.$.likescount": 1}
		tuple := bson.M{"username": data.Username, "reaction": data.Reaction}
		addentry := bson.M{"posts.$.likes": tuple}
		UpdateQueryIncrement := bson.M{"$inc": incrementlike}
		UpdateQueryAddEntry := bson.M{"$push": addentry}
		collection := client.Database("SocialMedia").Collection("profiles")
		addlikeresult, _ := collection.UpdateOne(context.TODO(), filter, UpdateQueryIncrement)
		addentryresult, _ := collection.UpdateOne(context.TODO(), filter, UpdateQueryAddEntry)
		fmt.Print(addentryresult, addlikeresult)
	}

}

func DisLike(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(response, "Method Not Allowed", http.StatusMethodNotAllowed)
	} else {

		var data AddLikeStruct
		request_body, _ := ioutil.ReadAll(request.Body)
		err = json.Unmarshal(request_body, &data)
		filter := bson.M{"posts.postid": data.Postid}
		decrementlike := bson.M{"posts.$.likescount": -1}
		tuple := bson.M{"username": data.Username, "reaction": data.Reaction}
		deleteentry := bson.M{"posts.$.likes": tuple}
		UpdateQueryIncrement := bson.M{"$inc": decrementlike}
		UpdateQueryAddEntry := bson.M{"$pull": deleteentry}
		collection := client.Database("SocialMedia").Collection("profiles")
		dellikeresult, _ := collection.UpdateOne(context.TODO(), filter, UpdateQueryIncrement)
		delentryresult, _ := collection.UpdateOne(context.TODO(), filter, UpdateQueryAddEntry)
		fmt.Print(delentryresult, dellikeresult)
	}
}

func AddComment(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(response, "Method Not Allowed", http.StatusMethodNotAllowed)
	} else {

		//query db.profiles.
		var data AddCommentStruct
		request_body, _ := ioutil.ReadAll(request.Body)
		err = json.Unmarshal(request_body, &data)
		filter := bson.M{"posts.postid": data.Postid}
		incrementcomment := bson.M{"posts.$.commentcount": 1}
		tuple := bson.M{"username": data.Username, "commenttext": data.CommentText}
		addentry := bson.M{"posts.$.comment": tuple}
		UpdateQueryIncrement := bson.M{"$inc": incrementcomment}
		UpdateQueryAddEntry := bson.M{"$push": addentry}
		collection := client.Database("SocialMedia").Collection("profiles")
		addcommentresult, _ := collection.UpdateOne(context.TODO(), filter, UpdateQueryIncrement)
		addentryresult, _ := collection.UpdateOne(context.TODO(), filter, UpdateQueryAddEntry)
		fmt.Print(addentryresult, addcommentresult)
	}
}

func DeleteComment(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(response, "Method Not Allowed", http.StatusMethodNotAllowed)
	} else {

		var data AddCommentStruct
		request_body, _ := ioutil.ReadAll(request.Body)
		err = json.Unmarshal(request_body, &data)
		filter := bson.M{"posts.postid": data.Postid}
		decrementcomment := bson.M{"posts.$.commentcount": -1}
		tuple := bson.M{"username": data.Username, "commenttext": data.CommentText}
		deleteentry := bson.M{"posts.$.comment": tuple}
		UpdateQueryIncrement := bson.M{"$inc": decrementcomment}
		UpdateQueryAddEntry := bson.M{"$pull": deleteentry}
		collection := client.Database("SocialMedia").Collection("profiles")
		delcommentresult, _ := collection.UpdateOne(context.TODO(), filter, UpdateQueryIncrement)
		delentryresult, _ := collection.UpdateOne(context.TODO(), filter, UpdateQueryAddEntry)
		fmt.Print(delentryresult, delcommentresult)
	}
}
