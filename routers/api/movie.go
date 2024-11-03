package api

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"gin-api-example/models"
	"strconv"
)

func GetMovies(c *gin.Context) {
	fmt.Println("Get all movies")
	c.String(http.StatusOK, "Get all movies")
}

func SaveMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Post movie with id", movie.Id)
	fmt.Println("Post movie with title ", movie.Title)
	fmt.Println("Post movie with release date ", movie.ReleaseDate)
	c.JSON(http.StatusOK, movie)
}

func DeleteMovie(c *gin.Context) {
	movieId, err := strconv.Atoi(c.Params.ByName("movieId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Delete movie with id", movieId)
	c.Status(http.StatusOK)
}
