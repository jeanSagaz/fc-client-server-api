package entity

import pkgEntity "github.com/jeanSagaz/server/pkg/entity"

type UsdBrl struct {
	Id         pkgEntity.ID `gorm:"size:40;column:Id;primary_key;not null"`
	Code       string       `gorm:"size:10;column:Code;"`
	Codein     string       `gorm:"size:10;column:Codein;"`
	Name       string       `gorm:"size:50;column:Name;"`
	High       string       `gorm:"size:10;column:High;"`
	Low        string       `gorm:"size:10;column:Low;"`
	VarBid     string       `gorm:"size:10;column:VarBid;"`
	PctChange  string       `gorm:"size:10;column:PctChange;"`
	Bid        string       `gorm:"size:10;column:Bid;"`
	Ask        string       `gorm:"size:10;column:Ask;"`
	Timestamp  string       `gorm:"size:15;column:Timestamp;"`
	CreateDate string       `gorm:"size:25;column:CreateDate;"`
}

func NewUsdBrl(code string,
	codein string,
	name string,
	high string,
	low string,
	varBid string,
	pctChange string,
	bid string,
	ask string,
	timestamp string,
	createDate string,
) *UsdBrl {
	return &UsdBrl{
		Id:         pkgEntity.NewId(),
		Code:       code,
		Codein:     codein,
		Name:       name,
		High:       high,
		Low:        low,
		VarBid:     varBid,
		PctChange:  pctChange,
		Bid:        bid,
		Ask:        ask,
		Timestamp:  timestamp,
		CreateDate: createDate,
	}
}
