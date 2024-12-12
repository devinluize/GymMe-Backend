package MenuPayloads

import "time"

type WeightHistoryPayloads struct {
	//UserId         int       `gorm:"column:user_id" json:"user_id"`
	UserWeight     float64   `gorm:"column:user_weight" json:"user_weight"`
	UserWeightTime time.Time `gorm:"column:user_weight_time" json:"user_weight_time"`
}

type LastWeightResponse struct {
	UserWeightTime time.Time `json:"user_weight_time"`
	UserId         int       `json:"user_id"`
	UserWeight     float64   `json:"user_weight"`
}
