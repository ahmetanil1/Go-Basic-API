package services

import (
	"context"
	"errors"
	"library-management/internal/models"
	"library-management/internal/repos"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookService struct {
	repo *repos.BookRepository
}

func NewBookService(repo *repos.BookRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (s *BookService) Create(ctx context.Context, book *models.Book) (primitive.ObjectID, error) {
	if book.Title == "" {
		return primitive.NilObjectID, errors.New("title is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return s.repo.CreateBook(ctx, book)
}

func (s *BookService) GetBooks(ctx context.Context) ([]models.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return s.repo.FindAll(ctx)
}

func (s *BookService) GetBookByID(ctx context.Context, id primitive.ObjectID) (*models.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return s.repo.FindByID(ctx, id)
}

func (s *BookService) UpdateBook(ctx context.Context, id primitive.ObjectID, book *models.Book) error {
	if book.ID.IsZero() {
		return errors.New("id is zero")
	}
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return s.repo.Update(ctx, book)
}

func (s *BookService) DeleteBook(ctx context.Context, id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return s.repo.Delete(ctx, id)
}
