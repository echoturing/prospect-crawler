package linkhome

import "fmt"

const (
	endpoint = `http://cd.lianjia.com`
)

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
