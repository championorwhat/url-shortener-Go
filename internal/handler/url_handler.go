package handler

import (
	"encoding/json"
	"net/http"
	"url-shortener/internal/service"

	"github.com/gorilla/mux"
)

type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(service *service.URLService) *URLHandler {
	return &URLHandler{service: service}
}

type shortenRequest struct {
	URL string `json:"url"`
}

func (h *URLHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	var req shortenRequest
	json.NewDecoder(r.Body).Decode(&req)

	result, err := h.service.CreateShortURL(req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]

	result, err := h.service.GetOriginalURL(code)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, result.OriginalURL, http.StatusFound)
}
