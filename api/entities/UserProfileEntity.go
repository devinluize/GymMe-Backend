package entities

type UserDetail struct {
	UserDetailId           int     `gorm:"column:user_detail_id;primaryKey;not null" json:"user_detail_id"`
	UserId                 int     `gorm:"column:user_id" json:"user_id"`
	UserWeight             float64 `gorm:"column:user_weight" json:"user_weight"`
	UserHeight             float64 `gorm:"column:user_height" json:"user_height"`
	UserGender             string  `gorm:"column:user_gender;size:1" json:"user_gender"`
	UserProfileDescription string  `gorm:"column:user_profile_description" json:"user_profile_description"`
}

//func (*UserDetail) TableName() string {
//	return "mtr_user_data_detail"
//}