package models

type Contact struct{
	Id 			uint64 	`json: "id"`
	Title    	string 	`json: "title"`
	SubTitle 	string 	`json: "subtitle"`
	Body     	string	`json: "body"`
	SubBody  	string	`json: subbody`
	Email    	string	`json: "email"`
	LineId   	string	`json: "linid"`
	Instagram 	string	`json: "instagram"`
	Facebook 	string	`json: "facebook"`
	Mobile   	string 	`json: "mobile"`
	Website  	string	`json: "website"`
	Status   	bool 	`json: "status"`
}

