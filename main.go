package main

import (
	"flag"
	"log"

	"wxrobot/wx"
)


func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := wx.Init(&wx.MessageHandler{
		TextHandler:textHandler,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	wx.Listening()
}

func textHandler(msg *wx.Message){
	wx.SendMsg(msg.FromUserName,"haha")
}
