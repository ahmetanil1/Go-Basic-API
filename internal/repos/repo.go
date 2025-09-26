package repos

import (
	"context"
	"library-management/internal/db"
	"library-management/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepository struct {
	collection *mongo.Collection
}

func NewBookRepository() *BookRepository {
	return &BookRepository{
		collection: db.Database.Collection("books"),
	}
}

func (r *BookRepository) CreateBook(ctx context.Context, book *models.Book) (primitive.ObjectID, error) {
	book.ID = primitive.NewObjectID()
	book.PublishedAt = time.Now()

	// add single object to collection (InsertOne)
	res, err := r.collection.InsertOne(ctx, book)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func (r *BookRepository) FindAll(ctx context.Context) ([]models.Book, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var books []models.Book
	if err := cursor.All(ctx, &books); err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Book, error) {
	var book models.Book

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) Update(ctx context.Context, book *models.Book) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": book.ID},
		bson.M{"$set": bson.M{
			"title":       book.Title,
			"author":      book.Author,
			"quantity":    book.Quantity,
			"publishedAt": time.Now(),
		}},
	)
	return err
}

func (r *BookRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
