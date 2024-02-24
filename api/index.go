package handler

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    // Create a new Gin router
    router := gin.Default()

    // Define a GET endpoint at the root path "/"
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "hello go from vercel !!!!",
        })
    })

    // Serve the HTTP request using the Gin router
    router.ServeHTTP(w, r)
}
