package entities

type RegisterPayloads struct {
	Username    string `gorm:"column:UserName" json:"username"`
	Useremail   string `gorm:"column:UserEmail" json:"useremail"`
	Userpasword string `gorm:"column:UserPassword" json:"userpasword"`
	IsVIP       bool   `gorm:"column:IsVIP" json:"isVIP"`
}

func (RegisterPayloads) TableName() string {
	return "Users"
}
