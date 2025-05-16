package controllers

import (
	"net/http"
	"time"

	"github.com/dohyeoplim/blog-server/config"
	"github.com/dohyeoplim/blog-server/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreatePostRequest struct {
	Title     string   `json:"title"`
	Slug      string   `json:"slug"`
	Excerpt   string   `json:"excerpt"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
	PostType  string   `json:"post_type"`
	Published bool     `json:"published"`
}

func CreatePost(c *gin.Context) {
	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	post := models.Post{
		ID:        uuid.New(),
		Title:     req.Title,
		Slug:      req.Slug,
		Content:   req.Content,
		Excerpt:   req.Excerpt,
		Tags:      req.Tags,
		PostType:  req.PostType,
		Published: req.Published,
	}

	config.DB.Create(&post)
	c.JSON(http.StatusCreated, post)
}

type PostSummary struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	Excerpt   string    `json:"excerpt"`
	Tags      []string  `json:"tags"`
	PostType  string    `json:"post_type"`
	Published bool      `json:"published"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetAllPosts(c *gin.Context) {
	var posts []models.Post
	config.DB.Find(&posts)

	var summaries []PostSummary
	for _, post := range posts {
		summaries = append(summaries, PostSummary{
			ID:        post.ID,
			Title:     post.Title,
			Slug:      post.Slug,
			Excerpt:   post.Excerpt,
			Tags:      post.Tags,
			PostType:  post.PostType,
			Published: post.Published,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, summaries)
}

func GetPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	result := config.DB.First(&post, "id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	if err := config.DB.First(&post, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	post.Title = req.Title
	post.Slug = req.Slug
	post.Content = req.Content
	post.Excerpt = req.Excerpt
	post.Tags = req.Tags
	post.PostType = req.PostType
	post.Published = req.Published
	config.DB.Save(&post)

	c.JSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Post{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}
