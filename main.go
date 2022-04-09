package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var courses []Course

type Course struct {
	ID   string `json:"id"`
	Name string `json:"course_name"`
}

func generateCourses() {
	course1 := Course{
		ID:   "1",
		Name: "Full Cycle",
	}

	course2 := Course{
		ID:   "2",
		Name: "Bonus FullCycle",
	}

	courses = append(courses, course1, course2)
}

func main() {
	generateCourses()
	http.HandleFunc("/courses", listCourses)
	http.ListenAndServe(":8081", nil)
}

func listCourses(w http.ResponseWriter, r *http.Request) {
	jsonCourses, _ := json.Marshal(courses)
	w.Write([]byte(jsonCourses))
}

func persistCourse() error {
	db, err := sql.Open("sqlite3", "teste.db")
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("insert into courses values ($1, $2)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec("abc", "curso de cloud")
	if err != nil {
		return err
	}
	return err
}
