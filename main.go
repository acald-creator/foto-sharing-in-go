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
	ID        int    `json:"id" gorm:"primaryKey"`
	Email     User   `gorm:"embedded"`
	Caption   string `json:"caption"`
	Url       string `json:"url"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

var items = []FeedItem{
	{ID: 1, Email: User{Email: "name@example.com"}, Caption: "Caption example 1", Url: "http://example.com/image.png", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 2, Email: User{Email: "name@example.com"}, Caption: "Caption example 2", Url: "http://example.com/image.png", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

func main() {
	r := gin.Default()
	r.GET("/items", getItems)
	r.Run()
}
