package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	session, err := mgo.Dial("mongodb://username:password@localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	cQuery := session.DB("db").C("col")
	dataDB := make(bson.M)
	err = cQuery.Find(bson.M{}).Select(bson.M{"_id": 0}).One(&dataDB)
	if err != nil {
		fmt.Print("Mongo Query Error!!")
	}
	fmt.Println(dataDB)
	
	cInsert := session.DB("db").C("col")
	InsertData := `{"a":50, "b":"ABC"}`
	_ = cInsert.Insert(InsertData)
}
