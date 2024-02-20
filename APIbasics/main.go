package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Price    string `json:"price"`
}

var movies = []Movie{
	{ID: "1", Title: "The Dark Knight", Director: "Christopher Nolan", Price: "5.99"},
	{ID: "2", Title: "Tommy Boy", Director: "Peter Segal", Price: "2.99"},
	{ID: "3", Title: "The Shawshank Redemption", Director: "Frank Darabont", Price: "7.99"},
}

func main() {
	router := gin.Default()

	router.GET("/movie", getMovies)

	router.GET("/movie/:id", getMovieByID)

	router.POST("/movie", createMovie)

	router.PATCH("/movie/:id", updateMoviePrice)

	router.DELETE("/movie/:id", deleteMovie)

	err := router.Run("localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
}

func getMovies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, movies)
}

func getMovieByID(c *gin.Context) {
	id := c.Param("id")
	var foundMovie Movie

	for _, v := range movies {
		if id == v.ID {
			foundMovie = v
			break
		}
	}

	if foundMovie.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, foundMovie)
}

func createMovie(c *gin.Context) {
	var newMovie Movie

	err := c.BindJSON(&newMovie)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	movies = append(movies, newMovie)
	c.IndentedJSON(http.StatusCreated, movies)
}

func updateMoviePrice(c *gin.Context) {
	var index int
	id := c.Param("id")
	for k, v := range movies {
		if id == v.ID {
			index = k
		}
	}
	movies[index].Price = "9.99"
	c.IndentedJSON(http.StatusOK, movies[index])
}

func deleteMovie(c *gin.Context) {
	var index int
	id := c.Param("id")
	for k, v := range movies {
		if id == v.ID {
			index = k
		}
	}
	movies = append(movies[:index], movies[index+1:]...)
	c.IndentedJSON(http.StatusOK, movies)
}
