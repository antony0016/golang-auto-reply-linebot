package main

import (
	"fmt"
	"log"
	"os"

	"../mylib/server"

	"github.com/line/line-bot-sdk-go/linebot"
)

type api = server.RoutAndHandler

// all of data that linebot needs
var channelSecret string = ""
var channelAccessToken string = ""
var serverPort string = ""

// linebot intance
var bot *linebot.Client

func main() {
	var err error
	channelSecret = os.Getenv("channelSecret")
	channelAccessToken = os.Getenv("channelAccessToken")
	serverPort = fmt.Sprintf(":%s", os.Getenv("PORT"))
	apis := []api{api{Rout: "/callback", Func: HandleAllEvents}}

	bot, err = linebot.New(channelSecret, channelAccessToken)

	log.Println("bot:", bot, ",err:", err)

	server.StartServer(serverPort, apis)

}
