// Copyright 2017 Rough Industries LLC. All rights reserved.
package alexa

import (
	"bytes"
	alexa "github.com/mikeflynn/go-alexa/skillserver"
	"html/template"
	"github.com/golang/glog"
)

type News struct {
	Source   string
	Status   string
	Articles []Article
}
type Article struct {
	Description string
}

type HelloSession struct {
	AWSID string
}

type Hello struct {
}

type Data struct {
	Description string
}

var app = alexa.EchoApplication{ // Route
	AppID:    "amzn1.ask.skill.b73e3eee-8022-412c-8a7c-1e3474d4757c", // Echo App ID from Amazon Dashboard
	OnIntent: EchoIntentHandler,
	OnLaunch: EchoIntentHandler,
}

var Applications = map[string]interface{}{
	"/echo/helloworld": app,
}

func EchoIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	t := template.New("response")

	glog.Infoln(echoReq)
	var err error
	t, err = t.ParseFiles("./view/webcontent/alexa/templates/hello.ssml")
	//t, err = t.ParseFiles("./view/webcontent/alexa/templates/news.ssml")
	if err != nil {
		return
	}

	//f2c9e83ee8d24953b40e1bc4093c60e9
	//myClient := &http.Client{Timeout: 10 * time.Second}

	//r, err := myClient.Get("https://newsapi.org/v1/articles?source=cnn&sortBy=top&apiKey=f2c9e83ee8d24953b40e1bc4093c60e9")
	//if err != nil {
	//	return
	//}
	//defer r.Body.Close()

	//bodyBytes, _ := ioutil.ReadAll(r.Body)
	//glog.Infoln(string(bodyBytes))

	//jsonResult := new(News)

	//json.NewDecoder(r.Body).Decode(&jsonResult)
	//glog.Infoln(jsonResult.Articles[0].Description)

	reqSlotValue := echoReq.Request.Intent.Slots["UpdateText"].Value
	reqName := echoReq.Request.Intent.Name
	data := Data{}

	if reqName != "WhatIfIntent" {
		data = Data{
			//Description: jsonResult.Articles[0].Description,
			Description: "Hi! " + reqSlotValue,
		}
	} else {
		glog.Infoln(reqSlotValue)
		if reqSlotValue == "a million dollars" {
			data = Data{
				Description: "If I Had a million dollars. I'd buy you a fur coat, but not a real fur coat that's cruel.",
			}
		} else {
			data = Data{
				Description: "I don't know.",
			}
		}
	}

	var tpl bytes.Buffer
	if err := t.ExecuteTemplate(&tpl, "base", data); err != nil {
		glog.Error(err)
		return
	}

	result := tpl.String()
	glog.Infoln(result)

	sessionAtt := map[string]interface{}{
		"session": HelloSession{ // Route
			AWSID: echoReq.Session.User.UserID,
		},
	}

	echoResp.SessionAttributes = sessionAtt
	echoResp.OutputSpeech(result).Card("Hello World", "This is a test card.")
}
