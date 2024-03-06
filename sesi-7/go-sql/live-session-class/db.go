package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Person struct {
	ID int
	Name string
	Address string
}

func main () {
	fmt.Println("Hello World")

	connectionString := "host=localhost port=5432 user=postgres password=rajawali02 dbname=hacktiv sslmode=disable"
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic (err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to Database!!")

	// insert(db, "Budi", "Jakarta")
	// insert(db, "Andi", "Bandung")

	getAll(db)


}


func insert(db *sql.DB, name, address string) {
	query := "insert into person(name, address) values($1, $2) returning *"

	row := db.QueryRow(query, name, address)

	var p Person

	err := row.Scan(&p.ID, &p.Name, &p.Address)

	if err != nil {
		panic(err)
	}

	fmt.Println(p)
}

func getAll(db *sql.DB) {
	rows, err := db.Query("select * from person")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var people []Person

	for rows.Next() { // i < 10; i++
		var p Person

		// arr[i] = i
		err := rows.Scan(&p.ID, &p.Name, &p.Address)

		if err != nil {
			panic(err)
		}

		people = append(people, p)
	}

	fmt.Println(people)
}