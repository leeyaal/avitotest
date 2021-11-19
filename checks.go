package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func check(w http.ResponseWriter, r *http.Request) {

	user := User{}
	usermass := []User{}

	// DB connection...
	connStr := "user=postgres password=qwerty7 dbname=avito sslmode=disable port=8080"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DataBase connection failed")
	}
	defer db.Close()

	rows, err := db.Query("select * from users")
	if err != nil {
		fmt.Fprint(w, "unexpected error")
	}

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Amount, &user.Status)
		if err != nil {
			fmt.Println(err)
			continue
		}
		usermass = append(usermass, user)
	}
	for _, user := range usermass {
		sntnc, _ := fmt.Printf("%v | %v | %s \n", user.Id, user.Amount, user.Status)
		fmt.Fprint(w, sntnc)
	}

}

func history(w http.ResponseWriter, r *http.Request) {

	trans := Transfer{}
	transmass := []Transfer{}

	// DB connection...
	connStr := "user=postgres password=qwerty7 dbname=avito sslmode=disable port=8080"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DataBase connection failed")
	}
	defer db.Close()

	rows, err := db.Query("select * from transfers")
	if err != nil {
		fmt.Fprint(w, "unexpected error")
	}

	for rows.Next() {
		err := rows.Scan(&trans.Idgeneral, &trans.Idreciever, &trans.Operation, &trans.Time, &trans.Sum)
		if err != nil {
			fmt.Println(err)
			continue
		}
		transmass = append(transmass, trans)
	}
	for _, trans := range transmass {
		sntnc, _ := fmt.Printf("%v | %v | %s | %v | %v | \n", trans.Idgeneral, trans.Idreciever, trans.Operation, trans.Time, trans.Sum)
		fmt.Fprint(w, sntnc)
	}

}
