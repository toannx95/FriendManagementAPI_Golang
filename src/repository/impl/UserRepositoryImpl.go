package impl

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func (u UserRepositoryImpl) GetAllUsers() []string {
	result, err := u.DB.Query("select email from user")
	if err != nil {
		panic(err.Error())
	}

	emails:= []string{}
	for result.Next() {
		var email string
		err = result.Scan(&email)
		if err != nil {
			panic(err.Error())
		}
		emails = append(emails, email)
	}
	return emails
}

func (u UserRepositoryImpl) CreateUser(email string) bool {
	query, err := u.DB.Prepare(`insert into user (email) values (?)`)
	if err != nil {
		return false
	}
	query.Exec(email)
	return true
}

func (u UserRepositoryImpl) ExistsByEmail(email string) bool {
	var id int
	err := u.DB.QueryRow(`select id from user where email=?`, email).Scan(&id)
	if err != nil {
		return false
	}
	return true
}

func (u UserRepositoryImpl) FindUserIdByEmail(email string) int64 {
	var id int64
	err := u.DB.QueryRow("select id from user where email=?", email).Scan(&id)
	if err != nil {
		return -1
	}
	return id
}

func (f UserRepositoryImpl) FindByIds(ids []int64) []string {
	strIds := make([]string, len(ids))
	for i, id := range ids {
		strIds[i] = strconv.FormatInt(id, 10)
	}

	stmt := `select x.email
			from user x
			where x.id in (%s);
			`
	query := fmt.Sprintf(stmt, strings.Join(strIds, ","))
	results, err := f.DB.Query(query)
	if err != nil {
		panic(err)
	}

	emails := []string{}
	for results.Next() {
		var email string
		results.Scan(&email)
		emails = append(emails, email)
	}
	return emails
}