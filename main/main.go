package main

import (
	"log"

	"../mylib/server"

	"github.com/line/line-bot-sdk-go/linebot"
)

type api = server.RoutAndHandler

// all of data that linebot needs
var channelSecret string = "6fb3aaf4b3867e72ce0e207e36a39efd"
var channelAccessToken string = "UB0GirpfktF+S9W04eLzMjh+1HbCzt3YxDwnbq1pVVikETNTXfWxhvAVhKPkTWjp5zws2d9wXqBfcW4NX+oZb8z8CNohTIyhSxeIpY2INu+idaD4BKG/+9vAf4djrYSoLJUrxrrA9WydULS9T0h3ywdB04t89/1O/w1cDnyilFU="
var httpReplyURL string = "https://api.line.me/v2/bot/message/reply"
var serverPort string = ":8000"

// linebot intance
var bot *linebot.Client

func main() {
	var err error
	apis := []api{api{Rout: "/callback", Func: HandleAllEvents}}

	bot, err = linebot.New(channelSecret, channelAccessToken)

	log.Println("bot:", bot, "\nerr:", err)

	server.StartServer(serverPort, apis)

}
