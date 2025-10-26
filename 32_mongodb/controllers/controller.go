package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongodbinstallation/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017/"

const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	//client options
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection success")
	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")
}

func insertOneMovie(movie models.Netflix) {
	//package name this thing comes from
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(inserted.InsertedID)

}

func updateOneMovie(movieId string) {

	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.ModifiedCount)
}
func deleteOneMovie(movieId string) {

	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.DeletedCount)
}
func deleteAll() int64 {

	result, err := collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.DeletedCount)
	return result.DeletedCount
}

func getAllMovies() []bson.M {
	var result []bson.M
	cursor,err:=collection.Find(context.TODO(),bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()){
		var movie bson.M
		cursor.Decode(&movie)
		result=append(result,movie)
	}
	return result
}

//all actual controllers that would be used

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	// var movie models.Netflix
	// _=json.NewDecoder(r.Body).Decode(&movie)
	// insertOneMovie(movie)
	var movie models.Netflix
	_=json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)		
}

func WatchedMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")
	params:=mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteOneMovies(w http.ResponseWriter,r *http.Request){
w.Header().Set("Content-Type","application/json")
w.Header().Set("Allow-Control-Allow-Methods","DELETE")
params:=mux.Vars(r)
deleteOneMovie(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	count:=deleteAll()
	json.NewEncoder(w).Encode(count)
}


