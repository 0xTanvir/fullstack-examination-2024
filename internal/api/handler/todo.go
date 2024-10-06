package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zuu-development/fullstack-examination-2024/internal/api/dto"
	"github.com/zuu-development/fullstack-examination-2024/internal/domain/bo"
	"github.com/zuu-development/fullstack-examination-2024/internal/domain/definition"
)

type todoHandler struct {
	Handler
	service definition.TodoService
}

// NewTodo returns a new instance of the todo handler.
func NewTodo(s definition.TodoService) definition.TodoHandler {
	return &todoHandler{service: s}
}

// Create todo godoc
//	@Summary	Create a new todo
//	@Tags		todos
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.CreateTodoRequest			true	"Create Todo Request Body"
//	@Success	201		{object}	dto.ResponseData{data=dto.Todo}	"Success Response with Todo"
//	@Failure	400		{object}	dto.ResponseError
//	@Failure	500		{object}	dto.ResponseError
//	@Router		/todos [post]
func (t *todoHandler) Create(c echo.Context) error {
	var req dto.CreateTodoRequest
	if err := t.MustBind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ResponseError{Errors: []dto.Error{{Code: dto.CodeBadRequest, Message: err.Error()}}})
	}

	if !dto.ValidPriorities[req.Priority] {
		return c.JSON(http.StatusBadRequest,
			dto.ResponseError{Errors: []dto.Error{{Code: dto.CodeBadRequest, Message: "invalid priority"}}})
	}

	todo, err := t.service.Create(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			dto.ResponseError{Errors: []dto.Error{{Code: dto.CodeInternalServerError, Message: err.Error()}}})
	}

	return c.JSON(http.StatusCreated, dto.ResponseData{Data: todo})
}

// Update todo godoc
//	@Summary	Update a todo
//	@Tags		todos
//	@Accept		json
//	@Produce	json
//	@Param		body	body		dto.UpdateTodoRequestBody		true	"body"
//	@Param		id		path		int								true	"todo id"
//	@Success	201		{object}	dto.ResponseData{data=dto.Todo}	"Success Response with Todo"
//	@Failure	400		{object}	dto.ResponseError
//	@Failure	500		{object}	dto.ResponseError
//	@Router		/todos/{id} [put]
func (t *todoHandler) Update(c echo.Context) error {
	var req dto.UpdateTodoRequest
	if err := t.MustBind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ResponseError{Errors: []dto.Error{{Code: dto.CodeBadRequest, Message: err.Error()}}})
	}

	todo, err := t.service.Update(req)
	if err != nil {
		if err == bo.ErrNotFound {
			return c.JSON(http.StatusNotFound,
				dto.ResponseError{Errors: []dto.Error{{Code: dto.CodeNotFound, Message: "todo not found"}}})
		}
		return c.JSON(http.StatusInternalServerError,
			dto.ResponseError{Errors: []dto.Error{{Code: dto.CodeInternalServerError, Message: err.Error()}}})
	}

	return c.JSON(http.StatusOK, dto.ResponseData{Data: todo})
}

// Delete todo godoc
//	@Summary	Delete a todo
//	@Tags		todos
//	@Param		id	path	int	true	"todo id"
//	@Success	204
//	@Failure	400	{object}	dto.ResponseError
//	@Failure	404	{object}	dto.ResponseError
//	@Failure	500	{object}	dto.ResponseError
//	@Router		/todos/{id} [delete]
func (t *todoHandler) Delete(c echo.Context) error {
	var wrappedID dto.IDWrapper
	if err := t.MustBind(c, &wrappedID); err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ResponseError{Errors: []dto.Error{{Code: dto.CodeBadRequest, Message: err.Error()}}})
	}

	if err := t.service.Delete(wrappedID.ID); err != nil {
		if err == bo.ErrNotFound {
			return c.JSON(http.StatusNotFound,
				dto.ResponseError{Errors: []dto.Error{{Code: dto.CodeNotFound, Message: "todo not found"}}})
		}
		return c.JSON(http.StatusInternalServerError,
			dto.ResponseError{Errors: []dto.Error{{Code: dto.CodeInternalServerError, Message: err.Error()}}})
	}
	return c.NoContent(http.StatusNoContent)
}

// Find todo godoc
//	@Summary	Find a todo
//	@Tags		todos
//	@Produce	json
//	@Param		id	path		int								true	"todo id"
//	@Success	200	{object}	dto.ResponseData{data=dto.Todo}	"Success Response with Todo"
//	@Failure	400	{object}	dto.ResponseError
//	@Failure	404	{object}	dto.ResponseError
//	@Failure	500	{object}	dto.ResponseError
//	@Router		/todos/{id} [get]
func (t *todoHandler) Find(c echo.Context) error {
	var wrappedID dto.IDWrapper
	if err := t.MustBind(c, &wrappedID); err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ResponseError{Errors: []dto.Error{{Code: dto.CodeBadRequest, Message: err.Error()}}})
	}

	res, err := t.service.Find(wrappedID.ID)
	if err != nil {
		if err == bo.ErrNotFound {
			return c.JSON(http.StatusNotFound,
				dto.ResponseError{Errors: []dto.Error{{Code: dto.CodeNotFound, Message: "todo not found"}}})
		}
		return c.JSON(http.StatusInternalServerError,
			dto.ResponseError{Errors: []dto.Error{{Code: dto.CodeInternalServerError, Message: err.Error()}}})
	}
	return c.JSON(http.StatusOK, dto.ResponseData{Data: res})
}

// FindAll todo godoc
//	@Summary	Find all todos
//	@Tags		todos
//	@Produce	json
//	@Param		task	query		string	false	"task search by task"
//	@Param		status	query		string	false	"status search by status"
//	@Param		order	query		string	false	"order by asc or desc"
//	@Success	200		{object}	dto.ResponseData{data=[]dto.Todo}
//	@Failure	500		{object}	dto.ResponseError
//	@Router		/todos [get]
func (t *todoHandler) FindAll(c echo.Context) error {
	var tq dto.TodoQuery
	if err := t.MustBind(c, &tq); err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ResponseError{Errors: []dto.Error{{Code: dto.CodeBadRequest, Message: err.Error()}}})
	}

	res, err := t.service.FindAll(tq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			dto.ResponseError{Errors: []dto.Error{{Code: dto.CodeInternalServerError, Message: err.Error()}}})
	}
	return c.JSON(http.StatusOK, dto.ResponseData{Data: res})
}
