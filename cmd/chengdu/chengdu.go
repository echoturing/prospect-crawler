package main

import (
	"fmt"

	"path/filepath"

	"os"

	"flag"

	"github.com/echoturing/buyhouse/db"
	"github.com/echoturing/buyhouse/ding_alert_service"
	"github.com/echoturing/buyhouse/etc"
	"github.com/echoturing/buyhouse/linkhome"
	"github.com/echoturing/buyhouse/logger"
	"go.uber.org/zap"
)

var (
	districts = []string{"gaoxin7", "tianfuxinqu"}
)

func main() {
	configPath := flag.String("config", "./etc/config.yaml", "The config file path")
	maxPage := flag.Int("max_page", 50, "the max page to crawl")
	flag.Parse()
	log := logger.GetLogger()
	defer log.Sync()
	filePath, err := filepath.Abs(*configPath)
	if err != nil {
		log.Error("load config file path failed",
			zap.Error(err),
		)
		os.Exit(1)
	}
	cfg, err := etc.LoadConfigFromFile(filePath)
	if err != nil {
		log.Error("load config failed",
			zap.Error(err),
		)
		os.Exit(1)
	}
	conn, err := db.NewConn(cfg)
	if err != nil {
		log.Error("init db conn failed",
			zap.Error(err),
		)
		os.Exit(1)
	}
	log.Info("init mysql conn",
		zap.String("cfg", fmt.Sprintf("%#v", cfg)),
		zap.String("conn", fmt.Sprintf("%#v", conn)),
	)
	houseInfoDal := linkhome.NewHouseInfoDal(conn)
	//timestamp := time.Now().Unix()
	var finalErr error
	if cfg.DingAlert.Url != "" {
		ding_alert_service.Alert("start crawling linkhome info", cfg.DingAlert.Url)
		defer ding_alert_service.Alert(fmt.Sprintf("crawling linkhome finished:%#v", finalErr), cfg.DingAlert.Url)
	}
	for _, district := range districts {
		items := linkhome.CrawlDistrict(linkhome.Chengdu, district, *maxPage)
		//resultFile := fmt.Sprintf("result_%s_%d.txt", district, timestamp)
		//linkhome.SaveHouseInfoToFile(resultFile, items)
		_, err = houseInfoDal.BatchCreateHouseInfo(items)
		if err != nil {
			finalErr = err
			log.Error("save house info to db failed", zap.Error(err))
		}
	}

}
