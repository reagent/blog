package posts

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
)

type Post struct {
  ID    uint    `json:"id"`
  Title string  `json:"title"`
}

func Index(ctx *gin.Context) {
  var posts []Post

  db := ctx.MustGet("db").(*gorm.DB)
  db.Find(&posts)

  ctx.JSON(http.StatusOK, posts)
}

func Create(ctx *gin.Context) {
  var post Post

  db := ctx.MustGet("db").(*gorm.DB)
  ctx.Bind(&post)

  post.ID = 0

  result := db.Create(&post)

  if result.Error != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"err": result.Error.Error()})
    return
  }

  ctx.JSON(http.StatusCreated, post)
}

func Get(ctx *gin.Context) {
  var post Post
  db := ctx.MustGet("db").(*gorm.DB)

  result := db.First(&post, ctx.Param("id"))

  if result.Error != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"err": "Post not found"})
    return
  }

  ctx.JSON(http.StatusOK, post)
}

func Destroy(ctx *gin.Context) {
  var post    Post
  var post_id string
  var result  *gorm.DB

  db := ctx.MustGet("db").(*gorm.DB)

  post_id = ctx.Param("id")
  result  = db.Find(&post, post_id)

  if result.Error != nil {
    ctx.JSON(http.StatusNotFound, gin.H{"err": "Post not found"})
    return
  }

  result = db.Delete(&post)

  if result.Error != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"err": result.Error.Error()})
    return
  }

  ctx.JSON(http.StatusOK, gin.H{"id": post_id})
}
