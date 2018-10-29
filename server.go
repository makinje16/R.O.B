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

func GetIgnHeadlines() NewsResponse {
	newsAPIKey := os.Getenv("NEWS_API_KEY")
	var ignResponse NewsResponse
	resp, err := http.Get("https://newsapi.org/v2/top-headlines?sources=ign&apiKey=" + newsAPIKey)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &ignResponse)
		if err == nil {
			return ignResponse
		}
	}
	return ignResponse
}

func GetPolygonHeadlines() NewsResponse {
	var polygonResponse NewsResponse
	resp, err := http.Get("https://newsapi.org/v2/top-headlines?sources=polygon&apiKey=" + newsAPIKey)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &polygonResponse)
		if err == nil {
			return polygonResponse
		}
	}
	return polygonResponse
}

func main() {
	r := gin.Default()
	r.GET("/ign", func(c *gin.Context) {
		c.JSON(http.StatusOK, GetIgnHeadlines())
	})

	r.GET("/polygon", func(c *gin.Context) {
		c.JSON(http.StatusOK, GetPolygonHeadlines())
	})
	r.Run()
}
