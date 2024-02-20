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

	router := gin.New()
	router.LoadHTMLGlob("templates/*.html")

	router.Static("/static", "./static")

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(middlewareFunc1, middlewareFunc2)

	router.GET("/", getHello)

	router.GET("/greet", getGreet)

	router.GET("/movie", getAllMovies)
	authRouter := router.Group("/auth", gin.BasicAuth(gin.Accounts{
		"joe":    "baseball",
		"manish": "123",
	}))

	authRouter.GET("/movie", createMovieForm)
	authRouter.POST("/movie", createMovie)

	err := router.Run("localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
}

func getHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func getGreet(c *gin.Context) {
	c.HTML(http.StatusOK, "greeting.html", nil)
}

func middlewareFunc1(c *gin.Context) {
	fmt.Println("Middleware one running")

	fmt.Println("Middleware one ended")

	c.Next()
}

func middlewareFunc2(c *gin.Context) {
	fmt.Println("Middleware two running")

	fmt.Println("Middleware two ended")

	c.Next()
}

func getAllMovies(c *gin.Context) {
	c.HTML(http.StatusOK, "all-movies.html", movies)
}

func createMovieForm(c *gin.Context) {
	c.HTML(http.StatusOK, "create-movie-form.html", nil)
}

func createMovie(c *gin.Context) {
	var newMovie Movie
	newMovie.ID = c.PostForm("id")
	newMovie.Title = c.PostForm("title")
	newMovie.Director = c.PostForm("director")
	newMovie.Price = c.PostForm("price")

	movies = append(movies, newMovie)
	c.HTML(http.StatusOK, "all-movies.html", movies)
}
