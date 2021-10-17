package api

import (
	"course-service/course"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateCourseHandlerFunc(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
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
		createRequest := &course.Course{}
		err = json.Unmarshal(body, createRequest)
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
			w.WriteHeader(500)
		}
		w.WriteHeader(201)
	}
}

func GetCourseByIDHandlerFunc(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Get Course endpoint hit!")
		if r.Method != "GET" {
			w.WriteHeader(405)
			return
		}
		keys, ok := r.URL.Query()["courseId"]
		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'courseId is missing'")
			w.WriteHeader(400)
			return
		}
		courseId := keys[0]
		course, err := course.GetCourseByID(courseId, db)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(500)
		}
		if course.Title == "" {
			w.WriteHeader(404)
			return
		}
		data, err := json.MarshalIndent(course, "", " ")
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
		w.Write(data)
	}
}

func DeleteCourseHandlerFunc(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
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
		createRequest := &course.Course{}
		err = json.Unmarshal(body, createRequest)
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
			w.WriteHeader(500)
		}
		w.WriteHeader(201)
	}
}

func GetAllCoursesHandlerFunc(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
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
		createRequest := &course.Course{}
		err = json.Unmarshal(body, createRequest)
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
			w.WriteHeader(500)
		}
		w.WriteHeader(201)
	}
}
func UpdateCourseHandlerFunc(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
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
		createRequest := &course.Course{}
		err = json.Unmarshal(body, createRequest)
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
			w.WriteHeader(500)
		}
		w.WriteHeader(201)
	}
}
