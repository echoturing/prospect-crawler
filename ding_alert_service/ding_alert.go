package ding_alert_service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/echoturing/buyhouse/logger"
	"go.uber.org/zap"
)

var HttpClient = &http.Client{Timeout: time.Second * 10}

func Alert(content string, url string) {
	log := logger.GetLogger()
	msg := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": content,
		},
		"isAtAll": true,
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
