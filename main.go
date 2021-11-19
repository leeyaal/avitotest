package main

import (
	"fmt"
	"net/http"
)

func main() {

	// перевод средств...
	http.HandleFunc("/transit", transit)
	// списание средств
	http.HandleFunc("/outcome", outcome)
	// получение списка юзеров
	http.HandleFunc("/check", check)
	//добавление в список нового юзера
	http.HandleFunc("/newuser", newuser)
	//получение баланса конкретного юзера
	http.HandleFunc("/infobalance", infobalance)
	//начисление средств
	http.HandleFunc("/income", income)
	//получение списка операций
	http.HandleFunc("/history", history)

	fmt.Println("listen...")
	http.ListenAndServe(":7000", nil)
}
