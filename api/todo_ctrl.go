package api

import (
	"github.com/asanobm/go-gin-mysql-example/repo"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TodoCreate /**
//   - @summary create todo
//   - @Description: 새로운 Todo를 생성합니다.
//   - @param c *gin.Context
//   - @router /todo [post]
//     */
func TodoCreate(c *gin.Context) {
	var req CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo := repo.Todo{Title: req.Title, Description: req.Description}
	repo.DB().Create(&todo)
	c.JSON(http.StatusOK, CreateTodoResponse{ID: todo.ID.String(), Title: todo.Title, Description: todo.Description})
}
