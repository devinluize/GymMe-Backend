package payloads

type RegisterPayloads struct {
	Username    string `gorm:"UserName" json:"username"`
	Useremail   string `gorm:"UserEmail" json:"useremail"`
	Userpasword string `gorm:"UserPassword" json:"userpasword"`
	IsVIP       bool   `gorm:"IsVIP" json:"isVIP"`
}

type LoginPaylods struct {
	Useremail   string `gorm:"UserEmail" json:"useremail"`
	Userpasword string `gorm:"UserPassword" json:"userpasword"`
}
type LoginRespons struct {
	UserName  string `json:"username"`
	UserEmail string `json:"userEmail"`
	IsVIP     bool   `json:"isVIP"`
	Token     string `json:"token"`
}
