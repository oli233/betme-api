package apis

import (
	"BetmeAPI/data"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"net/http"
)

type Api interface {
	InitData() int
	FetchData() int
}

type ApiClient struct {
	DbScocket *data.MongoSocket
	ApiKey    string
	Ctx       *context.Context
}

func (c *ApiClient) InitData() int {
	resp, err := http.Get("https://api.the-odds-api.com/v3/sports?apiKey=" + c.ApiKey)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//decoding to json
	var jsonResult data.Sports
	err = json.Unmarshal([]byte(string(body)), &jsonResult)
	if err != nil {
		return 1
	}

	// Check the connection
	collection, _ := c.DbScocket.GetCollection("sports", "gp3")
	collection1, _ := c.DbScocket.GetCollection("sports", "gp4")
	//database := c.DbScocket.Database("sports")
	//collection := database.Collection("gp3")

	for i := 0; i < len(jsonResult.DataArray); i++ {
		c.DbScocket.InsertData(collection, bson.M{"key": jsonResult.DataArray[i].Key}, jsonResult.DataArray[i])

		// Request odds & store
		resp, err := http.Get("https://api.the-odds-api.com/v3/odds/?sport=" + jsonResult.DataArray[i].Key + "&region=uk&apiKey=" + c.ApiKey)
		if err != nil {
			log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		//decoding to json
		var jsonOdds data.Odds
		err = json.Unmarshal([]byte(string(body)), &jsonOdds)
		if err != nil {
			return 1
		}

		for j := 0; j < len(jsonOdds.OddsDataArray); j++ {
			c.DbScocket.InsertData(collection1, bson.M{"id": jsonOdds.OddsDataArray[j].Id}, jsonOdds.OddsDataArray[j])
		}
	}

	return 0
}

func (c *ApiClient) FetchData() int {

	var sportsResult []bson.M
	var oddsResult []bson.M

	c.DbScocket.FetchAll(&sportsResult, &oddsResult)

	fmt.Println(sportsResult)
	fmt.Println(oddsResult)

	return 0
}
