package course

import (
	"database/sql"
)

func CreateCourse(request *Course, db *sql.DB) error {

	createQuery := `INSERT INTO COURSES (INSTRUCTOR,DESCRIPTION,DIFFICULTY,FEE,CERTPATH,TITLE,SUBTITLE,ENLISTED,LIKES) VALUES (?,?,?,?,?,?,?,?,?);`
	_, err := db.Exec(createQuery, request.Instructor, request.Description, request.Difficulty, request.Fee, request.Certpath, request.Title, request.Subtitle, request.Enlisted, request.Likes)
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
