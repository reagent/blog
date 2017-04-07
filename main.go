package main

import (
  "os"

  "github.com/gin-gonic/gin"

  "github.com/reagent/blog/middleware"
  "github.com/reagent/blog/posts"
)

func main() {
  connection_middleware := middleware.Connect(
    os.Getenv("DB_HOST"),
    os.Getenv("DB_USER"),
    os.Getenv("DB_NAME"),
  )

  router := gin.Default()
  router.Use(connection_middleware)

  router.GET("/posts", posts.Index)
  router.POST("/posts", posts.Create)
  router.GET("/posts/:id", posts.Get)
  router.DELETE("/posts/:id", posts.Destroy)

  router.Run()
}
