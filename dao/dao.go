package dao

import (
	"context"
	"github.com/1zidorius/project-management-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	DBName             = "project-management"
	USERCOLLECTION     = "users"
	TASKSCOLECCTION    = "tasks"
	PROJECTSCOLLECTION = "projects"
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

func CreateUser(user models.User) (models.User, error) {
	user.Id = primitive.NewObjectID()
	_, err := db.Collection(USERCOLLECTION).InsertOne(context.Background(), user)
	if err != nil {
		return user, err
	}
	return user, nil
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

func DeleteUser(Id primitive.ObjectID) (int, error) {
	filter := bson.M{"_id": Id}
	deleteResult, err := db.Collection(USERCOLLECTION).DeleteOne(context.Background(), filter, nil)
	dCount := deleteResult.DeletedCount
	if err != nil {
		return int(dCount), err
	}
	return int(dCount), nil
}

func UpdateUser(id primitive.ObjectID, user models.User) (models.User, error) {
	var u models.User
	filter := bson.M{"_id": id}
	_ = db.Collection(USERCOLLECTION).FindOne(context.Background(), filter).Decode(&u)
	if user.Password != "" {
		u.Password = user.Password
	}
	if user.Email != "" {
		u.Email = user.Email
	}

	if user.Name != "" {
		u.Name = user.Name
	}

	if user.Surname != "" {
		u.Surname = user.Surname
	}

	update := bson.M{"$set": bson.M{"password": u.Password, "email": u.Email, "name": u.Name, "surname": u.Surname}}
	_, err := db.Collection(USERCOLLECTION).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return user, err
	}
	return u, nil
}

func GetUser(id primitive.ObjectID) (models.User, error) {
	var user models.User
	filter := bson.M{"_id": id}
	err := db.Collection(USERCOLLECTION).FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		log.Println(err)
		return user, err
	}
	return user, nil
}

//func CreateTask(task models.Task) (models.Task, error) {
//	task.Id = primitive.NewObjectID()
//	task.CreatedOn = time.Now()
//	task.AddedBy = addedBy
//	_, err := db.Collection(USERCOLLECTION).InsertOne(context.Background(), user)
//	if err != nil {
//		return user, err
//	}
//	return user, nil
//}

