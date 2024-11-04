package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	db "gin-api-example/database"
	"gin-api-example/models"
)

func GetMovies(c *gin.Context) {
	queries := db.New(db.GetDbConnection())
	fmt.Println("Get all movies")

	results, err := queries.GetMovies(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	movies := make([]models.Movie, len(results))
    for i, result := range results {
        movies[i] = convertDbMovieToModelsMovie(result)
    }
	c.JSON(http.StatusOK, movies)
}

func SaveMovie(c *gin.Context) {
	queries := db.New(db.GetDbConnection())
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Post movie with id", movie.Id)
	fmt.Println("Post movie with title ", movie.Title)
	fmt.Println("Post movie with release date ", movie.ReleaseDate)

	queries.SaveMovie(context.Background(), db.SaveMovieParams{
		ID: movie.Id,
		Title: movie.Title,
		ReleaseDate: movie.ReleaseDate,
	})
	c.Status(http.StatusOK)
}

func DeleteMovie(c *gin.Context) {
	queries := db.New(db.GetDbConnection())
	movieId, err := strconv.Atoi(c.Params.ByName("movieId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Delete movie with id", movieId)
	queries.DeleteMovie(context.Background(), int32(movieId))
	c.Status(http.StatusOK)
}

func convertDbMovieToModelsMovie(dbMovie db.Movie) (models.Movie) {
	movie := models.Movie{
		Id: dbMovie.ID,
		Title: dbMovie.Title,
		ReleaseDate: dbMovie.ReleaseDate,
	}
	return movie
}
