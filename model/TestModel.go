package model

type Admin struct {
	Username string `form:"username" gorm:"primaryKey"`
	Password string `form:"password"`
}

func (Admin) TableName() string {
	return "admin"
}
