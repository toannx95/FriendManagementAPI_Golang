package repository

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	email := "a@gmail.com"
	rows := sqlmock.NewRows([]string{"email"}).AddRow(email)

	query := `select email from user`
	mock.ExpectQuery(query).WillReturnRows(rows)

	// passes the mock to our code
	myDB := NewUserRepository(db)
	results := myDB.GetAllUsers()

	expected := []string{email}
	assert.Equal(t, expected, results)
}

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	email := "a@gmail.com"

	query := regexp.QuoteMeta("insert into user (email) values (?)")
	mock.ExpectPrepare(query).ExpectExec().WithArgs(email).WillReturnResult(sqlmock.NewResult(1, 1))

	// passes the mock to our code
	myDB := NewUserRepository(db)
	result := myDB.CreateUser(email)

	expected := true
	assert.Equal(t, expected, result)
}

func TestExistsByEmail(t *testing.T)  {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	email := "a@gmail.com"
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	query := `select id from user where email=?`
	mock.ExpectQuery(query).WithArgs(email).WillReturnRows(rows)

	// passes the mock to our code
	myDB := NewUserRepository(db)
	results := myDB.ExistsByEmail(email)

	expected := true
	assert.Equal(t, expected, results)
}

func TestFindUserIdByEmail(t *testing.T)  {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	email := "a@gmail.com"
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	query := `select id from user where email=?`
	mock.ExpectQuery(query).WithArgs(email).WillReturnRows(rows)

	// passes the mock to our code
	myDB := NewUserRepository(db)
	results := myDB.FindUserIdByEmail(email)

	expected := int64(1)
	assert.Equal(t, expected, results)
}

func TestFindByIds(t *testing.T)  {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ids := []int64{1,2,3}
	rows := sqlmock.NewRows([]string{"email"}).
					AddRow("a@gmail.com").
					AddRow("b@gmail.com").
					AddRow("c@gmail.com")

	strIds := make([]string, len(ids))
	for i, id := range ids {
		strIds[i] = strconv.FormatInt(id, 10)
	}

	stmt := `select x.email 
			from user x 
			where x.id in (%s);
			`
	query := fmt.Sprintf(stmt, strings.Join(strIds, ","))
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	// passes the mock to our code
	myDB := NewUserRepository(db)
	results := myDB.FindByIds(ids)

	expected := []string{"a@gmail.com", "b@gmail.com", "c@gmail.com"}
	assert.Equal(t, expected, results)
}