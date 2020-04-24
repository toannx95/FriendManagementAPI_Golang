package repository

//func TestCreateUser(t *testing.T)  {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	email := "a@gmail.com"
//
//	mock.ExpectBegin()
//	mock.ExpectExec("INSERT INTO user").WithArgs(email).WillReturnResult(sqlmock.NewResult(0, 0))
//	mock.ExpectCommit()
//
//	//userRepositoryMock := repository.UserRepositoryMock{}
//
//	// now we execute our method
//	//if userRepositoryMock.CreateUser(email) != true {
//	//	t.Errorf("error was not expected while updating stats: %s", err)
//	//}
//
//	// we make sure that all expectations were met
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//	}
//}