package api

import (
	"simpleOpenapi/internal/http/gen"
	"simpleOpenapi/internal/http/usecase"

	"github.com/labstack/echo/v4"
)

type Api struct {
	pet  *usecase.Pet
	book *usecase.Book
}

func NewApi() *Api {
	return &Api{pet: usecase.NewPet(), book: usecase.NewBook()}
}

var _ gen.ServerInterface = (*Api)(nil)

func (p *Api) FindBooks(ctx echo.Context, params gen.FindBooksParams) error {
	return p.book.FindBooks(ctx, params)
}
func (p *Api) FindPets(ctx echo.Context, params gen.FindPetsParams) error {
	return p.pet.FindPets(ctx, params)
}
func (p *Api) AddBook(ctx echo.Context) error {
	return p.book.AddBook(ctx)
}
func (p *Api) AddPet(ctx echo.Context) error {
	return p.pet.AddPet(ctx)
}
func (p *Api) FindBookById(ctx echo.Context, id int64) error {
	return p.book.FindBookById(ctx, id)
}

func (p *Api) FindPetById(ctx echo.Context, id int64) error {
	return p.pet.FindPetById(ctx, id)
}
func (p *Api) DeleteBook(ctx echo.Context, id int64) error {
	return p.book.DeleteBook(ctx, id)
}
func (p *Api) DeletePet(ctx echo.Context, id int64) error {
	return p.pet.DeletePet(ctx, id)
}
