package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

var(
	app *gin.Engine
)


func myRoute(r *gin.RouterGroup){
	r.Get("/admin", func(c *gin.Context){
		c.String(http.StatusOK,"Hello from vercek")
	})
}

func init(){
	app = gin.New()
	r := app.Group("/api")
	myRoute(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
    

    // Serve the HTTP request using the Gin router
    app.ServeHTTP(w, r)
}
