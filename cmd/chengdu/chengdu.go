package main

import (
	"fmt"
	"time"

	"github.com/echoturing/buyhouse/linkhome"
)

var (
	districts = []string{"gaoxin7", "tianfuxinqu"}
)

func main() {
	timestamp := time.Now().Unix()
	for _, name := range districts {
		items := linkhome.CrawlDistrict(linkhome.Chengdu, name, 100)
		resultFile := fmt.Sprintf("result_%s_%d.txt", name, timestamp)
		linkhome.SaveHouseInfo(resultFile, items)
	}
}
