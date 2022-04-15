package Courses

type WebApp struct {
	CoursePic   []byte `json:"course_pic"`
	Thumbnail   []byte `json:"thumbnail"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Link        string `json:"link"`
}
