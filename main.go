package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

// type User struct {
// 	id       primitive.ObjectID `json:"_id" bosn:"_id"`
// 	Name     string             `json:"Name" bosn:"Name"`
// 	Email    string             `json:"Email" bosn:"Email"`
// 	Password string             `json:"Password" bosn:"Password"`
// }
type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}

// import (
//
// 	//"github.com/Subham-RKB/mongo-golang-INSTA/controllers"
// 	"github.com/Subham-RKB/monog-golang-INSTA/controllers"
// 	"github.com/julienschmidt/httprouter"
// 	"gopkg.in/mgo.v2"
// )

// func main() {
// 	r := httprouter.New()
// 	uc := controllers.NewUserController(getSession())
// 	r.GET("/user/:id", uc.GetUser)
// 	r.POST("/user", uc.CreateUser)
// 	r.DELETE("/user/:id", uc.DeleteUser)
// 	http.ListenAndServe("localhost:8080", r)
// }
// func getSession() *mgo.Session {
// 	s, err := mgo.Dial("mongodb://localhost:2000")
// 	if err != nil {
// 		panic(err)
// 	}
// 	return s
// }
var client *mongo.Client

//MongoDB Connection
func connect_mongodb() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://subham:subham@cluster0.unw6a.mongodb.net/INSTA"))
	if err != nil {
		fmt.Println("Disconected")
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	// collection := client.Database("INSTA").Collection("APPOINTY")

	// ash := User{primitive.NewObjectID(), "Ash", "ash@comp.com", "Pallet Town"}
	// insertResult, err := collection.InsertOne(context.TODO(), ash)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(insertResult)

}

//POST
func PersonEndpoint(w http.ResponseWriter, r *http.Request) {
	connect_mongodb()
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://subham:subham@cluster0.unw6a.mongodb.net/INSTA"))
	// if err != nil {
	// 	fmt.Println("Disconected")
	// 	log.Fatal(err)
	// }
	//response.Header().Set("content-type", "application/json")
	// collection := client.Database("INSTA").Collection("APPOINTY")
	// user00 := User{primitive.NewObjectID(), "Subham", "abc@gmail.com", "abc"}
	// insertResult, err := collection.InsertOne(context.TODO(), user00)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	// var user User
	// _ = json.NewDecoder(request.Body).Decode(&user)

	// ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// collection := client.Database("INSTA")
	// details := collection.Collection("APPOINTY")
	// result, err := details.InsertOne(ctx, bson.D{
	// 	{Key: "Name", Value: "Subham"},
	// 	{Key: "Email", Value: "abc@gmail.com"},
	// 	{Key: "Password", Value: "abc"},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(result.InsertedID)
	// user.Name = "Subham"
	// user.Email = "abc@gmail.com"
	// user.Password = "ok"
	//result, _ := collection.InsertOne(ctx, user00)
	//json.NewEncoder(response).Encode(result)

	collection := client.Database("INSTA").Collection("APPOINTY")

	ash := User{primitive.NewObjectID(), "Ash", "ash@comp.com", "Pallet Town"}
	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insertResult)
}

//
func main() {
	fmt.Println("Starting........")
	//PersonEndpoint(w http.ResponseWriter, r *http.Request)
	//http.ListenAndServe(":8080", nil)
	if err := http.ListenAndServe("", nil); err != nil {
		log.Fatal(err)
	}
}
