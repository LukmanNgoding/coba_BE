package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID       uint
	Content  string
	Owner    string
	Comment  string
	Photo    string
	Username string
}

type Repository interface {
	Insert(newPost Core) (Core, error)    //Add post
	Delete(ID uint) error                 //Delete post
	Update(updateData Core) (Core, error) //Update post
	GetAllPost() ([]Core, error)          //shows all
}

type Service interface {
	AddPost(newPost Core) (Core, error)
	Delete(id uint) error
	UpdatePost(updateData Core) (Core, error)
	ShowAllPost() ([]Core, error)
}

type Handler interface {
	AddPost() echo.HandlerFunc
	DeleteByID() echo.HandlerFunc
	UpdatePost() echo.HandlerFunc
	ShowAllPost() echo.HandlerFunc
}
