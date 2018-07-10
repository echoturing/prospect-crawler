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
	"city",
	"district",
	"created_at",
}

type HouseInfoDal struct {
	conn *dbr.Connection
}

func NewHouseInfoDal(conn *dbr.Connection) *HouseInfoDal {
	return &HouseInfoDal{
		conn: conn,
	}
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
	defer session.Close()
	tx, err := session.Begin()
	if err != nil {
		return 0, err
	}
	var count int64
	defer tx.RollbackUnlessCommitted()
	for _, info := range infoList {
		result, err := tx.InsertInto(TableHouseInfo).
			Columns(TableHouseColumns...).
			Record(info).
			Exec()
		if err != nil {
			return 0, err
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return 0, err
		}
		count += rowsAffected
	}
	tx.Commit()
	return count, nil
}

func (h *HouseInfoDal) GetHouseInfoListByCityAndDistrict(city, district string, limit, offset uint64) ([]*HouseInfo, int64, error) {
	session := h.conn.NewSession(nil)
	defer session.Close()
	var infoList []*HouseInfo
	count, err := session.Select(TableHouseColumns...).
		From(TableHouseInfo).
		Where(dbr.And(dbr.Eq("city", city), dbr.Eq("district", district))).
		Limit(limit).
		Offset(offset).
		Load(&infoList)
	if err != nil {
		return nil, -1, err
	}
	return infoList, int64(count), nil
}
