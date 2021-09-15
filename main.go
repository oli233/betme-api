package main

import (
	"BetmeAPI/apis"
	"BetmeAPI/data"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

//const APIKEY = "9317664c724de4ae8eb40bbb38a21d18"
//const MONGOADDR = "mongodb://192.227.147.138:29017"

func main() {

	//command line get key and DBaddr
	var apik, mongoaddr string
	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		fmt.Println("arg is", args[i])
		switch args[i] {
		case "-k":
			apik = args[i+1]
			fmt.Println(apik)
			i++
			break
		case "-a":
			mongoaddr = args[i+1]
			fmt.Println("addr is", mongoaddr)
			i++
			break
		default:
			log.Fatalln("Args missing or here is no implemented command fot that.")
		}
	}

	//Init back-end
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	mongoSocket := data.Init(mongoaddr, &ctx)

	//Create an instance
	var client apis.Api
	client = &apis.ApiClient{Ctx: &ctx, DbScocket: mongoSocket, ApiKey: apik}

	clientResult := client.InitData() //change api here as needed
	if clientResult != 0 {
		fmt.Println("client finished with error code", clientResult)
	}
}
