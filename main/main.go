package main

import (
	"fmt"
	"time"
)

const (
	endpoint = `http://cd.lianjia.com`
)

var (
	blockNames = []string{"gaoxin7", "tianfuxinqu"}
	resultFile = "%s_result_%s.txt"
)

func main() {
	for _, blockName := range blockNames {
		var allItems []HouseInfo
		var preItems []HouseInfo
		for i := 1; i < 100; i++ {
			requestURL := fmt.Sprintf("%s/ershoufang/%s/pg%d/", endpoint, blockName, i)
			time.Sleep(time.Second * 5)
			res, _ := GetItemFromUrl(requestURL)
			if NeedContinue(preItems, res) {
				preItems = res
				allItems = append(allItems, res...)
			} else {
				fmt.Println("发现重复的,停止")
				break
			}

		}
		resultFile := fmt.Sprintf(resultFile, time.Now(), blockName)
		WriteHouseInfoToFile(resultFile, allItems)
	}
}
