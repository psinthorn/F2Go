package main

import (
	"net/http"
)

func PublicStatic() {
	fs := http.FileServer(http.Dir("assets/"))

}
