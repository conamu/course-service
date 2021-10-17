package course

type Course struct {
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Description string `json:"description"`
	Instructor  string `json:"instructor"`
	Difficulty  int    `json:"difficulty"`
	Fee         string `json:"fee"`
	Certpath    string `json:"certpath"`
	Enlisted    string `json:"enlisted"`
	Likes       int    `json:"likes"`
}

type CourseMin struct {
	Name       string `json:"name"`
	Difficulty int    `json:"difficulty"`
	Fee        string `json:"fee"`
	Likes      int    `json:"likes"`
}

type AllCoursesResponse struct {
	Courses map[string]CourseMin `json:"courses"`
}

type ValidationResponse struct {
	Role string `json:"role"`
}
