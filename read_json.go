package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result bson.M
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result)
}
