package dao

import (
	"context"
	"fmt"
	"github.com/1zidorius/project-management-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	DBName             = "project-management"
	USERCOLLECTION     = "users"
	tasksCollection    = "tasks"
	projectsCollection = "projects"
	URI                = "mongodb://localhost:27017"
)

var db *mongo.Database

func init() {
	clientOption := options.Client().ApplyURI(URI)
	client, err := mongo.NewClient(clientOption)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database(DBName)
}

func CreateUser(user models.User) {
	//user.Id = primitive.NewObjectID()
	fmt.Println(user)
	_, err := db.Collection(USERCOLLECTION).InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}

func GetAllUsers() []*models.ResultUser {
	cur, err := db.Collection(USERCOLLECTION).Find(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	users := make([]*models.ResultUser, 0)
	for cur.Next(context.Background()) {
		user := &models.ResultUser{}
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())

	return users
}

func DeleteUser(user models.User) {
	_, err := db.Collection(USERCOLLECTION).DeleteOne(context.Background(), user, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateUser(user models.User) error {
	filter := bson.M{"id": user.Id}
	update := bson.M{"$set": bson.M{"password": user.Password, "email": user.Email, "name": user.Name, "surname": user.Surname}}
	_, err := db.Collection(USERCOLLECTION).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
}
