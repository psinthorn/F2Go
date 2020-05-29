package utils

import (
	"fmt"
	"net/http"
)

func ImplemenftMe() string {
	fmt.Println("Implement Me")
	c.HTML(http.StatusOK, "Implement Me")
	return
}
