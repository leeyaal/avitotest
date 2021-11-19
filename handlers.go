package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	Id     int    `json:"id"`
	Amount int    `"json:"amount"`
	Status string `json:"status"`
}

type Transfer struct {
	Idgeneral  int `json:"idg"`
	Idreciever int `json:"idr"`
	Operation  string
	Time       string
	Sum        int
}

type Transit struct {
	Idgeneral  int `json:"idg"`
	Idreciever int `json:"idr"`
	Sum        int `json:"sum"`
}

//создание нового юзера...
func newuser(w http.ResponseWriter, r *http.Request) {

	// JSON...
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprint(w, "request can not be decoded")
	}
	if user.Id < 0 {
		fmt.Fprint(w, "Id should be more than 0")
	} else {
		// DB connection...
		connStr := "user=postgres password=qwerty7 dbname=avito sslmode=disable port=8080"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			fmt.Fprint(w, "DataBase connection failed")
		}
		defer db.Close()

		rows := db.QueryRow("select * from users where id = $1", user.Id)
		if user.Id == 0 {
			fmt.Fprint(w, "Id = 0 belongs to Avito, try another one")
		} else {
			err = rows.Scan(&user.Id)
			if err != sql.ErrNoRows {
				fmt.Fprint(w, "id exists, try another one")
			} else {

				_, err := db.Exec("insert into users (id, balance, status) values ($1, $2, $3)", user.Id, 0, "new")
				if err != nil {
					fmt.Fprint(w, "unexpected error")
				} else {
					fmt.Fprint(w, "new user was added")
				}

			}
		}
	}
}

//проверка баланса юзера...
func infobalance(w http.ResponseWriter, r *http.Request) {

	user := User{}
	usermass := []User{}

	// DB connection...
	connStr := "user=postgres password=qwerty7 dbname=avito sslmode=disable port=8080"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DataBase connection failed")
	}
	defer db.Close()

	id := r.URL.Query().Get("id")
	if err != nil {
		fmt.Fprint(w, "unexpected error")
	}

	rows := db.QueryRow("select * from users where id = $1", id)
	if err != nil {
		fmt.Fprint(w, "unexpected error")
	}

	error := rows.Scan(&user.Id, &user.Amount, &user.Status)
	if error != nil {
		if error == sql.ErrNoRows {
			fmt.Fprint(w, "id does not exist")
		}
	} else {

		usermass = append(usermass, user)

		for _, user := range usermass {
			sntnc, _ := fmt.Printf("%v \n", user.Amount)
			fmt.Fprint(w, sntnc)
		}
	}
}

//зачисление на счет, например, от работодателя...
func income(w http.ResponseWriter, r *http.Request) {

	time := time.Now().Format(time.RFC850)

	// JSON...
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprint(w, "request can not be decoded")
	}
	if user.Amount < 0 || user.Amount == 0 {
		fmt.Fprint(w, "Amount should be more than 0")
	} else {
		// DB connection...
		connStr := "user=postgres password=qwerty7 dbname=avito sslmode=disable port=8080"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			fmt.Fprint(w, "DataBase connection failed")
		}
		defer db.Close()

		row := db.QueryRow("select id from users where id = $1", user.Id)
		error := row.Scan(&user.Id)
		if error != nil {
			if error == sql.ErrNoRows {
				fmt.Fprint(w, "id does not exist")
			}
		} else {

			rows := db.QueryRow("select * from users where id = $1", user.Id)
			if err != nil {
				fmt.Fprint(w, "unexpected error")
			} else {
				_, err := db.Exec("update users set balance = balance + $1 where id = $2", user.Amount, user.Id)
				if err != nil {
					fmt.Fprint(w, "unexpected error")
				}
				_, error := db.Exec("update users set status = $1 where id = $2", "active", user.Id)
				if error != nil {
					fmt.Fprint(w, rows)
				}
				// запись операции в таблицу...
				_, error2 := db.Exec("insert into transfers (id_creator, id_reciever, operation, date, sum) values ($1, $2, $3, $4, $5)", 0, user.Id, "salary", time, user.Amount)
				if error2 != nil {
					fmt.Fprint(w, "unexpected error")
				}
				fmt.Fprint(w, "operation is completed")
			}
		}
	}
}

//списание средств...
func outcome(w http.ResponseWriter, r *http.Request) {

	time := time.Now().Format(time.RFC850)
	// JSON...
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprint(w, "request can not be decoded")
	}
	if user.Amount < 0 || user.Amount == 0 {
		fmt.Fprint(w, "Amount should be more than 0")
	} else {
		// DB connection...
		connStr := "user=postgres password=qwerty7 dbname=avito sslmode=disable port=8080"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			fmt.Fprint(w, "DataBase connection failed")
		}
		defer db.Close()

		rows := db.QueryRow("select id from users where id = $1", user.Id)
		error := rows.Scan(&user.Id)
		if error != nil {
			if error == sql.ErrNoRows {
				fmt.Fprint(w, "id does not exist")
			}
		} else {
			_, error1 := db.Exec("update users set balance = balance - $1 where id = $2", user.Amount, user.Id)
			if error != nil {
				fmt.Fprint(w, error1)
			}
			row := db.QueryRow("select id from users where balance < 0")
			err := row.Scan(&user.Id)
			if err == sql.ErrNoRows {

				_, error := db.Exec("update users set status = $1 where id = $2", "active", user.Id)
				if error != nil {
					fmt.Fprint(w, rows)
				}
				_, error2 := db.Exec("insert into transfers (id_creator, id_reciever, operation, date, sum) values ($1, $2, $3, $4, $5)", 0, user.Id, "outcome", time, user.Amount)
				if error2 != nil {
					fmt.Fprint(w, "smth happened")
				}
				fmt.Fprint(w, "operation is completed")
			} else {
				fmt.Fprint(w, "balance is less than sum of transfer")
				_, error1 := db.Exec("update users set balance = balance + $1 where id = $2", user.Amount, user.Id)
				if error != nil {
					fmt.Fprint(w, error1)
				}
			}
		}
	}
}

//перевод средств от юзера к юзеру...
func transit(w http.ResponseWriter, r *http.Request) {

	time := time.Now().Format(time.RFC850)

	// JSON...
	decoder := json.NewDecoder(r.Body)
	var transit Transit
	err := decoder.Decode(&transit)
	if err != nil {
		fmt.Fprint(w, "request can not be decoded")
	}
	//проверка Id...
	if transit.Idgeneral < 0 || transit.Idgeneral == 0 || transit.Idreciever < 0 || transit.Idreciever == 0 {
		fmt.Fprint(w, "id should be more than 0")
	} else {
		// DB connection...
		connStr := "user=postgres password=qwerty7 dbname=avito sslmode=disable port=8080"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			fmt.Fprint(w, "DataBase connection failed")
		}
		defer db.Close()

		//проверка наличия Id в БД...
		row1 := db.QueryRow("select id from users where id = $1", transit.Idgeneral)
		err1 := row1.Scan(&transit.Idgeneral)
		if err1 != nil {
			fmt.Fprint(w, "one of id or both do not exist")
		} else {
			row := db.QueryRow("select id from users where id = $1", transit.Idreciever)
			err := row.Scan(&transit.Idreciever)
			if err != nil {
				fmt.Fprint(w, "one of id or both do not exist")
			} else {
				//списание средств у id-инициатора операции...
				_, error1 := db.Exec("update users set balance = balance - $1 where id = $2", transit.Sum, transit.Idgeneral)
				if error1 != nil {
					fmt.Fprint(w, error1)
				}
				row := db.QueryRow("select id from users where balance < 0")
				err := row.Scan(&transit.Idgeneral)
				if err == sql.ErrNoRows {

					_, error := db.Exec("update users set status = $1 where id = $2", "active", transit.Idgeneral)
					if error != nil {
						fmt.Fprint(w, "unexpected error")
					}
					_, error2 := db.Exec("insert into transfers (id_creator, id_reciever, operation, date, sum) values ($1, $2, $3, $4, $5)", transit.Idgeneral, transit.Idreciever, "transfer", time, transit.Sum)
					if error2 != nil {
						fmt.Fprint(w, "smth happened")
					}
					// зачисление средств юзеру...
					_, errors := db.Exec("update users set balance = balance + $1 where id = $2", transit.Sum, transit.Idreciever)
					if errors != nil {
						fmt.Fprint(w, "unexpected error")
					}
					_, error3 := db.Exec("update users set status = $1 where id = $2", "active", transit.Idreciever)
					if error3 != nil {
						fmt.Fprint(w, err)
					}

					fmt.Fprint(w, "operation is completed")
				} else {
					fmt.Fprint(w, "balance is less than sum of transfer")
					_, error1 := db.Exec("update users set balance = balance + $1 where id = $2", transit.Sum, transit.Idgeneral)
					if error1 != nil {
						fmt.Fprint(w, error1)
					}
				}

			}

		}

	}

}
