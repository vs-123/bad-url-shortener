package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var urlCodes = make(map[string]string, 0)

func updateUrlCodeData() error {
	dataFile, err := os.ReadFile("data.txt")
	if err != nil {
		return err
	}

	data := strings.ReplaceAll(string(dataFile), "\r", "")
	urlCodes = make(map[string]string, 0)

	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}
		splitLine := strings.Fields(line)

		code := splitLine[0]
		url := splitLine[1]
		urlCodes[code] = url
	}

	return nil
}

func appendToDataFile(newLine string) error {
	dataFile, err := os.ReadFile("data.txt")
	if err != nil {
		return nil
	}

	os.WriteFile("data.txt", []byte(fmt.Sprintf("%s\n%s", string(dataFile), newLine)), os.ModeAppend)
	return nil
}

func main() {
	r := gin.Default()
	updateUrlCodeData()

	// Initital page, no use
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusFound, "Hello world!")
	})

	// Go to a shortened url
	r.GET("/:code", func(c *gin.Context) {
		if url, ok := urlCodes[c.Param("code")]; ok {
			c.Redirect(http.StatusFound, url)
		} else {
			c.String(http.StatusNotFound, "404 Not Found")
		}
	})

	// Add new url
	r.GET("/addnew/https/:newCode/:newUrl", func(c *gin.Context) {
		if _, ok := urlCodes[c.Param("newCode")]; ok {
			c.String(http.StatusConflict, "Name taken.")
		} else {
			newCode := c.Param("newCode")
			newUrl := c.Param("newUrl")
			newLine := fmt.Sprintf("%s https://%s", newCode, newUrl)

			time.Sleep(2 * time.Second)
			appendToDataFile(newLine)

			updateUrlCodeData()
			c.String(http.StatusFound, "Success!")
		}
	})

	r.Run()
}
