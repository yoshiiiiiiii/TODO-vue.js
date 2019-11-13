package entity

//テーブル情報
type Plan struct {
	ID    int    `gorm:"primary_key;not null"       json:"id"`
	Name  string `gorm:"type:varchar(200);not null" json:"name"`
	Cost  int    `gorm:"type:int"					json:"cost"`
	Priority int `gorm:"type:int"					json:"priority"`
	Memo  string `gorm:"type:varchar(400)"          json:"memo"`
	State int    `gorm:"not null"                   json:"state"`
}