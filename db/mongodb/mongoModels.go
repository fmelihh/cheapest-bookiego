package mongodb

import (
	"cheapest-bookiego/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoDatabase = "books"
var MongoCollection = "scraped"

type MongoBookModel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Keyword  string             `bson:"keyword,omitempty"`
	BookData []models.Book      `bson:"book_data"`
}

type BookModel struct {
	C *mongo.Collection
}

func NewMongoBookModel() *BookModel {
	client, _, _, _ := connect()
	return &BookModel{C: client.Database(MongoDatabase).Collection(MongoCollection)}
}

func (c *BookModel) FindByKeyword(keyword string) (MongoBookModel, error) {
	ctx := context.TODO()
	var b MongoBookModel
	err := c.C.FindOne(ctx, bson.D{{"keyword", keyword}}).Decode(&b)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return MongoBookModel{}, nil
		}
		return MongoBookModel{}, err
	}

	return b, nil
}

func (c *BookModel) Insert(bookModel MongoBookModel) (*mongo.InsertOneResult, error) {
	ctx := context.TODO()
	return c.C.InsertOne(ctx, bookModel)
}
