package mysql

type Users struct {
	Address   string `db:"address"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Password  string `db:"password"`
	Username  string `db:"username"`
}
