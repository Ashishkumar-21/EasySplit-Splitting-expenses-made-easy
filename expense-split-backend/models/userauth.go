package models

type Userauth struct {
	User_id string `orm:"pk;size(5)" json:"user_id"` // Primary key
	Mobile  string `orm:"size(10);unique" json:"mobile"`       // Unique constraint
	Name    string `orm:"size(255)" json:"name"`      // Name field
}
// TableName overrides the default table name
func (u *Userauth) TableName() string {
	return "userauth"
}
