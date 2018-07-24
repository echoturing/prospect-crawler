package ding_alert_service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/echoturing/prospect-crawler/logger"
	"go.uber.org/zap"
)

var HttpClient = &http.Client{Timeout: time.Second * 10}

func Alert(content string, url string) {
	log := logger.GetLogger()
	msg := struct {
		MsgType string            `json:"msgtype"`
		Text    map[string]string `json:"text"`
		IsAtAll bool              `json:"isAtAll"`
	}{
		MsgType: "text",
		Text:    map[string]string{"content": content},
		IsAtAll: true,
	}
	msgString, _ := json.Marshal(msg)
	resp, err := HttpClient.Post(url, "application/json", strings.NewReader(string(msgString)))
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	log.Info("ding",
		zap.ByteString("response", response),
	)
}
