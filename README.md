## Sample Blog API

This is just an example on using [PostgreSQL](http://postgresql.org) + [GORM](http://jinzhu.me/gorm/) + [Gin](https://github.com/gin-gonic/gin) + [Migrate](https://github.com/mattes/migrate) to build a simple API endpoint.

### Installation

```
$ git clone git@github.com:reagent/blog.git $GOPATH/src/github.com/reagent/blog
$ cd $GOPATH/src/github.com/reagent/blog
$ go get ./vendor/github.com/mattes/migrate
$ make create-db migrate run
```

### Usage

```
$ curl -s -H "Content-Type: application/json" -d '{"title":"First Post"}' http://localhost:8080/posts | python -m json.tool
$ curl -s http://localhost:8080/posts | python -m json.tool
$ curl -s http://localhost:8080/posts/1 | python -m json.tool
$ curl -s -X DELETE http://localhost:8080/posts/1 | python -m json.tool
$ curl -s http://localhost:8080/posts/1 | python -m json.tool
```
