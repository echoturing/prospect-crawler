package linkhome

import (
	"github.com/gocraft/dbr"
)

const TableHouseInfo = "house_info"

var TableHouseColumns = []string{
	"house_code",
	"title",
	"detail_url",
	"address",
	"total_price",
	"unit_price",
	"follow_info",
	"subway",
	"tax_free",
	"has_key",
}

type HouseInfoDal struct {
	conn *dbr.Connection
}

func (h *HouseInfoDal) CreateHouseInfo(info *HouseInfo) (*HouseInfo, error) {
	session := h.conn.NewSession(nil)
	result, err := session.InsertInto(TableHouseInfo).
		Columns(TableHouseColumns...).
		Record(info).
		Exec()
	if err != nil {
		return nil, err
	}
	info.ID, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (h *HouseInfoDal) BatchCreateHouseInfo(infoList []*HouseInfo) (int64, error) {
	session := h.conn.NewSession(nil)
	tx, err := session.Begin()
	if err != nil {
		return 0, err
	}
	var count int64
	defer tx.RollbackUnlessCommitted()
	for info := range infoList {
		result, err := tx.InsertInto(TableHouseInfo).
			Record(info).
			Columns(TableHouseColumns...).
			Exec()
		if err != nil {
			tx.Rollback()
			return 0, err
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			tx.Rollback()
			return 0, err
		}
		count += rowsAffected
	}

	return count, nil
}
