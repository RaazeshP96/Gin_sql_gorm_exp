package Models

type User struct {
	Id      uint   `json:"id" gorm:"primary"`
	Name    string `json:"name" gorm:"varchar(50)"`
	Email   string `json:"email" gorm:"varchar(50)"`
	Phone   string `json:"phone" gorm:"varchar(50)"`
	Address string `json:"address" gorm:"varchar(50)"`
}

func (b *User) TableName() string {
	return "user"
}
