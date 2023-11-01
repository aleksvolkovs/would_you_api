package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
)

func main() {
	r := gin.Default()

	//Example
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World",
		})
	})

	r.GET("/question/:lang/:id", func(c *gin.Context) {
		lang := c.Param("lang")
		idStr := c.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}

		filename := lang + ".json"
		filePath := filepath.Join(".", filename)

		fileContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}

		var data map[string][]map[string]interface{}
		err = json.Unmarshal(fileContent, &data)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error parsing JSON"})
			return
		}

		var foundQuestion string
		for _, q := range data["questions"] {
			if int(q["id"].(float64)) == id {
				foundQuestion = q["question"].(string)
				break
			}
		}

		if foundQuestion == "" {
			c.JSON(404, gin.H{"error": "Question not found"})
		} else {
			c.JSON(200, gin.H{"question": foundQuestion})
		}
	})

	r.Run(":8080")
}
