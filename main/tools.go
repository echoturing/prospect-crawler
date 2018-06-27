package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func WriteHouseInfoToFile(filePath string, houseItems []HouseInfo) {
	file, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	for _, houseInfo := range houseItems {
		file.WriteString(houseInfo.unitPrice + "\t" + strconv.Itoa(houseInfo.totalPrice) + "\t" + houseInfo.title + "\t" + houseInfo.address + "\t" + houseInfo.followInfo)
		file.WriteString("\n")
	}
}

type Tag struct {
	subway  string
	taxfree string
	haskey  string
}

type HouseInfo struct {
	title      string
	detailURL  string
	address    string
	totalPrice int
	unitPrice  string
	followInfo string
	Tag
}

func NeedContinue(items1 []HouseInfo, items2 []HouseInfo) bool {
	if len(items1) != len(items2) {
		return true
	} else if len(items1) > 0 && len(items1) > 0 {
		if items1[0] == items2[0] {
			return false
		}
	}
	return true
}

func GetItemFromUrl(url string) ([]HouseInfo, error) {
	fmt.Println(url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var itemList []HouseInfo
	doc.Find(".sellListContent").Find(".info").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".title").Text()
		detailURL, _ := s.Find(".title").Find("a").Attr("href")
		address := s.Find(".address").Text()
		followInfo := s.Find(".followInfo").Text()
		subway := s.Find(".subway").Text()
		taxfree := s.Find(".taxfree").Text()
		haskey := s.Find(".haskey").Text()
		totalPrice, _ := strconv.Atoi(s.Find(".totalPrice").Find("span").Text())
		unitPrice := s.Find(".unitPrice").Text()
		tag := Tag{
			subway:  subway,
			taxfree: taxfree,
			haskey:  haskey,
		}
		houseInfo := HouseInfo{
			Tag:        tag,
			title:      title,
			detailURL:  detailURL,
			address:    address,
			followInfo: followInfo,
			totalPrice: totalPrice,
			unitPrice:  unitPrice,
		}
		fmt.Println(houseInfo)
		itemList = append(itemList, houseInfo)
	})
	return itemList, nil
}
