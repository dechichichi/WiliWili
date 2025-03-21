package model

type User struct {
	Username  string `json:"username" gorm:"primaryKey"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Signature string `json:"signature"`
	Uid       int64  `json:"uid" gorm:"primaryKey"`
}
type UserInfo struct {
	Username string `json:"username"`
	Id       int64  `json:"id"`
}

type UserProfile struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Signature string `json:"signature"`
}

type Image struct {
	Uid     int64  `json:"uid"`
	ImageID int64  `json:"image_id" gorm:"column:image_id"`
	Url     string `json:"url"`
}

func (User) TableName() string {
	return "users"
}

func (Image) TableName() string {
	return "images"
}
