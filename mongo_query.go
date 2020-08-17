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
	cConfig := session.DB("iSolarWhiz").C("config_criteriaAlarm")
	dataDB := make(bson.M)
	err = cConfig.Find(bson.M{}).Select(bson.M{"_id": 0}).One(&dataDB)
	if err != nil {
		fmt.Print("Mongo Query Error!!")
	}
	fmt.Println(dataDB)
}
