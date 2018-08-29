package main

import (
	"log"

	"wxrobot/wx"
	"github.com/larspensjo/config"
	"flag"
)

var TextReplyPath string
func init(){
	textReplyPath :=flag.String("textReplyPath","","")

	flag.Parse()

	TextReplyPath=*textReplyPath
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err0 := wx.Init(&wx.MessageHandler{
		TextHandler:textHandler,
	})
	if err0 != nil {
		log.Fatal(err0.Error())
	}


	wx.Listening()
}


func textHandler(msg *wx.Message){
	c, _ := config.ReadDefault(TextReplyPath)
	reply,err:=c.String("default", msg.Content)
	if err!=nil {
		log.Println("textHandler : get reply is err ! err is : ",err)
		return
	}
	if reply!=""{
		wx.SendMsg(msg.FromUserName,reply)
	}
}
