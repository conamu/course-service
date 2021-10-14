package course

import (
	"database/sql"
)

func CreateCourse(request *CreateRequest, db *sql.DB) error {

	createQuery := `INSERT INTO COURSES (INSTRUCTOR,DESCRIPTION,DIFFICULTY,FEE,CERTPATH,TITLE,SUBTITLE,ENLISTED,LIKES) VALUES (?,?,?,?,?,?,?,?,?);`
	_, err := db.Exec(createQuery, request.Instructor, request.Description, request.Difficulty, request.Fee, request.Certpath, request.Title, request.Subtitle, request.Enlisted, request.Likes)
	if err != nil {
		return err
	}

	return nil
}
