package entities

import "time"

type EquipmentSearchHistoryEntities struct {
	UserId     int       `json:"user_id"`
	SearchKey  string    `json:"search_key"`
	DateSearch time.Time `json:"date_search"`
}

func (*EquipmentSearchHistoryEntities) TableName() string {
	return "trx_equipment_search_history"
}
