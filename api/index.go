package handler

import (
	"fmt"
	"net/http"

	gee "github.com/tbxark/g4vercel"
)

var server = gee.Default()

func init() {
	server.GET("/", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"message": "hello go from vercel!",
		})
	})

	server.GET("/hello", func(c *gee.Context) {
		name := c.Query("name")
		if name == "" {
			c.JSON(400, gee.H{
				"message": "name not found",
			})
		} else {
			c.JSON(200, gee.H{
				"data": fmt.Sprintf("Hello %s!", name),
			})
		}
	})
	server.GET("/user/:id", func(c *gee.Context) {
		c.JSON(400, gee.H{
			"data": gee.H{
				"id": c.Param("id"),
			},
		})
	})
	server.GET("/long/long/long/path/*test", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"data": gee.H{
				"url": c.Path,
			},
		})
	})

}

// Handler è compatibile sia con Vercel che con localhost
func Handler(w http.ResponseWriter, r *http.Request) {
	// Usa il metodo Handler() invece di ServeHTTP
	server.Handle(w, r)
}
