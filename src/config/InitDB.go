package config

import (
	"database/sql"
	"fmt"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/joho/godotenv"
	//"os"
)

func OpenDB() (*sql.DB, error) {
	//env := godotenv.Load("./src/.env")
	//if env != nil {
	//	fmt.Print(env)
	//}

	//username := os.Getenv("DB_USER")
	//password := os.Getenv("DB_PASSWORD")
	//dbType := os.Getenv("DB_TYPE")
	//dbName := os.Getenv("DB_NAME")
	//dbHost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv("DB_PORT")

	username := "root"
	password := "123456"
	dbType := "mysql"
	dbName := "testdb"
	//dbHost := "localhost"
	dbHost := "mydb"
	dbPort := "3306"

	url := username + ":" + password + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Asia%2FKolkata"
	fmt.Println(url)
	conn, err := sql.Open(dbType, url)

	if err != nil {
		return nil, err
	} else {
		fmt.Println("DB connected!")
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}