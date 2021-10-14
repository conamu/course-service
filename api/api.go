package api

import (
	"course-service/course"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func SignUpHandlerFunc(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Course Create endpoint hit!")
		if r.Method != "POST" {
			w.WriteHeader(405)
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(500)
			log.Println(err.Error())
			return
		}
		if r.ContentLength == 0 {
			w.WriteHeader(400)
			return
		}
		createRequest := &course.CreateRequest{}
		err = json.Unmarshal(body, createRequest)
		if err := authenticateToken(createRequest.Token); err != nil {
			w.WriteHeader(401)
			return
		}
		if err != nil {
			w.WriteHeader(500)
			log.Println(err.Error())
			return
		}
		if createRequest == nil {
			w.WriteHeader(400)
			return
		}
		log.Println(createRequest)

		err = course.CreateCourse(createRequest, db)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(401)
		}
		w.WriteHeader(201)
	}
}
