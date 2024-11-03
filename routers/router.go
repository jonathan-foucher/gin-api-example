package routers

import (
	"github.com/gin-gonic/gin"
	"gin-api-example/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	movies := r.Group("/api/movies")

	{
		movies.GET("", api.GetMovies)
		movies.POST("", api.SaveMovie)
		movies.DELETE("/:movieId", api.DeleteMovie)
	}
	return r
}
