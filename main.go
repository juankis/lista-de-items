package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
		title := c.PostForm("title")
		moverArchivo(c)

		fmt.Printf("title: %s;", title)
		c.JSON(200, gin.H{
			"status": "posted",
			"title":  title,
		})
	})
	router.Run(":9000")
}

func moverArchivo(c *gin.Context) string {
	path := "./pictures/"
	file, err := c.FormFile("picture")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
	}
	nameFile := path + strconv.Itoa(rand.Intn(10000)) + "_" + file.Filename
	if err := c.SaveUploadedFile(file, nameFile); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
	}
	c.String(http.StatusOK, fmt.Sprintf("El archivo %s ha sido trasladado con exito", file.Filename))
	return nameFile
}
