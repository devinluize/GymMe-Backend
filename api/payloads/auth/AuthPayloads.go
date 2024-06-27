package payloads

type RegisterPayloads struct {
	Username    string `gorm:"UserName" json:"username"`
	Useremail   string `gorm:"UserEmail" json:"useremail"`
	Userpasword string `gorm:"UserPassword" json:"userpasword"`
	IsVIP       bool   `gorm:"IsVIP" json:"isVIP"`
}
