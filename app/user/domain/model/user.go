package model

type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Signature string `json:"signature"`
	Uid       int64  `json:"uid"`
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
	Imageid int64  `json:"imageid"`
	Url     string `json:"url"`
}
