package MenuPayloads

import "time"

type CalendarInsertPayload struct {
	CalendarName     string    `gorm:"column:calendar_name" json:"calendar_name"`
	CalendarDate     time.Time `gorm:"column:calendar_date" json:"calendar_date"`
	UserId           int       `gorm:"column:user_id" json:"user_id"`
	CalendarTimeFrom time.Time `gorm:"column:calendar_time_from" json:"calendar_time_from"`
	CalendarTimeTo   time.Time `gorm:"column:calendar_time_to" json:"calendar_time_to"`
}
type CalendarUpdatePayload struct {
	CalendarId       int       `gorm:"column:calendar_id;primaryKey;not null" json:"calendar_id"`
	CalendarName     string    `gorm:"column:calendar_name" json:"calendar_name"`
	CalendarDate     time.Time `gorm:"column:calendar_date" json:"calendar_date"`
	UserId           int       `gorm:"column:user_id" json:"user_id"`
	CalendarTimeFrom time.Time `gorm:"column:calendar_time_from" json:"calendar_time_from"`
	CalendarTimeTo   time.Time `gorm:"column:calendar_time_to" json:"calendar_time_to"`
}
type CalendarGetByIdResponse struct {
	CalendarName     string    `gorm:"column:calendar_name" json:"calendar_name"`
	CalendarDate     time.Time `gorm:"column:calendar_date" json:"calendar_date"`
	UserId           int       `gorm:"column:user_id" json:"user_id"`
	CalendarTimeFrom time.Time `gorm:"column:calendar_time_from" json:"calendar_time_from"`
	CalendarTimeTo   time.Time `gorm:"column:calendar_time_to" json:"calendar_time_to"`
	CalendarId       int       `gorm:"column:calendar_id;primaryKey;not null" json:"calendar_id"`
}
