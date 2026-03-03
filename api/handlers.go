package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sv-test/database"
	"sv-test/models"
)

var db = database.ConnectDB()

func GetArticles(w http.ResponseWriter, r *http.Request) {
    var articles []models.Article
    // Parse query parameters
    limitStr := r.URL.Query().Get("limit")
    offsetStr := r.URL.Query().Get("offset")

    // Default values
    limit := 10
    offset := 0

    // Convert string to int if provided
    if lim, err := strconv.Atoi(limitStr); err == nil {
        limit = lim
    }
    if off, err := strconv.Atoi(offsetStr); err == nil {
        offset = off
    }

    // Query with limit and offset
    if err := db.Limit(limit).Offset(offset).Find(&articles).Error; err != nil {
        http.Error(w, "Failed to fetch articles", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(articles)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
    var article models.Article
    // Decode JSON request body into Article struct
    if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Save to database
    if err := db.Create(&article).Error; err != nil {
        http.Error(w, "Failed to create article", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(article)
}

func GetArticleByID(w http.ResponseWriter, r *http.Request) {
    // Extract ID from query parameter (?id=1)
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    var article models.Article
    if err := db.First(&article, id).Error; err != nil {
        http.Error(w, "Article not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(article)
}

// UpdateArticle handles PUT /articles/update?id=1
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    var article models.Article
    if err := db.First(&article, id).Error; err != nil {
        http.Error(w, "Article not found", http.StatusNotFound)
        return
    }

    // Decode new values from request body
    var updated models.Article
    if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Update fields
    article.Title = updated.Title
    article.Content = updated.Content
    article.Category = updated.Category
    article.Status = updated.Status

    db.Save(&article)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(article)
}

// DeleteArticle handles DELETE /articles/delete?id=1
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    if err := db.Delete(&models.Article{}, id).Error; err != nil {
        http.Error(w, "Failed to delete article", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent) // 204 No Content
}
