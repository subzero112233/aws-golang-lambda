package mysql

import (
	"aws-golang-lambda/entity"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MysqlDB struct {
	db *sqlx.DB
}

func NewMysqlDB(db *sqlx.DB) *MysqlDB {
	return &MysqlDB{
		db: db,
	}
}

func (s *MysqlDB) GetUser(username string) (details entity.User, err error) {
	var users []Users
	err = s.db.Select(&users, "SELECT * FROM users WHERE username=?", username)
	if err != nil {
		return details, err
	}

	if len(users) > 0 {
		details = entity.User(users[0])
	}

	return details, nil
}

func (s *MysqlDB) AddUser(details entity.User) (err error) {
	user := Users(details)

	_, err = s.db.NamedExec("INSERT INTO users (username, password, first_name, last_name, address) VALUES (:username, :password, :first_name, :last_name, :address)", &user)
	return err
}
