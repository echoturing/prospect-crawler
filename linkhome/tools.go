package linkhome

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func SaveHouseInfoToFile(filePath string, houseItems []HouseInfo) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, houseInfo := range houseItems {
		fmt.Fprintf(file, "%s\n", houseInfo)
	}
	return nil
}

func getItemFromURL(url string) ([]HouseInfo, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	var itemList []HouseInfo
	doc.Find(".sellListContent").Find(".info").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".title").Text()
		detailURL, _ := s.Find(".title").Find("a").Attr("href")
		houseCode, _ := s.Find(".title").Find("a").Attr("data-housecode")
		address := s.Find(".address").Text()
		followInfo := s.Find(".followInfo").Text()
		subway := s.Find(".subway").Text()
		taxfree := s.Find(".taxfree").Text()
		haskey := s.Find(".haskey").Text()
		totalPrice, _ := strconv.Atoi(s.Find(".totalPrice").Find("span").Text())
		unitPrice := s.Find(".unitPrice").Text()
		houseInfo := HouseInfo{
			HouseCode:  houseCode,
			Title:      title,
			DetailURL:  detailURL,
			Address:    address,
			FollowInfo: followInfo,
			TotalPrice: totalPrice,
			UnitPrice:  unitPrice,
			Subway:     subway,
			TaxFree:    taxfree,
			HasKey:     haskey,
		}
		fmt.Println(houseInfo)
		itemList = append(itemList, houseInfo)
	})
	return itemList, nil
}

func getRentItemsFromURL(url string) ([]RentInfo, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	var items []RentInfo
	doc.Find("#house-lst").Find("li").Each(func(i int, s *goquery.Selection) {
		var ri RentInfo
		if err := ri.fromSelection(s); err != nil {
			log.Printf("failed to extract item: %v", err)
			return
		}
		items = append(items, ri)
		fmt.Println(ri)
	})
	return items, nil
}

func needContinue(items1 []HouseInfo, items2 []HouseInfo) bool {
	if len(items1) != len(items2) {
		return true
	} else if len(items1) > 0 && len(items1) > 0 {
		if items1[0] == items2[0] {
			return false
		}
	}
	return true
}

func CrawlDistrict(city City, district string, maxPage int) []HouseInfo {
	var allItems []HouseInfo
	var prevItems []HouseInfo
	for i := 1; i <= maxPage; i++ {
		requestURL := fmt.Sprintf("%s/%s/pg%d/", endpoint(city, SecondHand), district, i)
		log.Printf("crawling page %d from %s", i, requestURL)
		res, err := getItemFromURL(requestURL)
		if err != nil {
			log.Printf("failed to crawl %s, err: %v", requestURL, err)
			continue
		}
		if needContinue(prevItems, res) {
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

func CrawlDistrictRent(city City, district string, maxPage int) []RentInfo {
	var allItems []RentInfo
	for i := 1; i <= maxPage; i++ {
		requestURL := fmt.Sprintf("%s/%s/pg%d/", endpoint(city, Rent), district, i)
		log.Printf("crawling page %d from %s", i, requestURL)
		items, err := getRentItemsFromURL(requestURL)
		if err != nil {
			log.Printf("failed to crawl %s, err: %v", requestURL, err)
			continue
		}
		log.Printf("got %d items from page %d", len(items), i)
		allItems = append(allItems, items...)
		time.Sleep(time.Second * 5)
	}
	return allItems
}
