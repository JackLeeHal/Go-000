package data

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Hello interface {
	ReSendMessage(context.Context, string) string
}

type MongoRepo struct {
	collection *mongo.Collection
}

func NewMongoRepo(collection *mongo.Collection) Hello {
	return &MongoRepo{collection: collection}
}

func (m *MongoRepo) ReSendMessage(ctx context.Context, req string) string {
	return "hi," + req
}
