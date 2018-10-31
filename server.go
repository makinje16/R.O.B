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
	Code         string
	Message      string
	URL          string
	Articles     []headlines
}

func GetSourceHeadlines(source string, newsAPIKey string) NewsResponse {
	var newsResponse NewsResponse
	url := "https://newsapi.org/v2/top-headlines?sources=" + source + "&apiKey=" + newsAPIKey
	resp, err := http.Get(url)
	println(url)
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
		panic(err)
	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		_ = json.Unmarshal(bodyBytes, &newsResponse)
		newsResponse.URL = url
		return newsResponse
	}
}

func main() {
	if len(os.Args) != 2 {
		print("Usage: ./server <news-api-key>\n")
		os.Exit(99)
	}
	r := gin.Default()
	r.GET("/headlines/ign", func(c *gin.Context) {
		c.JSON(http.StatusOK, GetSourceHeadlines("ign", os.Args[1]))
	})

	r.GET("/headlines/polygon", func(c *gin.Context) {
		c.JSON(http.StatusOK, GetSourceHeadlines("polygon", os.Args[1]))
	})

	r.GET("/headlines/techcrunch", func(c *gin.Context) {
		c.JSON(http.StatusOK, GetSourceHeadlines("techcrunch", os.Args[1]))
	})

	r.GET("/headlines/hacker-news", func(c *gin.Context) {
		c.JSON(http.StatusOK, GetSourceHeadlines("hacker-news", os.Args[1]))
	})
	r.Run(":80")
}
