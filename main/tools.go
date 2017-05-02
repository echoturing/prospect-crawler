package main

import (
	"os"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

func WriteFile(filePath string, content interface{}) (int, error) {
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		fmt.Println(filePath, err)
		return -1, err
	}
	switch instance := content.(type) {
	case string:
		file.WriteString(instance)
	case []byte:
		file.Write(instance)
	default:
		return -9, nil
	}
	return 0, nil

}

func WriteHouseInfoToFile(filePath string, houseItems []HouseInfo) {
	file, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	for _, houseInfo := range houseItems {
		file.WriteString(houseInfo.unitPrise + "\t" + strconv.Itoa(houseInfo.totalPrise) + "\t" + houseInfo.title + "\t" + houseInfo.address + "\t" + houseInfo.followInfo)
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
	detailUrl  string
	address    string
	totalPrise int
	unitPrise  string
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
		detailUrl, _ := s.Find(".title").Find("a").Attr("href")
		address := s.Find(".address").Text()
		followInfo := s.Find(".followInfo").Text()
		subway := s.Find(".subway").Text()
		taxfree := s.Find(".taxfree").Text()
		haskey := s.Find(".haskey").Text()
		totalPrise, _ := strconv.Atoi(s.Find(".totalPrice").Find("span").Text())
		unitPrise := s.Find(".unitPrice").Text()
		tag := Tag{
			subway:  subway,
			taxfree: taxfree,
			haskey:  haskey,
		}
		houseInfo := HouseInfo{
			Tag:        tag,
			title:      title,
			detailUrl:  detailUrl,
			address:    address,
			followInfo: followInfo,
			totalPrise: totalPrise,
			unitPrise:  unitPrise,
		}
		fmt.Println(houseInfo)
		itemList = append(itemList, houseInfo)
	})
	return itemList, nil

}
