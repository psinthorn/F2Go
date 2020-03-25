package domain

type About struct {
	Id       uint64 `json:"id"`
	Title    string `json:"title"`
	SubTitle string `json:"sub_title"`
	Body     string `json:"body"`
	SubBody  string `json:"sub_body"`
	Status   bool   `json:"status"`
}
