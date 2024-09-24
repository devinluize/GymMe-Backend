package entities

type Users struct {
	UserId       int     `gorm:"column:user_id;primaryKey;not null" json:"user_id"`
	UserName     string  `gorm:"column:user_name" json:"user_name"`
	UserEmail    string  `gorm:"column:user_email" json:"user_email"`
	UserPassword string  `gorm:"column:user_password" json:"user_password"`
	IsVIP        bool    `gorm:"column:is_vip" json:"is_vip"`
	UserWeight   float64 `gorm:"column:user_weight" json:"user_weight"`
	UserHeight   float64 `gorm:"column:user_height" json:"user_height"`
	UserGender   string  `gorm:"column:user_gender" json:"user_gender"`
}

func (*Users) TableName() string {
	return "mtr_user"
}
