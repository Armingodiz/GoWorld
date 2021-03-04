package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//User - struct to map with mongodb documents
type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Name      string             `bson:"naem"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
}

//CreateIssue - Insert a new document in the collection.
func CreateUser(user User) error {
	client, err := GetMongoClient()
	if err != nil {
		fmt.Println(err)
		return err
	}
	//Create a handle to the respective collection in the database.
	// client.Database() create a DB if the db doesnt exists.
	collection := client.Database(DB).Collection(Users)
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//CreateMany - Insert multiple documents at once in the collection.
func CreateMany(list []User) error {
	//Map struct slice to interface slice as InsertMany accepts interface slice as parameter
	insertableList := make([]interface{}, len(list))
	for i, v := range list {
		insertableList[i] = v
	}
	client, err := GetMongoClient()
	if err != nil {
		fmt.Println(err)
		return err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(DB).Collection(Users)
	//Perform InsertMany operation & validate against the error.
	_, err = collection.InsertMany(context.TODO(), insertableList)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//PrintList - Print list of users on console
func PrintList(issues []User) {
	for i, u := range issues {
		fmt.Printf("%d: %s    %s\n", i+1, u.Name, u.Email)
	}
}

//GetAllIssues - Get All issues for collection
func GetAllUsers() ([]User, error) {
	//Define filter query for fetching specific document from collection
	filter := bson.D{{}} //     ***************  bson.D{{}} specifies 'all documents'   ****************
	var users []User
	//Get MongoDB connection using connectionhelper.
	client, err := GetMongoClient()
	if err != nil {
		return users, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(DB).Collection(Users)
	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return users, findError
	}
	//Map result to slice
	for cur.Next(context.TODO()) {
		var u User
		err := cur.Decode(&u)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}
	// once exhausted, close the cursor
	cur.Close(context.TODO())
	if len(users) == 0 {
		return users, mongo.ErrNoDocuments
	}
	return users, nil
}

//GetIssuesByCode - Get All issues for collection
func GetUserByName(name string) (User, error) {
	result := User{}
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "naem", Value: name}}
	//Get MongoDB connection using connectionhelper.
	client, err := GetMongoClient()
	if err != nil {
		return result, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(DB).Collection(Users)
	//Perform FindOne operation & validate against the error.
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	//Return result without any error.
	return result, nil
}

func DeleteByName(name string) error {
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "naem", Value: name}}
	//Get MongoDB connection using connectionhelper.
	client, err := GetMongoClient()
	if err != nil {
		return err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(DB).Collection(Users)

	//Perform DeleteOne operation & validate against the error.
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAll() error {
	//Define filter query for fetching specific document from collection
	selector := bson.D{{}} // bson.D{{}} specifies 'all documents'
	//Get MongoDB connection using connectionhelper.
	client, err := GetMongoClient()
	if err != nil {
		return err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(DB).Collection(Users)
	//Perform DeleteMany operation & validate against the error.
	_, err = collection.DeleteMany(context.TODO(), selector)
	if err != nil {
		return err
	}
	return nil
}

func UpdateByName(name string) error {
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "naem", Value: name}}

	//Define updater for to specifiy change to be updated.
	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "updated_at", Value: time.Now()},
	}}}

	//Get MongoDB connection using connectionhelper.
	client, err := GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(DB).Collection(Users)
	//Perform UpdateOne operation & validate against the error.
	_, err = collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}

func main() {
	/*user := User{
		Name:      "armin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     "armin3011",
	}
	CreateUser(user)*/
	//CreateMany(users)
	//us, err2 := GetUserByName("armin")
	//DeleteByName("armin")
	//DeleteAll()
	UpdateByName("armin")
}
