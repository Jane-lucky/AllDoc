package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func listHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "请求方式错误", http.StatusMethodNotAllowed)
	}
	//rows, err := db.Query("SELECT * FROM half_year")
	rows, err := db.Query("SELECT * FROM users")
	checkErr(err)
	var users []User
	var user User

	//var projects []Project
	//var project Project
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username)
		//err = rows.Scan(&project.Base, &project.Bd, &project.Sa, &project.Manager, &project.Province,
		//	&project.City, &project.Company, &project.Spec, &project.Workload, &project.Starting_date,
		//	&project.Ending_date, &project.Status, &project.Payment, &project.Remarks)
		checkErr(err)
		users = append(users, user)
		//projects = append(projects, project)
	}
	//t, err := template.New("list.html").ParseFiles("tpl/list.html")
	t, err := template.New("listcopy.html").ParseFiles("tpl/listcopy.html")
	checkErr(err)
	err = t.Execute(w, users)
	//err = t.Execute(w, projects)

	checkErr(err)
}

func addHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.New("create.html").ParseFiles("tpl/create.html")
		checkErr(err)
		t.Execute(w, nil)
	}
	if r.Method == "POST" {
		r.ParseForm()
		username := r.Form["username"][0]
		fmt.Println(r.Form["username"])
		stmt, err := db.Prepare("INSERT INTO users(username) VALUES(?)")

		if err != nil {

			panic(err)

		}

		_, err = stmt.Exec(username)

		if err != nil {

			panic(err)

		}

		http.Redirect(w, r, "/list", http.StatusMovedPermanently)

	}
}

func deleteHandle(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		http.Error(w, "请求放松错误", http.StatusMethodNotAllowed)
	}

	
	r.ParseForm()
	id := r.Form["id"][0]
	stmt, err := db.Prepare("DELETE FROM users HWERE id=?")
	checkErr(err)
	_, err = stmt.Exec(id)
	checkErr(err)
	http.Redirect(w, r, "/list", http.StatusMovedPermanently)

}

func updateHandle(w http.ResponseWriter, r *http.Request) {

}
