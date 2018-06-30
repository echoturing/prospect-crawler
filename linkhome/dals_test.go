package linkhome

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"github.com/echoturing/buyhouse/etc"
	"github.com/echoturing/buyhouse/db"
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
	}
	info, err := c.HouseInfoDal.CreateHouseInfo(info)
	c.Nil(err)
	c.NotEqual(0, info.ID)
}

func TestConnectionTestSuite(t *testing.T) {
	suite.Run(t, new(DalTestSuit))
}
