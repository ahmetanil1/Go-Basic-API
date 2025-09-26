package handlers

import (
	"library-management/internal/models"
	"library-management/internal/services"
	"net/http"

	"github.com/goccy/go-json"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookHandler struct {
	Service *services.BookService
}

func (h *BookHandler) GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := h.Service.GetBooks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	er := json.NewEncoder(w).Encode(books)
	if er != nil {
		http.Error(w, er.Error(), http.StatusInternalServerError)
	}
}

func (h *BookHandler) GetByIDBookHandler(w http.ResponseWriter, r *http.Request, id string) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	book, err := h.Service.GetBookByID(r.Context(), objectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	er := json.NewEncoder(w).Encode(book)
	if er != nil {
		http.Error(w, er.Error(), http.StatusInternalServerError)
	}
}

func (h *BookHandler) CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := h.Service.Create(r.Context(), &book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	book.ID = id
	w.Header().Set("Content-Type", "application/json")
	er := json.NewEncoder(w).Encode(book)
	if er != nil {
		http.Error(w, er.Error(), http.StatusInternalServerError)
	}
}

func (h *BookHandler) UpdateBookHandler(w http.ResponseWriter, r *http.Request, id string) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var book models.Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book.ID = objectID
	if book.ID.IsZero() {
		http.Error(w, "id is zero", http.StatusBadRequest)
	}

	if err := h.Service.UpdateBook(r.Context(), objectID, &book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	er := json.NewEncoder(w).Encode(book)
	if er != nil {
		http.Error(w, er.Error(), http.StatusInternalServerError)
	}
}

func (h *BookHandler) DeleteBookHandler(w http.ResponseWriter, r *http.Request, id string) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteBook(r.Context(), objectID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	er := json.NewEncoder(w).Encode(objectID)
	if er != nil {
		http.Error(w, er.Error(), http.StatusInternalServerError)
	}
}
