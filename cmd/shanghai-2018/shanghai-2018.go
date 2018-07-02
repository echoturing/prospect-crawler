package main

import (
	"github.com/echoturing/buyhouse/linkhome"
)

var (
	districts = []string{
		"changning",
	}
)

func main() {
	// timestamp := time.Now().Unix()
	for _, name := range districts {
		items := linkhome.CrawlDistrictRent(linkhome.Shanghai, name, 1)
		// resultFile := fmt.Sprintf("result_%s_%d.txt", name, timestamp)
		// linkhome.SaveHouseInfoToFile(resultFile, items)
		_ = items
	}
}
