package entities

const timerEntityTableName = "trx_timer"

type TimerEntity struct {
	TimerId   int    `gorm:"column:timer_id;primary_key;auto_increment;not null" json:"timer_id"`
	UserId    int    `gorm:"column:user_id;" json:"user_id"`
	TimerName string `gorm:"column:timer_name" json:"timer_name"`
	User      Users
	//RemindingHours   int                `gorm:"column:reminding_hours;" json:"reminding_hours"`
	//RemindingMinutes int                `gorm:"column:reminding_minutes;" json:"reminding_minutes"`
	Timer []TimerQueueEntity `gorm:"foreignKey:TimerId;references:TimerId" json:"timer"`
}

func (*TimerEntity) TableName() string {
	return timerEntityTableName
}