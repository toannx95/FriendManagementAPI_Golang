package repository

import (
	"fmt"
	"friend/entity"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestCreateRelationship(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	relationship := entity.Relationship{FirstEmailId: 1, SecondEmailId: 2, Status: 0}

	query := regexp.QuoteMeta("insert into relationship (first_email_id, second_email_id, status) values (?, ?, ?);")
	mock.ExpectPrepare(query).
			ExpectExec().
			WithArgs(relationship.FirstEmailId, relationship.SecondEmailId, relationship.Status).
			WillReturnResult(sqlmock.NewResult(1, 1))

	// passes the mock to our code
	myDB := NewRelationshipRepository(db)
	result := myDB.CreateRelationship(relationship)

	expected := true
	assert.Equal(t, expected, result)
}

func TestFindByTwoEmailIdsAndStatus(t *testing.T)  {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	firstEmailId := int64(1)
	secondEmailId := int64(2)
	status := []int64{0, 1}
	rows := sqlmock.
		NewRows([]string{"id", "first_email_id", "second_email_id", "status"}).
		AddRow(1, 1, 2, 0).
		AddRow(2, 1, 3, 1)

	strStatusIds := make([]string, len(status))
	for i, id := range status {
		strStatusIds[i] = strconv.FormatInt(id, 10)
	}

	stmt := `select x.*
			from relationship x
			where x.first_email_id in (%s, %s)
			and x.second_email_id in (%s, %s)
			and x.status in (%s);
			`
	query := fmt.Sprintf(
		stmt,
		strconv.FormatInt(firstEmailId, 10),
		strconv.FormatInt(secondEmailId, 10),
		strconv.FormatInt(firstEmailId, 10),
		strconv.FormatInt(secondEmailId, 10),
		strings.Join(strStatusIds, ","))
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	// passes the mock to our code
	myDB := NewRelationshipRepository(db)
	results := myDB.FindByTwoEmailIdsAndStatus(firstEmailId, secondEmailId, status)

	rela1 := entity.Relationship{1, 1, 2, 0}
	rela2 := entity.Relationship{2, 1, 3, 1}
	expected := []entity.Relationship{rela1, rela2}
	assert.Equal(t, expected, results)
}

func TestFindByEmailIdAndStatus(t *testing.T)  {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	emailId := int64(1)
	status := []int64{0, 1}
	rows := sqlmock.
		NewRows([]string{"id", "first_email_id", "second_email_id", "status"}).
		AddRow(1, 1, 2, 0).
		AddRow(2, 1, 3, 1)

	strStatusIds := make([]string, len(status))
	for i, id := range status {
		strStatusIds[i] = strconv.FormatInt(id, 10)
	}

	stmt := `select x.*
			from relationship x
			where (x.first_email_id = %s
			or x.second_email_id = %s) 
			and x.status in (%s);
			`
	query := fmt.Sprintf(
		stmt,
		strconv.FormatInt(emailId, 10),
		strconv.FormatInt(emailId, 10),
		strings.Join(strStatusIds, ","))
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	// passes the mock to our code
	myDB := NewRelationshipRepository(db)
	results := myDB.FindByEmailIdAndStatus(emailId, status)

	rela1 := entity.Relationship{1, 1, 2, 0}
	rela2 := entity.Relationship{2, 1, 3, 1}
	expected := []entity.Relationship{rela1, rela2}
	assert.Equal(t, expected, results)
}

func TestFindByFirstOrSecondEmailIdAndStatus(t *testing.T)  {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	firstEmailId := int64(1)
	secondEmailId := int64(2)
	status := []int64{0, 1}
	rows := sqlmock.
		NewRows([]string{"id", "first_email_id", "second_email_id", "status"}).
		AddRow(1, 1, 2, 0).
		AddRow(2, 1, 3, 1)

	strStatusIds := make([]string, len(status))
	for i, id := range status {
		strStatusIds[i] = strconv.FormatInt(id, 10)
	}

	stmt := `select x.*
			from relationship x
			where x.first_email_id in (%s, %s)
			or x.second_email_id in (%s, %s)
			and x.status in (%s);
			`
	query := fmt.Sprintf(
		stmt,
		strconv.FormatInt(firstEmailId, 10),
		strconv.FormatInt(secondEmailId, 10),
		strconv.FormatInt(firstEmailId, 10),
		strconv.FormatInt(secondEmailId, 10),
		strings.Join(strStatusIds, ","))
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	// passes the mock to our code
	myDB := NewRelationshipRepository(db)
	results := myDB.FindByFirstOrSecondEmailIdAndStatus(firstEmailId, secondEmailId, status)

	rela1 := entity.Relationship{1, 1, 2, 0}
	rela2 := entity.Relationship{2, 1, 3, 1}
	expected := []entity.Relationship{rela1, rela2}
	assert.Equal(t, expected, results)
}

func TestFindSubscribersByEmailId(t *testing.T)  {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	emailId := int64(1)
	rows := sqlmock.
		NewRows([]string{"x.first_email_id"}).
		AddRow(2).
		AddRow(3).
		AddRow(4)

	query := `select x.first_email_id 
			from relationship x 
			where x.second_email_id = ? 
			and x.status = 1;
			`
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(emailId).WillReturnRows(rows)

	// passes the mock to our code
	myDB := NewRelationshipRepository(db)
	results := myDB.FindSubscribersByEmailId(emailId)

	expected := []int64{2, 3, 4}
	assert.Equal(t, expected, results)
}