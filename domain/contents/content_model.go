package contents

type ContentDefault struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	SubTitle    string `json:"sub_tilte"`
	Body        string `json:"body"`
	SubBody     string `json:"sub_body"`
	Status      bool   `json:"status"`
	CreatedDate string `json:"created_date"`
}

type Content ContentDefault {
	Category string `json"category"`
}


