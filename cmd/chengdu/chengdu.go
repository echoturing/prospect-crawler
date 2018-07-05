package main

import (
	"fmt"
	"time"

	"path/filepath"

	"os"

	"flag"

	"github.com/echoturing/buyhouse/db"
	"github.com/echoturing/buyhouse/etc"
	"github.com/echoturing/buyhouse/linkhome"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

var (
	districts = []string{"gaoxin7", "tianfuxinqu"}
)

func main() {
	configPath := flag.String("config", "./etc/config.yaml", "The config file path")
	flag.Parse()
	logger, err := zap.NewProduction()
	if err != nil {
		print(err.Error())
		os.Exit(1)
	}
	defer logger.Sync()
	filePath, err := filepath.Abs(*configPath)
	if err != nil {
		logger.Error("load config file path failed",
			zap.Error(err),
		)
		os.Exit(1)
	}
	cfg, err := etc.LoadConfigFromFile(filePath)
	if err != nil {
		logger.Error("load config failed",
			zap.Error(err),
		)
		os.Exit(1)
	}
	conn, err := db.NewConn(cfg)
	if err != nil {
		logger.Error("init db conn failed",
			zap.Error(err),
		)
		os.Exit(1)
	}
	log.Info("init mysql conn",
		zap.String("cfg", fmt.Sprintf("%#v", cfg)),
		zap.String("conn", fmt.Sprintf("%#v", conn)),
	)
	houseInfoDal := linkhome.NewHouseInfoDal(conn)
	timestamp := time.Now().Unix()
	for _, district := range districts {
		items := linkhome.CrawlDistrict(linkhome.Chengdu, district, 1)
		resultFile := fmt.Sprintf("result_%s_%d.txt", district, timestamp)
		linkhome.SaveHouseInfoToFile(resultFile, items)
		_, err = houseInfoDal.BatchCreateHouseInfo(items)
		if err != nil {
			logger.Error("save house info to db failed", zap.Error(err))
		}
	}
}
