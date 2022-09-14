package service

import (
	"github.com/gtxiqbal/altera-agmc/Day-2/model/dto"
)

type BookService interface {
	GetAll() dto.ResponseDTO[[]dto.BookDtoRes]
	GetByID(id uint) dto.ResponseDTO[dto.BookDtoRes]
	Create(bookDtoReq dto.BookDtoReq) dto.ResponseDTO[dto.BookDtoRes]
	Update(bookDtoReq dto.BookDtoReq) dto.ResponseDTO[dto.BookDtoRes]
	Delete(id uint) dto.ResponseDTO[any]
}
