package model

type Book struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Title       string `json:"title" db:"title"`
	ISBN        string `json:"isbn" db:"isbn"`
	Description string `json:"description" db:"description"`
}

type BookCreateInput struct {
	Name        string `json:"name" db:"name" binding:"required"`
	Title       string `json:"title" db:"title" binding:"required"`
	ISBN        string `json:"isbn" db:"isbn" binding:"required"`
	Description string `json:"description" db:"description" binding:"required"`
}

type BookUpdateInput struct {
	Name        *string `json:"name" db:"name"`
	Title       *string `json:"title" db:"title"`
	ISBN        *string `json:"isbn" db:"isbn"`
	Description *string `json:"description" db:"description"`
}
