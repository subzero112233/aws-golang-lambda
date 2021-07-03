package mysql

import (
	"aws-golang-lambda/entity"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
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
	query := fmt.Sprintf("SELECT * FROM %s WHERE username = '%s'", usersTable, username)
	err = s.db.Select(&users, query)
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

	query := fmt.Sprintf("INSERT INTO %s (username, password, first_name, last_name, address) VALUES (:username, :password, :first_name, :last_name, :address)", usersTable)
	_, err = s.db.NamedExec(query, &user)
	return err
}
