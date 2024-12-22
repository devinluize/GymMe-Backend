package payloads

type RegisterPayloads struct {
	Username        string  `gorm:"UserName" json:"user_name"`
	Useremail       string  `gorm:"UserEmail" json:"user_email"`
	Userpasword     string  `gorm:"UserPassword" json:"user_password"`
	UserPhoneNumber string  `gorm:"UserPhoneNumber" json:"user_phone_number"`
	UserGender      string  `gorm:"UserGender" json:"user_gender"`
	UserHeight      float64 `gorm:"UserHeight" json:"user_height"`
	UserWeight      float64 `gorm:"UserWeight" json:"user_weight"`
}

type LoginPaylods struct {
	Useremail   string `gorm:"UserEmail" json:"user_email"`
	Userpasword string `gorm:"UserPassword" json:"user_password"`
}
type LoginRespons struct {
	UserName  string `json:"username"`
	UserEmail string `json:"userEmail"`
	IsVIP     bool   `json:"isVIP"`
	Token     string `json:"token"`
}

type UserBmiResponse struct {
	UserId     int     `json:"user_id"`
	Bmi        float64 `json:"bmi"`
	UserWeight float64 `json:"user_weight"`
	UserHeight float64 `json:"user_height"`
}
