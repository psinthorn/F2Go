package users

import "encoding/json"

type PublicUser struct {
	Id          int64  `json:"id"`
	Status      string `json:"status"`
	DateCreated string `json:"date_created`
}

type PrivateUser struct {
	Id         int64  `json:"id"`
	First_Name string `json:"first_name`
	Last_Name  string `json:"last_name"`
	//Email       string `json:"email"`
	Status      string `json:"status"`
	DateCreated string `json:"date_created`
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}

func (user *User) Marshall(isPubllic bool) interface{} {
	if isPubllic {
		return PublicUser{
			Id:          user.Id,
			Status:      user.Status,
			DateCreated: user.DateCreated,
		}
	}
	//Return private
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}
