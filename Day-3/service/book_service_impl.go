package service

import (
	"errors"
	"github.com/gtxiqbal/altera-agmc/Day-3/helper"
	"github.com/gtxiqbal/altera-agmc/Day-3/model"
	"github.com/gtxiqbal/altera-agmc/Day-3/model/dto"
	"github.com/gtxiqbal/altera-agmc/Day-3/repository"
)

type BookServiceImpl struct {
	bookRepo repository.BookRepository
}

func NewBookServiceImpl(bookRepo repository.BookRepository) *BookServiceImpl {
	return &BookServiceImpl{bookRepo: bookRepo}
}

func (bookSvc *BookServiceImpl) GetAll() dto.ResponseDTO[[]dto.BookDtoRes] {
	books := bookSvc.bookRepo.FindAll()
	var booksDtoRes []dto.BookDtoRes
	helper.PanicIfError(helper.MappingStruct(&booksDtoRes, books))
	return dto.ResponseDTO[[]dto.BookDtoRes]{
		Code:    200,
		Status:  dto.StatusSuccess,
		Message: "success get all book",
		Data:    booksDtoRes,
	}
}

func (bookSvc *BookServiceImpl) getByID(id uint) *model.Book {
	book := bookSvc.bookRepo.FindByID(id)
	if book == nil {
		helper.PanicErrorCode(404, errors.New("book not found"))
	}
	return book
}

func (bookSvc *BookServiceImpl) GetByID(id uint) dto.ResponseDTO[dto.BookDtoRes] {
	book := bookSvc.getByID(id)

	var bookDtoRes dto.BookDtoRes
	helper.PanicIfError(helper.MappingStruct(&bookDtoRes, book))

	return dto.ResponseDTO[dto.BookDtoRes]{
		Code:    200,
		Status:  dto.StatusSuccess,
		Message: "success get book",
		Data:    bookDtoRes,
	}
}

func (bookSvc *BookServiceImpl) Create(bookDtoReq dto.BookDtoReq) dto.ResponseDTO[dto.BookDtoRes] {
	var book model.Book
	helper.PanicIfError(helper.MappingStruct(&book, bookDtoReq))
	bookSvc.bookRepo.Insert(&book)

	var bookDtoRes dto.BookDtoRes
	helper.PanicIfError(helper.MappingStruct(&bookDtoRes, book))

	return dto.ResponseDTO[dto.BookDtoRes]{
		Code:    200,
		Status:  dto.StatusSuccess,
		Message: "success create book",
		Data:    bookDtoRes,
	}
}

func (bookSvc *BookServiceImpl) Update(bookDtoReq dto.BookDtoReq) dto.ResponseDTO[dto.BookDtoRes] {
	b := bookSvc.getByID(bookDtoReq.ID)
	var book model.Book
	helper.PanicIfError(helper.MappingStruct(&book, bookDtoReq))
	book.CreatedAt = b.CreatedAt
	bookSvc.bookRepo.Update(&book)

	var bookDtoRes dto.BookDtoRes
	helper.PanicIfError(helper.MappingStruct(&bookDtoRes, book))

	return dto.ResponseDTO[dto.BookDtoRes]{
		Code:    200,
		Status:  dto.StatusSuccess,
		Message: "success update book",
		Data:    bookDtoRes,
	}
}

func (bookSvc *BookServiceImpl) Delete(id uint) dto.ResponseDTO[any] {
	book := bookSvc.getByID(id)
	bookSvc.bookRepo.Delete(book)
	return dto.ResponseDTO[any]{
		Code:    200,
		Status:  dto.StatusSuccess,
		Message: "success delete book",
	}
}
