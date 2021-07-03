package entity

type User struct {
	Address   string `json:"address"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Username  string `json:"username"`
}
