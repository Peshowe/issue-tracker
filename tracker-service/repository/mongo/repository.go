package mongo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/Peshowe/issue-tracker/tracker-service/tracker"
)

//mongoRepository is the struct for interfacing with the database (should implement TrackerRepository)
type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

//newMongoClient creates a MongoDB client to a db hosted at the given mongoURL
func newMongoClient(mongoURL, mongoUser, mongoPass string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()

	//define the options
	opts := options.Client().ApplyURI(mongoURL)

	if mongoUser != "" && mongoPass != "" {
		credentials := options.Credential{
			Username: mongoUser,
			Password: mongoPass,
		}
		opts = opts.SetAuth(credentials)
	}

	//create the connection
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	//make sure we can read from our db
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

//NewMongoRepository creates a mongoRepository instance (that implements the TrackerRepository interface)
func NewMongoRepository(mongoURL, mongoDB, mongoUser, mongoPass string, mongoTimeout int) (tracker.TrackerRepository, error) {
	repo := &mongoRepository{
		timeout:  time.Duration(mongoTimeout) * time.Second,
		database: mongoDB,
	}
	client, err := newMongoClient(mongoURL, mongoUser, mongoPass, mongoTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewMongoRepo")
	}
	repo.client = client
	return repo, nil
}
