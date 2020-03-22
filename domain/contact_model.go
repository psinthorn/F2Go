package models

type Contact struct{
	Id 			uint64 	`json:"id"`
	Title    	string 	`json:"title"`
	SubTitle 	string 	`json:"sub_title"`
	Body     	string	`json:"body"`
	SubBody  	string	`json:"sub_body"`
	Email    	string	`json:"email"`
	LineId   	string	`json:"line_id"`
	Instagram 	string	`json:"instagram"`
	Facebook 	string	`json:"facebook"`
	Mobile   	string 	`json:"mobile"`
	Website  	string	`json:"website"`
	Status   	bool 	`json:"status"`
}

