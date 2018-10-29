package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type headlines struct {
	Author      string
	Title       string
	Description string
	Url         string
	UrlToImage  string
	PublishedAt string
	Content     string
}

type NewsResponse struct {
	Status       string
	TotalResults int
	Articles     []headlines
}

func GetSourceHeadlines(source string) NewsResponse {
	newsAPIKey := os.Getenv("NEWS_API_KEY")
	var newsResponse NewsResponse
	resp, err := http.Get("https://newsapi.org/v2/top-headlines?sources=" + source + "&apiKey=" + newsAPIKey)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &newsResponse)
		if err == nil {
			return newsResponse
		}
	}
	return newsResponse
}

func main() {
	r := gin.Default()
	r.GET("/headlines/ign", func(c *gin.Context) {
		c.JSON(http.StatusOK, GetSourceHeadlines("ign"))
	})

	r.GET("/headlines/polygon", func(c *gin.Context) {
		c.JSON(http.StatusOK, GetSourceHeadlines("polygon"))
	})

	r.GET("/headlines/techcrunch", func(c *gin.Context) {
		c.JSON(http.StatusOK, GetSourceHeadlines("techcrunch"))
	})

	r.GET("/headlines/hacker-news", func(c *gin.Context) {
		c.JSON(http.StatusOK, GetSourceHeadlines("hacker-news"))
	})
	r.Run()
}
