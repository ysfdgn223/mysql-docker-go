package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/mysqlgo/config"
	"example.com/mysqlgo/db"
	"example.com/mysqlgo/models"
)

var cnf config.Config

func main() {
	cnf = config.NewConfig()

	http.HandleFunc("/home", HomePage)

	log.Fatal(http.ListenAndServe(cnf.GetServerCnf().GetPort(), nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home page hit")
	// fmt.Fprintf(w, "hommeeeee")
	u, e := GetUsers()
	if e != nil {
		fmt.Println("my error = ", e)
		fmt.Fprintf(w, "str")
	}
	json.NewEncoder(w).Encode(u)
}

func GetUsers() ([]*models.User, error) {
	db := db.NewDBConnection(cnf.GetDBCnf()).GetDB()
	defer db.Close()

	results, err := db.Query("select * from users")
	defer results.Close()

	if err != nil {
		fmt.Println("in errr = ", err)
		return nil, err
	}

	var users []*models.User

	for results.Next() {
		var user models.User
		err = results.Scan(&user.ID, &user.Name, &user.Weight)
		if err != nil {
			log.Println("Unable to parse row:", err)
		}
		users = append(users, &user)
	}

	return users, nil
}
