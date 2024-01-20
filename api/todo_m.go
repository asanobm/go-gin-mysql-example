package api

// CreateTodoRequest /**
//   - @title TodoCreate Todo Model
//   - @description create todo model
//   - @param title formData string true "title"
//   - @param description formData string true "description"
//     */
type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type CreateTodoResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
