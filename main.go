package main

import (
	"course-service/api"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var initCoursesDBQuery = `CREATE TABLE IF NOT EXISTS COURSES 
(
ID INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
INSTRUCTOR VARCHAR(30) NOT NULL,
DESCRIPTION TEXT NOT NULL,
DIFFICULTY INTEGER NOT NULL,
FEE VARCHAR(10) NOT NULL,
CERTPATH VARCHAR(30) NOT NULL,
TITLE VARCHAR(30) NOT NULL,
SUBTITLE VARCHAR(40) NOT NULL,
ENLISTED TEXT NOT NULL,
LIKES INTEGER NOT NULL
);`

type Server struct {
	db     *sql.DB
	router *mux.Router
}

func NewServer(db *sql.DB, router *mux.Router) *Server {
	return &Server{
		db:     db,
		router: router,
	}
}

func main() {
	log.Println("Waiting for DB to be up...")
	time.Sleep(time.Second * 20)

	dbConfig := mysql.Config{
		User:   "kb-course",
		Passwd: "kb-course",
		Net:    "tcp",
		Addr:   "course-service-db:3306",
		DBName: "courses",
	}

	router := mux.NewRouter()

	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	_, err = db.Exec(initCoursesDBQuery)
	if err != nil {
		fmt.Println(err.Error())
	}

	server := NewServer(db, router)

	// Create a course
	server.router.HandleFunc("/create", checkAdminAuthHeader(api.CreateCourseHandlerFunc(server.db)))
	// Update a course
	server.router.HandleFunc("/update", checkAdminAuthHeader(api.UpdateCourseByIdHandlerFunc(server.db)))
	// Delete a course
	server.router.HandleFunc("/delete", checkAdminAuthHeader(api.DeleteCourseHandlerFunc(server.db)))
	// Get All courses IDs, Names, Difficulty Ratings
	server.router.HandleFunc("/courses", checkAuthHeader(api.GetAllCoursesHandlerFunc(server.db)))
	// Get one course by ID
	server.router.HandleFunc("/course", checkAuthHeader(api.GetCourseByIDHandlerFunc(server.db)))
	server.router.HandleFunc("/ping", api.Ping())

	log.Println("KB-Course-Service listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", server.router))
	defer db.Close()

}

// Auth header middleware
func checkAuthHeader(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		header := req.Header.Get("X-KBU-Auth")
		sessionHeader := req.Header.Get("X-KBU-Login")
		if header != "abcdefghijklmnopqrstuvwxyz" {
			res.WriteHeader(403)
			log.Println("Auth header does not match!")
			return
		}
		_, err := api.AuthenticateToken(sessionHeader)
		if err != nil {
			res.WriteHeader(401)
			log.Println("Login header does not match!")
			return
		}
		next(res, req)
	}
}

// Auth header middleware admin
func checkAdminAuthHeader(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		header := req.Header.Get("X-KBU-Auth")
		sessionHeader := req.Header.Get("X-KBU-Login")
		if header != "abcdefghijklmnopqrstuvwxyz" {
			res.WriteHeader(403)
			log.Println("Auth header does not match!")
			return
		}
		role, err := api.AuthenticateToken(sessionHeader)
		if err != nil {
			res.WriteHeader(401)
			log.Println("Login header does not match!")
			return
		}
		if role != "admin" {
			res.WriteHeader(401)
			log.Println("Role not sufficient!")
			return
		}
		next(res, req)
	}
}
