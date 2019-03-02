package controller

import (
	"context"
	"log"
	"mini-services/models"
	"mini-services/pkg/e"
	"mini-services/pkg/mgo"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func GetTags(c *gin.Context) {
	var tag models.Tag
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	collection := mgo.MongoClient.Database("baz").Collection("tag")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result....
		log.Println(result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": tag})
}

func AddTag(c *gin.Context) {

	var tag models.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": e.INVALID_PARAMS, "Msg": e.GetMsg(e.INVALID_PARAMS), "error": err.Error()})
		return
	}

	data, err := bson.Marshal(tag)
	if err != nil {
		log.Println(`bson.Marshal`, err.Error())
	}

	collection := mgo.MongoClient.Database("baz").Collection("tag")
	_, err = collection.InsertOne(context.Background(), data)
	if err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": "ok"})
}

func EditTag(c *gin.Context) {
	// id := c.Param("id")
	var tag models.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": e.INVALID_PARAMS, "Msg": e.GetMsg(e.INVALID_PARAMS), "error": err.Error()})
		return
	}
	// filter := bson.M{"name": "name2"}
	// collection := mgo.MongoClient.Database("baz").Collection("tag")
	// err := collection.FindOne(nil, filter).Decode(&tag)
	// if err != nil {
	// 	log.Println(err)
	// }

	collection := mgo.MongoClient.Database("baz").Collection("tag")
	doc1 := bsonx.Doc{{"name", bsonx.String("123")}}
	update := bsonx.Doc{{"$inc", bsonx.Document(bsonx.Doc{{"state", bsonx.Int32(1)}})}}
	result := collection.FindOneAndUpdate(nil, doc1, update)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": result})
}

func DeleteTag(c *gin.Context) {
	collection := mgo.MongoClient.Database("baz").Collection("tag")
	doc1 := bsonx.Doc{{"name", bsonx.String("name2")}}
	collection.FindOneAndDelete(nil, doc1)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": "ok"})
}
