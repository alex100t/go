package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost" // Точка входа
	port     = 5432
	user     = "pguser" // Пользователь БД
	dbname   = "db1"    // Идентификатор подключения
	password = "zzz"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}

func main() {
	//psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	_, err = db.Query("select 1")
	if err != nil {
		panic(err)
	}
}

/*
// Получение IAM-токена сервисного аккаунта, указанного в настройках функции
func getToken(ctx context.Context) string {
	resp, err := ycsdk.InstanceServiceAccount().IAMToken(ctx)
	if err != nil {
		panic(err)
	}
	return resp.IamToken
}


// Подключение к БД
func Handler(ctx context.Context) (*Response, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=require",
		host, port, user, getToken(ctx), dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	_, err = db.Query("select 1")
	if err != nil {
		panic(err)
	}

	return &Response{
		StatusCode: 200,
		Body:       "Successfully connected!",
	}, nil
}
*/
