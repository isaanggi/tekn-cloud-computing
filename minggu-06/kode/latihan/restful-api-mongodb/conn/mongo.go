package conn

import (
	"fmt"
	"log"
	"os"
	
	"github.com/joho/godotenv"
	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env File")
	}
	host := os.Getenv("MONGO_HOST")
	dbName := os.Getenv("MONGO_DB_NAME")
	session, err := mgo.Dial(host)
	if err != nil {
		fmt.Println("Session err: ", err)
		os.Exit(2)
	}
	
	db = session.DB(dbName)
}

func GetMongoDB() *mgo.Database {
	return db
}