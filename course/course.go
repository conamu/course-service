package course

import (
	"database/sql"
	"errors"
)

func CreateCourse(request *Course, db *sql.DB) error {
	err := validateCourse(request)
	if err != nil {
		return err
	}
	createQuery := `INSERT INTO COURSES (INSTRUCTOR,DESCRIPTION,DIFFICULTY,FEE,CERTPATH,TITLE,SUBTITLE,ENLISTED,LIKES) VALUES (?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(createQuery, request.Instructor, request.Description, request.Difficulty, request.Fee, request.Certpath, request.Title, request.Subtitle, request.Enlisted, request.Likes)
	if err != nil {
		return err
	}

	return nil
}

func GetCourseByID(courseId string, db *sql.DB) (*Course, error) {
	var (
		id          string
		instructor  string
		description string
		difficulty  int
		fee         string
		certpath    string
		title       string
		subtitle    string
		enlisted    string
		likes       int
	)
	GetByIdQuery := `SELECT * FROM COURSES WHERE ID=?`

	rows, err := db.Query(GetByIdQuery, courseId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&id, &instructor, &description, &difficulty, &fee, &certpath, &title, &subtitle, &enlisted, &likes)
	}

	return &Course{
		Id:          id,
		Title:       title,
		Subtitle:    subtitle,
		Description: description,
		Instructor:  instructor,
		Difficulty:  difficulty,
		Fee:         fee,
		Certpath:    certpath,
		Enlisted:    enlisted,
		Likes:       likes,
	}, nil

}

func DeleteCourseById(id string, db *sql.DB) error {
	deleteQuery := `DELETE FROM COURSES WHERE ID=?;`
	_, err := db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCourseById(id string, course *Course, db *sql.DB) error {
	err := validateCourse(course)
	if err != nil {
		return err
	}
	updateQuery := `UPDATE COURSES SET INSTRUCTOR=?,DESCRIPTION=?,DIFFICULTY=?,FEE=?,CERTPATH=?,TITLE=?,SUBTITLE=?,ENLISTED=?,LIKES=? WHERE ID=?;`
	_, err = db.Exec(updateQuery, course.Instructor, course.Description, course.Difficulty, course.Fee, course.Certpath, course.Title, course.Subtitle, course.Enlisted, course.Likes, id)
	if err != nil {
		return err
	}
	return nil
}

func GetAllCourses(pageLength int, db *sql.DB) ([]CourseMin, error) {
	var (
		id         string
		title      string
		difficulty int
		fee        string
		likes      int
	)
	query := `SELECT ID,TITLE,DIFFICULTY,FEE,LIKES FROM COURSES LIMIT ?;`
	rows, err := db.Query(query, pageLength)
	if err != nil {
		return nil, err
	}

	response := []CourseMin{}
	for rows.Next() {
		err := rows.Scan(&id, &title, &difficulty, &fee, &likes)
		if err != nil {
			return nil, err
		}
		courseResult := CourseMin{
			Id:         id,
			Name:       title,
			Difficulty: difficulty,
			Fee:        fee,
			Likes:      likes,
		}
		response = append(response, courseResult)
	}
	return response, nil
}

func validateCourse(course *Course) error {
	if course.Fee == "" {
		return errors.New("course fields not filled in correctly")
	}
	if course.Title == "" {
		return errors.New("course fields not filled in correctly")
	}
	if course.Subtitle == "" {
		return errors.New("course fields not filled in correctly")
	}
	if course.Enlisted == "" {
		return errors.New("course fields not filled in correctly")
	}
	if course.Description == "" {
		return errors.New("course fields not filled in correctly")
	}
	if course.Instructor == "" {
		return errors.New("course fields not filled in correctly")
	}
	if course.Certpath == "" {
		return errors.New("course fields not filled in correctly")
	}
	return nil
}
