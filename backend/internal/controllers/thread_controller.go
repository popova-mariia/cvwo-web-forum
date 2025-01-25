package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/popova-mariia/cvwo-web-forum/internal/db"
	"github.com/popova-mariia/cvwo-web-forum/internal/models"
)

// CreateThread ...
func CreateThread(c *gin.Context) {
	var thread models.Thread
	if err := c.ShouldBindJSON(&thread); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO threads (title, topic, creator_id, theme_preference)
              VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.DB.QueryRow(query, thread.Title, thread.Topic, thread.CreatorID, thread.ThemePreference).Scan(&thread.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, thread)
}

// ListThreads ...
func ListThreads(c *gin.Context) {
	rows, err := db.DB.Query(`SELECT id, title, topic, creator_id, theme_preference FROM threads`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var threads []models.Thread
	for rows.Next() {
		var t models.Thread
		err := rows.Scan(&t.ID, &t.Title, &t.Topic, &t.CreatorID, &t.ThemePreference)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		threads = append(threads, t)
	}

	c.JSON(http.StatusOK, threads)
}

// CreatePost ...
func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO posts (content, thread_id, author_id) VALUES ($1, $2, $3) RETURNING id, created_at`
	err := db.DB.QueryRow(query, post.Content, post.ThreadID, post.AuthorID).Scan(&post.ID, &post.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

// ListPosts ...
func ListPosts(c *gin.Context) {
	threadIDParam := c.Param("thread_id")
	threadID, err := strconv.Atoi(threadIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid thread ID"})
		return
	}

	query := `SELECT id, content, thread_id, author_id, created_at FROM posts WHERE thread_id=$1 ORDER BY created_at DESC`
	rows, err := db.DB.Query(query, threadID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var p models.Post
		err := rows.Scan(&p.ID, &p.Content, &p.ThreadID, &p.AuthorID, &p.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		posts = append(posts, p)
	}
	c.JSON(http.StatusOK, posts)
}

// Placeholder - Summaries (AI, polls, notifications, etc.)
func GetThreadSummary(c *gin.Context) {
	// In real use-case, call your AI service or ChatGPT API, etc.
	// This is just a placeholder returning some static text
	c.JSON(http.StatusOK, gin.H{"summary": "This is a placeholder for the AI-based thread summary."})
}
