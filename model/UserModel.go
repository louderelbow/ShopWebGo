package model

type User struct {
	ID       int // GORM v2 必须用大写 ID 才识别为主键
	Phone    string
	Password string
	AddTime  int
	LastIp   string
	Email    string
	Status   int
}

func (User) TableName() string {
	return "user"
}
