package middleware

import (
  "fmt"

  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect(hostname, username, dbname string) gin.HandlerFunc {
  var db  *gorm.DB
  var err error

  return func (ctx *gin.Context) {
    if db == nil {
      connection_string := fmt.Sprintf(
        "host=%s user=%s dbname=%s sslmode=disable",
        hostname, username, dbname,
      )

      db, err = gorm.Open("postgres", connection_string)
    }

    ctx.Set("db", db)

    ctx.Next()
  }
}
