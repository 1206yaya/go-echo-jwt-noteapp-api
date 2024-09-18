package controller

import (
	"net/http"
	"strconv"

	"github.com/1206yaya/go-echo-jwt-noteapp-api/model"
	"github.com/1206yaya/go-echo-jwt-noteapp-api/usecase"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type INoteController interface {
	GetAllNotes(c echo.Context) error
	GetNoteById(c echo.Context) error
	CreateNote(c echo.Context) error
	UpdateNote(c echo.Context) error
	DeleteNote(c echo.Context) error
}

type noteController struct {
	usecase usecase.INoteUsecase
}

func NewNoteController(tu usecase.INoteUsecase) INoteController {
	return &noteController{tu}
}

func (tc *noteController) GetAllNotes(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	notesRes, err := tc.usecase.GetAllNotes(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, notesRes)
}

func (controller *noteController) GetNoteById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("noteId")
	noteId, _ := strconv.Atoi(id)
	noteRes, err := controller.usecase.GetNoteById(uint(userId.(float64)), uint(noteId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, noteRes)
}

func (controller *noteController) CreateNote(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	note := model.Note{}
	if err := c.Bind(&note); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	note.UserId = uint(userId.(float64))
	noteRes, err := controller.usecase.CreateNote(note)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, noteRes)
}

func (controller *noteController) UpdateNote(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("noteId")
	noteId, _ := strconv.Atoi(id)

	note := model.Note{}
	if err := c.Bind(&note); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	noteRes, err := controller.usecase.UpdateNote(note, uint(userId.(float64)), uint(noteId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, noteRes)
}

func (controller *noteController) DeleteNote(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("noteId")
	noteId, _ := strconv.Atoi(id)

	err := controller.usecase.DeleteNote(uint(userId.(float64)), uint(noteId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
