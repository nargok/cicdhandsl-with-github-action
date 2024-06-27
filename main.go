package main

import (
  "fmt"
  "het/http"
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  route.GET("/hello/:name", func(c *gin.Context) {
    name := c.Param("name")
    c.String(http.StatusOK, makeGreeting(name))
  })
}

func makeGreeting(name String) String {
  return fmt.Sprintf("Hello", %d, name)
}
