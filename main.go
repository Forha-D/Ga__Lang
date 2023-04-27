package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {

	var tmplt = template.Must(template.ParseFiles("template/login.html"))
	tmplt.Execute(w, nil)
}

func signup(w http.ResponseWriter, r *http.Request) {

	var tmplt = template.Must(template.ParseFiles("template/signup.html"))
	tmplt.Execute(w, nil)
}

func adduser(Username string, password string) bool {
	db, err := sql.Open("mysql", "root:kanon@1to5#@tcp(127.0.0.1:3306)/mockweb")
	if err != nil {
		panic(err)
	}
	add, err := db.Query("INSERT INTO users (Name,Password) VALUES (?,?)", (Username), (password))
	if err != nil {
		panic(err)
	}
	fmt.Println(add)
	defer db.Close()
	return true
}

func checkuser(Username string, password string) bool {
	db, err := sql.Open("mysql", "root:kanon@1to5#@tcp(127.0.0.1:3306)/mockweb")
	if err != nil {
		panic(err)
	}
	var exists bool
	var query string

	query = fmt.Sprintf("SELECT EXISTS (SELECT Name FROM Users WHERE Name ='%s' AND Password='%s')", (Username), (password))
	row := db.QueryRow(query).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	fmt.Println(row)
	return exists
}

func signupuser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var Username = r.Form["name"]
	var Password = r.Form["pass"]
	fmt.Println(Username, " ", Password)

	if adduser(Username[0], Password[0]) {
		var tmplt = template.Must(template.ParseFiles("template/index.html"))
		tmplt.Execute(w, nil)

	} else {
		var tmplt = template.Must(template.ParseFiles("template/error.html"))
		tmplt.Execute(w, nil)

	}

}

func loginuser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var Username = r.Form["name"]
	var Password = r.Form["pass"]
	fmt.Println(Username, " ", Password)

	if checkuser(Username[0], Password[0]) {
		var tmplt = template.Must(template.ParseFiles("template/index.html"))
		tmplt.Execute(w, nil)

	} else {
		var tmplt = template.Must(template.ParseFiles("template/error.html"))
		tmplt.Execute(w, nil)

	}

}

func main() {

	http.HandleFunc("/", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/loginuser", loginuser)
	http.HandleFunc("/signupuser", signupuser)
	http.ListenAndServe(":8000", nil)

}
