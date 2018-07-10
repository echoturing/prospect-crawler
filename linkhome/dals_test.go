package linkhome

import (
	"testing"
	"time"

	"github.com/echoturing/buyhouse/db"
	"github.com/echoturing/buyhouse/etc"
	"github.com/stretchr/testify/suite"
)

type DalTestSuit struct {
	suite.Suite
	HouseInfoDal
}

func (c *DalTestSuit) SetupSuite() {
	cfg, err := etc.LoadConfigFromFile("../etc/config.test.yaml")
	if err != nil {
		c.FailNow("load config failed", err)
	}
	conn, err := db.NewConn(cfg)
	if err != nil {
		c.FailNow("create mysql conn failed", err)
	}
	c.HouseInfoDal = HouseInfoDal{conn: conn}
}

func (c *DalTestSuit) TestCreateHouseInfo() {
	tm := time.Now()
	info := &HouseInfo{
		HouseCode:  "10000001",
		Title:      "house title",
		DetailURL:  "house detail url",
		Address:    "address",
		TotalPrice: 100,
		UnitPrice:  "18000",
		FollowInfo: "follow info",
		Subway:     "一号线",
		TaxFree:    "满5唯一",
		HasKey:     "有钥匙",
		City:       "chengdu",
		District:   "gaoxin7",
		CreatedAt:  tm,
	}
	info, err := c.HouseInfoDal.CreateHouseInfo(info)
	c.Nil(err)
	if err != nil {
		c.FailNow(err.Error())
	}
	c.NotEqual(0, info.ID)
	infoList, count, err := c.HouseInfoDal.GetHouseInfoListByCityAndDistrict("chengdu", "gaoxin7", 10, 0)
	c.Nil(err)
	c.Equal(int64(1), count)
	c.Equal(1, len(infoList))
	c.Equal(tm.UnixNano()/1000*1000, infoList[0].CreatedAt.UnixNano())
	info.ID = 0
	info.City = "BEIJING"
	_, err = c.HouseInfoDal.BatchCreateHouseInfo([]*HouseInfo{info})
	if err != nil {
		c.FailNow("batch create house info failed", err.Error())
	}
}

func TestConnectionTestSuite(t *testing.T) {
	suite.Run(t, new(DalTestSuit))
}
