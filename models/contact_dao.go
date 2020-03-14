package models

var {
	contents := map[id]*Content{
		1: {	Id:  1,			
				Title: "Contact Us",    
				SubTitle:  	"F2 Co.,Ltd.",
				Body:      	"Contact F2 with below infomation and we will reply you within 24 hrs.",
				SubBody:   	"",
				Email:     	"psinthorn@gmail.com",	
				LineId: 	"sinthorn83",   		
				Instagram:  "sinthorn",	
				Facebook:  	"fb.me/f2coltd",				
				Mobile:    	"0989358228",		
				Website:   	"https://www.f2.co.th",				
				Status:   	true,	
			}
	}

func getContent(id intint64) *Content {
	content := 
}