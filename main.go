package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	FeedItems []FeedItem
}

type FeedItem struct {
	gorm.Model
	ID        string `json:"id" gorm:"primaryKey"`
	Email     User   `gorm:"embedded"`
	Caption   string `json:"caption"`
	Url       string `json:"url"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

var items = []FeedItem{
	{ID: "1", Email: User{Email: "name@example.com"}, Caption: "Caption example 1", Url: "http://example.com/image.png", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: "2", Email: User{Email: "name@example.com"}, Caption: "Caption example 2", Url: "http://example.com/image.png", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

func getItems(c *gin.Context) {
	c.SecureJSON(http.StatusOK, items)
}

func postItems(c *gin.Context) {
	var newItem FeedItem

	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	items = append(items, newItem)
	c.SecureJSON(http.StatusCreated, newItem)
}

func getItemById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range items {
		if a.ID == id {
			c.SecureJSON(http.StatusOK, a)
			return
		}
	}
	c.SecureJSON(http.StatusNotFound, gin.H{"message": "item is not found"})
}
func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/items", getItems)
		r.GET("/items/:id", getItemById)
		r.POST("/items", postItems)
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/items", getItems)
		v2.GET("/items/:id", getItemById)
		v2.POST("/items", postItems)
	}
	r.Run(":8080")
}
