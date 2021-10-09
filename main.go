package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var client *mongo.Client
var lock sync.Mutex

// func Paginate(collection *mongo.Collection, startValue objectid.ObjectID, nPerPage int64) ([]bson.Document, *bson.Value, error) {

//     // Query range filter using the default indexed _id field. 
//     filter := bson.VC.DocumentFromElements(
//         bson.EC.SubDocumentFromElements(
//             "_id",
//             bson.EC.ObjectID("$gt", startValue),
//         ),
//     )

//     var opts []findopt.Find
//     opts = append(opts, findopt.Sort(bson.NewDocument(bson.EC.Int32("_id", -1))))
//     opts = append(opts, findopt.Limit(nPerPage))

//     cursor, _ := collection.Find(context.Background(), filter, opts...)

//     var lastValue *bson.Value
//     var results []bson.Document
//     for cursor.Next(context.Background()) {
//         elem := bson.NewDocument()
//         err := cursor.Decode(elem)
//         if err != nil {
//             return results, lastValue, err
//         }
//         results = append(results, *elem)
//         lastValue = elem.Lookup("_id")
//     }

//     return results, lastValue, nil
// }

func CreateUserEndpoint(response http.ResponseWriter, request *http.Request){
	lock.Lock()
    defer lock.Unlock()
	response.Header().Add("content-type", "application/json")
	var user User
	json.NewDecoder(request.Body).Decode(&user)
	hash, _ := HashPassword(user.Password)
	user.Password = hash
	collection := client.Database("insta").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, user)
	json.NewEncoder(response).Encode(result)
	time.Sleep(1 * time.Second)
}

func GetUserEndpoint(response http.ResponseWriter, request *http.Request){
	lock.Lock()
    defer lock.Unlock()
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var user User
	collection := client.Database("insta").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, User{ID: id}).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	fmt.Println(user.Password)
	json.NewEncoder(response).Encode(user)
	time.Sleep(1 * time.Second)
}

func HashPassword(password string) (string, error) {
	lock.Lock()
    defer lock.Unlock()
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	time.Sleep(1 * time.Second)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	lock.Lock()
    defer lock.Unlock()
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	time.Sleep(1 * time.Second)
    return err == nil
}

func CreatePostEndpoint(response http.ResponseWriter, request *http.Request){
	lock.Lock()
    defer lock.Unlock()
	response.Header().Add("content-type", "application/json")
	var post Post
	json.NewDecoder(request.Body).Decode(&post)
	post.PostedTimestamp= time.Now()
	collection := client.Database("insta").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, post)
	json.NewEncoder(response).Encode(result)
	time.Sleep(1 * time.Second)
}

func GetPostEndpoint(response http.ResponseWriter, request *http.Request){
	lock.Lock()
    defer lock.Unlock()
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var post Post
	collection := client.Database("insta").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, User{ID: id}).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	fmt.Println(post.Caption)
	json.NewEncoder(response).Encode(post)
	time.Sleep(1 * time.Second)
}

func GetUserPosts(response http.ResponseWriter, request *http.Request){
	lock.Lock()
    defer lock.Unlock()
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := params["id"]
	var posts []Post
	collection := client.Database("insta").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, Post{UserID: id})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx){
		var post Post
		cursor.Decode(&post)
		posts = append(posts, post)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(posts)
	time.Sleep(1 * time.Second)

}


func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	router := mux.NewRouter()
	router.HandleFunc("/user", CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/user/{id}", GetUserEndpoint).Methods("GET")
	router.HandleFunc("/post", CreatePostEndpoint).Methods("POST")
	router.HandleFunc("/post/{id}", GetPostEndpoint).Methods("GET")
	router.HandleFunc("/post/user/{id}", GetUserPosts).Methods("GET")
	http.ListenAndServe(":12345", router)
}