package entity

type User struct {
	UserID string `json:"user_id" gorm:"primaryKey"`
	Name   string `json:"name"`
}

func (u *User) TableName() string {
	return "users"
}
