package responseobject

type Course struct {
	ID    string `json:"id"`
	Steps []Step `json:"steps"`
	Name  string `json:"name"`
}
