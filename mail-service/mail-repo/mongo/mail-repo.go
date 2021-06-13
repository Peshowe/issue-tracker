package mongo

import (
	"context"
	"time"

	"github.com/Peshowe/issue-tracker/mail-service/mailer"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//mongoRepository is the struct for interfacing with the database (should implement MailRepository)
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

//NewMongoRepository creates a mongoRepository instance (that implements the MailRepository interface)
func NewMongoRepository(mongoURL, mongoDB, mongoUser, mongoPass string, mongoTimeout int) (mailer.MailRepo, error) {
	repo := &mongoRepository{
		timeout:  time.Duration(mongoTimeout) * time.Second,
		database: mongoDB,
	}
	client, err := newMongoClient(mongoURL, mongoUser, mongoPass, mongoTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewMongoRepo")
	}

	repo.client = client

	//set the index to be the user field
	collection := repo.getPreferencesCollection()
	mod := mongo.IndexModel{
		Keys: bson.M{
			"user": 1, // index in ascending order
		}, Options: options.Index().SetUnique(true),
	}
	_, err = collection.Indexes().CreateOne(context.TODO(), mod)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewMongoRepo")
	}

	return repo, nil
}

//getPreferencesCollection gets a reference to the user preferences collection in the db
func (r *mongoRepository) getPreferencesCollection() *mongo.Collection {
	return r.client.Database(r.database).Collection("preferences")
}

func (r *mongoRepository) SetUserPreference(ctx context.Context, userPreference *mailer.UserPreference) error {
	//get the collection
	collection := r.getPreferencesCollection()

	//contruct the filter
	filter := bson.M{"user": userPreference.User}

	//insert into the db
	_, err := collection.ReplaceOne(ctx, filter, userPreference, options.Replace().SetUpsert(true))
	if err != nil {
		return errors.Wrap(err, "mailRepo.SetUserPreference")
	}

	return nil

}

func (r *mongoRepository) GetUserPreference(ctx context.Context, user string) (*mailer.UserPreference, error) {
	//get the collection
	collection := r.getPreferencesCollection()

	//contruct the filter
	filter := bson.M{"user": user}

	//query the db
	userPreference := &mailer.UserPreference{}
	err := collection.FindOne(ctx, filter).Decode(&userPreference)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			//it's ok if there are no entries
			return userPreference, nil
		}
		return nil, errors.Wrap(err, "mailRepo.GetUserPreference")
	}
	return userPreference, nil
}
