package entities

import "time"

type SearchHistoryEntities struct {
	UserId     int       `json:"user_id"`
	SearchKey  string    `json:"search_key"`
	DateSearch time.Time `json:"date_search"`
}

func (*SearchHistoryEntities) TableName() string {
	return "trx_search_history"
}
