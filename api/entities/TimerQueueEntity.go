package entities

const timerQueueEntityTableName = "trx_timer_queue"

type TimerQueueEntity struct {
	TimerId int `gorm:"column:timer_id" json:"timer_id"`
	//Timer                      TimerEntity
	TimerQueueId               int    `gorm:"column:timer_queue_id;primary_key;AUTO_INCREMENT;not null" json:"timer_queue_id"`
	TimerQueueName             string `gorm:"column:timer_queue_name;not null" json:"timer_queue_name"`
	TimerQueueRemindingHour    int    `gorm:"column:timer_queue_reminding_hour" json:"timer_queue_reminding_hour"`
	TimerQueueRemindingMinutes int    `gorm:"column:timer_queue_reminding_minutes" json:"timer_queue_reminding_minutes"`
	TimerQueueLineNumber       int    `gorm:"column:timer_queue_line_number" json:"timer_queue_line_number"`
}

func (*TimerQueueEntity) TableName() string {
	return timerQueueEntityTableName
}
