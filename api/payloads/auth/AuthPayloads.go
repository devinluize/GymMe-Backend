package payloads

type RegisterPayloads struct {
	Username        string `gorm:"UserName" json:"user_name"`
	Useremail       string `gorm:"UserEmail" json:"user_email"`
	Userpasword     string `gorm:"UserPassword" json:"user_password"`
	UserPhoneNumber string `gorm:"UserPhoneNumber" json:"user_phone_number"`
	UserGender      string `gorm:"UserGender" json:"user_gender"`
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
