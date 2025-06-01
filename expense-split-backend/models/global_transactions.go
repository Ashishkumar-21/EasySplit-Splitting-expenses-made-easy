package models

import "github.com/astaxie/beego/orm"

type Global_transactions struct {
	ID          int     `orm:"column(ID);auto" json:"id"`
	PayerID     string  `orm:"column(PayerID);size(255)" json:"payer_id"`
	PayeeID     string  `orm:"column(PayeeID);size(255)" json:"payee_id"`
	Amount      float64 `orm:"column(Amount);digits(10);decimals(2)" json:"amount"`
	Description string  `orm:"column(Description);type(text)" json:"description"`
}

func (t *Global_transactions) TableName() string {
	return "Global_transactions" // Explicitly set table name
}

func init() {
	orm.RegisterModel(new(Global_transactions))
}
