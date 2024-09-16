package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"checkmarx/api/helpers"
	"checkmarx/internal/domain/entity"
	"checkmarx/internal/domain/service"
)

type PostHandler struct {
	Service *service.PostService
}

func NewPostHandler(s *service.PostService) *PostHandler {
	return &PostHandler{
		Service: s,
	}
}

func (ph *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post entity.Post

	if err := helpers.ReadJSON(w, r, &post); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, "malformed json", nil)
		return
	}

	if err := ph.Service.CreatePost(&post); err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, "couldn't create post", nil)
		return
	}

	helpers.WriteJSON(w, http.StatusCreated, nil, nil)
}

func (ph *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, "Invalid Post ID", nil)
		return
	}

	data, err := ph.Service.GetPost(id)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, "couldn't find the post", nil)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, data, nil)
}

func (ph *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := ph.Service.GetAll()
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, "couldn't fetch posts", nil)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, posts, nil)
}

func (ph *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	var post entity.Post

	if err := helpers.ReadJSON(w, r, &post); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, "malformed json", nil)
		return
	}

	if err := ph.Service.Update(&post); err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, "couldn't update post", nil)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, nil, nil)
}

func (ph *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, "Invalid Post ID", nil)
		return
	}

	if err := ph.Service.Delete(id); err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, "couldn't delete the post", nil)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, nil, nil)
}
