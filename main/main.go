package main

import (
	"fmt"
	"time"

	"github.com/echoturing/buyhouse/linkhome"
)

var (
	blockNames = []string{"gaoxin7", "tianfuxinqu"}
)

func main() {
	timestamp := time.Now().Unix()
	for _, blockName := range blockNames {
		items := linkhome.CrowlBlock(blockName, 1)
		resultFile := fmt.Sprintf("result_%s_%d.txt", blockName, timestamp)
		linkhome.SaveHouseInfo(resultFile, items)
	}
}
