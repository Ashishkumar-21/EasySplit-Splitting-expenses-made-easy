package models

import "github.com/astaxie/beego/orm"

type Transaction struct {
	ID          int     `orm:"auto" json:"id"`
	PayerID     string  `orm:"size(100)" json:"payer_id"`
	PayeeID     string  `orm:"size(100)" json:"payee_id"`
	Amount      float64 `json:"amount"`
	Description string  `orm:"size(100)" json:"description"`
}

func init() {
	orm.RegisterModel(new(Transaction))
}
