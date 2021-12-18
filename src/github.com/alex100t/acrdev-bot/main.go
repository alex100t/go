package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost" // Точка входа
	port     = 5432
	user     = "pguser"   // Пользователь БД
	dbname   = "postgres" // Идентификатор подключения
	password = "zzz"
)

type recinfo struct {
	description string
}

func main() {
	//psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	fmt.Println("Connect to database...")

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

	fmt.Println("connected")

	_, err = db.Query("select 1")
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection valid")

	fmt.Println("Get existed record:")

	acr := "ДРИБС"
	row := db.QueryRow("select description from acronyms where acronym = $1", acr)
	if err != nil {
		panic(err)
	}
	//defer row.Close()
	record := recinfo{}

	err = row.Scan(&record.description)
	if err != nil {
		panic(err)
	}

	fmt.Println(record.description)

	fmt.Println("Inser new record")
	result, err := db.Exec("insert into acronyms (acronym, description) values ($1, $2)",
		"ЗЗЗ", "Тестовый акроним")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // не поддерживается
	fmt.Println(result.RowsAffected()) // количество добавленных строк

	fmt.Println("Select new record:")
	acr = "ЗЗЗ"
	row = db.QueryRow("select description from acronyms where acronym = $1", acr)
	if err != nil {
		panic(err)
	}
	//defer row.Close()
	record = recinfo{}

	err = row.Scan(&record.description)
	if err != nil {
		panic(err)
	}

	fmt.Println(record.description)

	fmt.Println("Delete new record:")
	result, err = db.Exec("delete from acronyms where acronym = $1", acr)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество удаленных строк

}
