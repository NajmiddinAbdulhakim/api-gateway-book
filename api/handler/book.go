package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	pb "github.com/NajmiddinAbdulhakim/ude/api-gateway/genproto"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/NajmiddinAbdulhakim/ude/api-gateway/api/model"
	"github.com/gin-gonic/gin"
)

// CreateBook creates book
// @Summary Create book summary
// @Description This api is using create new book
// @Tags Book
// @Accept json
// @Produce json
// @Param category query string false "Category" Enums(Roman, Qissa, Hikoyalar, Ilmiy-ommabop, Diniy)
// @Success 201 {string} Success
// @Param book body model.CreateBook true "book body"
// @Router /book/create [post]
func (h *Handler) CreateBook(c *gin.Context) {

	var (
		body        model.Book
		jspbMarshel protojson.MarshalOptions
	)
	jspbMarshel.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			`error`: err.Error(),
		})
		log.Println("failed to bind json for book in craeate")
		return
	}
	category := c.Query("category")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()

	req := pb.CreateBookReq{
		Title:        body.Title,
		AuthorName:   body.AuthorName,
		CategoryName: category,
	}

	response, err := h.serviceManager.BookService().Create(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			`error`: err.Error(),
		})
		log.Println("failed to create book: ", err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetByIdBook gets book by id
// @Summary Create book summary
// @Description This api is using getting book by id
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "book id"
// @Success 200 {string} model.Book
// @Router /book/{id} [get]
func (h *Handler) GetBookById(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()

	res, err := h.serviceManager.BookService().GetById(
		ctx, &pb.BookByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			`error`: err.Error(),
		})
		log.Println("failed to get book by id: ", err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateBook updates book
// @Summary update book summary
// @Description This api is using update new book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "book id"
// @Param category query string false "Category" Enums(Roman, Qissa, Hikoyalar, Ilmiy-ommabop, Diniy)
// @Param book body model.UpdateBook true "book body"
// @Success 200 {string}
// @Router /book/{id} [put]
func (h *Handler) UpdateBook(c *gin.Context) {
	var (
		body        model.Book
		jspbMarshel protojson.MarshalOptions
	)
	jspbMarshel.UseProtoNames = true
	
	guid := c.Param("id")
	
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			`error`: err.Error(),
		})
		log.Println("failed to bind json for book in update")
		return
	}
	category := c.Query("category")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()

	req := pb.UpdateBookReq{
		Id:           guid,
		Title:        body.Title,
		AuthorName:   body.AuthorName,
		CategoryName: category,
	}

	res, err := h.serviceManager.BookService().Update(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			`error`: err.Error(),
		})
		log.Println("failed to update book: ", err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteById deletes book by id
// @Summary Create book summary
// @Description This api is using deleteting book by id
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "book id"
// @Success 200 {string}
// @Router /book/{id} [delete]
func (h *Handler) DeleteBookById(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()

	res, err := h.serviceManager.BookService().Delete(
		ctx, &pb.BookByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			`error`: err.Error(),
		})
		log.Println("failed to delete book by id: ", err)
		return
	}

	c.JSON(http.StatusOK, res)
}
