package main

import (
	"github.com/asanobm/go-gin-mysql-example/api"
	"github.com/asanobm/go-gin-mysql-example/repo"
	"github.com/gin-gonic/gin"
)

func main() {

	// migrate the schema
	repo.Migration()

	r := gin.Default()
	// todo group
	todo := r.Group("/todos")
	{
		todo.POST("", api.TodoCreate)
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
