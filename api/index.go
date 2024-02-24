package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
//   fmt.Fprintf(w, "<h1>Hello from Go!</h1>")

server := New()

	server.GET("/", func(context *Context) {
		context.JSON(200, H{
			"message": "hello go from vercel !!!!",
		})
	})

	server.Handle(w, r)
}