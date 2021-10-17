package course

type CreateRequest struct {
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
