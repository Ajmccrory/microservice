// main.go

package main

import "github.com/gin-gonic/gin"

var router *gin.Engine

func main() {
	// Set the router as the default
	router := gin.Default()

	//process the templates at the start, makes serving pages faster
	router.LoadHTMLGlob("templates/*")

	// Define the route for the index page and display the idnex.html
	router.GET("/", func(c *gin.Context) {

		// Call the HTMl method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200
			http.StatusOK,
			//use template
			"index.html",
			// pass the data that page uses
			gin.H{
				"title": "Home Page",
			},
		)
	})

	//serve the app
	router.Run()
}