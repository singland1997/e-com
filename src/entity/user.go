package entity

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Roles    string `json:"roles"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}