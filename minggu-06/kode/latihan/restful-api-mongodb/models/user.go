package user

import (
	"mongodb-restfulapi/conn"
	"time"
	
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID			bson.ObjectID `bson: "_id"`
	Name		string		  `bson: "name"`
	Address		string		  `bson: "address"`
	Age			int			  `bson: "age"`
	CreatedAt 	time.Time	  `bson: "created_at"`
	UpdatedAt	time.Time	  `bson: "updated_at"`
}

type Users []User

func UserInfo(id bson.ObjectID, userCollection string) (User, error) {
	// Get DB from Mongo Config
	db := conn.GetMongoDB()
	user := User{}
	err := db.C(userCollection).Find(bson.M{"_id": &id}).One(&user)
	return user, err
}