package controllers

import (
	"net/http"
	"time"

	"github.com/dohyeoplim/blog-server/config"
	"github.com/dohyeoplim/blog-server/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// --- DTOs ---

type CreatePostRequest struct {
	Title     string   `json:"title"`
	Slug      string   `json:"slug"`
	Excerpt   string   `json:"excerpt"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
	PostType  string   `json:"post_type"`
	Published bool     `json:"published"`
}

type PatchPostRequest struct {
	Title     *string   `json:"title"`
	Slug      *string   `json:"slug"`
	Excerpt   *string   `json:"excerpt"`
	Content   *string   `json:"content"`
	Tags      *[]string `json:"tags"`
	PostType  *string   `json:"post_type"`
	Published *bool     `json:"published"`
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

// --- Handlers ---

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

func GetAllPosts(c *gin.Context) {
	var posts []models.Post
	query := config.DB

	_, authenticated := c.Get("user_id")
	if !authenticated {
		query = query.Where("published = ?", true)
	}

	if err := query.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

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
	slug := c.Param("slug")

	var post models.Post
	query := config.DB.Where("slug = ?", slug)

	if _, authenticated := c.Get("user_id"); !authenticated {
		query = query.Where("published = ?", true)
	}

	if err := query.First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found or not accessible"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func UpdatePost(c *gin.Context) {
	slug := c.Param("slug")

	var post models.Post
	if err := config.DB.Where("slug = ?", slug).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	var req PatchPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if req.Title != nil {
		post.Title = *req.Title
	}
	if req.Slug != nil {
		post.Slug = *req.Slug
	}
	if req.Excerpt != nil {
		post.Excerpt = *req.Excerpt
	}
	if req.Content != nil {
		post.Content = *req.Content
	}
	if req.Tags != nil {
		post.Tags = *req.Tags
	}
	if req.PostType != nil {
		post.PostType = *req.PostType
	}
	if req.Published != nil {
		post.Published = *req.Published
	}

	config.DB.Save(&post)
	c.JSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	slug := c.Param("slug")

	if err := config.DB.Where("slug = ?", slug).Delete(&models.Post{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}
