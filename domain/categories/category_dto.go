package categories

type Category struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DateCreated string `json:"date_reated"`
}
