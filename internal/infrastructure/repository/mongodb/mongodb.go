package mongodb

import (
	"akim/internal/config"
	"akim/internal/domain/model"
	"akim/utility/tools/mongoFilter"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoDBRepository struct {
	Client     *mongo.Client
	Db         *mongo.Database
	Collection *mongo.Collection
}

func NewMongoDBRepository(config *config.Config) (*MongoDBRepository, error) {
	clientOptions := options.Client().ApplyURI("mongodb://" + config.MongoDB.Host + ":" + config.MongoDB.Port)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("error creating MongoDB client:", err)
		return nil, errors.New(fmt.Sprintf("trouble with connection mongo: %v", err))
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("error creating MongoDB client:", err)
		return nil, errors.New(fmt.Sprintf("trouble with ping mongo: %v", err))
	}

	db := client.Database(config.MongoDB.DataBase)
	collection := db.Collection(config.Collection)
	return &MongoDBRepository{
		Client:     client,
		Db:         db,
		Collection: collection,
	}, nil
}

func (m *MongoDBRepository) FindByInterval(ctx context.Context, artifact *model.FuzzyArchitecturalArtifact) ([]model.FuzzyArchitecturalArtifact, error) {
	filter := mongoFilter.SearchFiler(artifact.IntervalStart, artifact.IntervalEnd)

	cur, err := m.Collection.Find(ctx, filter)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error with finding in mongo: %v", err))
	}
	defer cur.Close(ctx)
	var buildings []model.FuzzyArchitecturalArtifact
	if err := cur.All(ctx, &buildings); err != nil {
		return nil, errors.New(fmt.Sprintf("error with marshalling to struct: %v", err))
	}
	if len(buildings) == 0 {
		return nil, model.ErrNoResults
	}
	return buildings, nil
}
