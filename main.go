package main

import (
	"net/http"
	"github.com/Henock-Yonas/project1/routes"
	"github.com/Henock-Yonas/project1/models"
	"github.com/Henock-Yonas/project1/utils"
)

func main() {
	models.Init()
	utils.LoadTemplates("templates/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
