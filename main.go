package main

import (
	"log"
	"net/http"
	"sv-test/api"
	"sv-test/database"
)

func main() {
    db := database.ConnectDB()
    database.MigrateDB(db)

    http.HandleFunc("/articles", api.GetArticles)
    http.HandleFunc("/articles/get", api.GetArticleByID)
    http.HandleFunc("/articles/create", api.CreateArticle)
    http.HandleFunc("/articles/update", api.UpdateArticle)
    http.HandleFunc("/articles/delete", api.DeleteArticle)

    log.Println("Article Service running on http://localhost:5000")
    http.ListenAndServe(":5000", nil)
}