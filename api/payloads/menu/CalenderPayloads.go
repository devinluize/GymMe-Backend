package MenuPayloads

import "time"

type CalenderInsertPayload struct {
	CalenderName     string    `gorm:"column:calender_name" json:"calender_name"`
	CalenderDate     time.Time `gorm:"column:calender_date" json:"calender_date"`
	UserId           int       `gorm:"column:user_id" json:"user_id"`
	CalenderTimeFrom time.Time `gorm:"column:calender_time_from" json:"calender_time_from"`
	CalenderTimeTo   time.Time `gorm:"column:calender_time_to" json:"calender_time_to"`
}
type CalenderUpdatePayload struct {
	CalenderId       int       `gorm:"column:calender_id;primaryKey;not null" json:"calender_id"`
	CalenderName     string    `gorm:"column:calender_name" json:"calender_name"`
	CalenderDate     time.Time `gorm:"column:calender_date" json:"calender_date"`
	UserId           int       `gorm:"column:user_id" json:"user_id"`
	CalenderTimeFrom time.Time `gorm:"column:calender_time_from" json:"calender_time_from"`
	CalenderTimeTo   time.Time `gorm:"column:calender_time_to" json:"calender_time_to"`
}
type CalendarGetByIdResponse struct {
	CalenderName     string    `gorm:"column:calender_name" json:"calender_name"`
	CalenderDate     time.Time `gorm:"column:calender_date" json:"calender_date"`
	UserId           int       `gorm:"column:user_id" json:"user_id"`
	CalenderTimeFrom time.Time `gorm:"column:calender_time_from" json:"calender_time_from"`
	CalenderTimeTo   time.Time `gorm:"column:calender_time_to" json:"calender_time_to"`
}
