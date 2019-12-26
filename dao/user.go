package dao

import (
	"context"
	"github.com/1zidorius/project-management-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"log"
)

const (
	DBName             = "project-management"
	USERCOLLECTION     = "users"
	tasksCollection    = "tasks"
	projectsCollection = "projects"
	URI                = "mongodb://admin:admin@localhost:27017"
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
	user.Id, _ = uuid.New()
	_, err := db.Collection(USERCOLLECTION).InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
}

func GetAllUsers() []models.User {
	cur, err := db.Collection(USERCOLLECTION).Find(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	var users []models.User
	var user models.User
	for cur.Next(context.Background()) {
		err := cur.Decode(user)
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

func UpdateUser(user models.User) {
	filter := bson.M{"id": user.Id}
	update := bson.M{"$set": bson.M{"password": user.Password, "email": user.Email, "name": user.Name, "surname": user.Surname}}
	_, err := db.Collection(USERCOLLECTION).UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}
