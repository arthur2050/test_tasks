package entity

type User struct {
	ID       int      `json:"id"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}
