package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"checkmarx/api/helpers"
	"checkmarx/internal/domain/entity"
	"checkmarx/internal/domain/service"
)

type CommentHandler struct {
	Service *service.CommentService
}

func NewCommentHandler(s *service.CommentService) *CommentHandler {
	return &CommentHandler{
		Service: s,
	}
}

func (ch *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment entity.Comment

	if err := helpers.ReadJSON(w, r, &comment); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, "malformed json", nil)
		return
	}

	if comment.Content == "" {
		helpers.WriteJSON(w, http.StatusBadRequest, "content can't be empty", nil)
		return
	}

	if err := ch.Service.CreateComment(&comment); err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, "couldn't create comment", nil)
		return
	}

	helpers.WriteJSON(w, http.StatusCreated, nil, nil)
}

func (ch *CommentHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	var comment entity.Comment

	if err := helpers.ReadJSON(w, r, &comment); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, "malformed json", nil)
		return
	}

	if err := ch.Service.UpdateComment(&comment); err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, "couldn't update comment", nil)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, nil, nil)
}

func (ch *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, "invalid comment ID", nil)
		return
	}

	if err := ch.Service.DeleteComment(id); err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, "couldn't delete comment", nil)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, nil, nil)
}
