package main

import (
	"fmt"
	"log"
	"time"
)

const (
	endpoint = `http://cd.lianjia.com`
)

var (
	blockNames = []string{"gaoxin7", "tianfuxinqu"}
)

func main() {
	timestamp := time.Now().Unix()
	for _, blockName := range blockNames {
		items := run(blockName, 1)
		resultFile := fmt.Sprintf("result_%s_%d.txt", blockName, timestamp)
		WriteHouseInfoToFile(resultFile, items)
	}
}

func run(blockName string, maxPage int) []HouseInfo {
	var allItems []HouseInfo
	var prevItems []HouseInfo
	for i := 1; i <= maxPage; i++ {
		requestURL := fmt.Sprintf("%s/ershoufang/%s/pg%d/", endpoint, blockName, i)
		log.Printf("crawling page %d from %s", i, requestURL)
		res, err := GetItemFromURL(requestURL)
		if err != nil {
			log.Printf("failed to crawl %s, err: %v", requestURL, err)
			continue
		}
		if NeedContinue(prevItems, res) {
			prevItems = res
			allItems = append(allItems, res...)
		} else {
			fmt.Println("发现重复的,停止")
			break
		}
		time.Sleep(time.Second * 5)
	}
	return allItems
}
