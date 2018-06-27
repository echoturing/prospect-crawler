package linkhome

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// City name
type City string

// List of Supported cities
const (
	Chengdu  City = `cd`
	Shanghai City = `sh`
)

// Category name
type Category string

// List of Supported categories
const (
	Rent       Category = `zufang`
	SecondHand Category = `ershoufang`
)

func endpoint(city City, cat Category) string {
	return fmt.Sprintf("https://%s.lianjia.com/%s", city, cat)
}

type Tags struct {
	Subway  string
	TaxFree string
	HasKey  string
}

type HouseInfo struct {
	Title      string `json:"title"`
	DetailURL  string `json:"detail_url"`
	Address    string `json:"address"`
	TotalPrice int    `json:"total_price"`
	UnitPrice  string `json:"unit_price"`
	FollowInfo string `json:"follow_info"`
	Tags       Tags   `json:"tags"`
}

func (hi HouseInfo) String() string {
	return fmt.Sprintf("%s %d %s %s %s", hi.UnitPrice, hi.TotalPrice, hi.Title, hi.Address, hi.FollowInfo)
}

type RentInfo struct {
	DataID        string // TODO: int64
	DataHouseCode string // TODO: int64
	// TODO: image URL

	Where string
	Price string
}

func (ri *RentInfo) fromSelection(s *goquery.Selection) error {
	ri.DataID, _ = s.Attr("data-id")
	ri.DataHouseCode, _ = s.Attr("data-housecode")
	ri.Where = s.Find(".where").Text()
	ri.Price = s.Find(".price").Text()
	return nil
}

func (ri RentInfo) String() string {
	return fmt.Sprintf("%s    %-16s %s", ri.DataID, ri.Price, ri.Where)
}
