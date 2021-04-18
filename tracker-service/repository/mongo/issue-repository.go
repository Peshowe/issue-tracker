package mongo

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Peshowe/issue-tracker/tracker-service/tracker/issue"
)

//getIssuesCollection gets a reference to the issues collection in the db
func (r *mongoRepository) getIssuesCollection() *mongo.Collection {
	return r.client.Database(r.database).Collection("issues")
}

//GetIssueById returns a single issue that matches the given id
func (r *mongoRepository) GetIssueById(id string) (*issue.Issue, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	//get the collection
	collection := r.getIssuesCollection()

	//contruct the filter
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Issue.GetIssueById")
	}
	filter := bson.M{"_id": idPrimitive}

	//query the db
	issueObj := &issue.Issue{}
	err = collection.FindOne(ctx, filter).Decode(&issueObj)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(issue.ErrIssueNotFound, "repository.Issue.GetIssueById")
		}
		return nil, errors.Wrap(err, "repository.Issue.GetIssueById")
	}
	return issueObj, nil
}

// //GetIssuesById returns a slice of issues that matche the given ids
// func (r *mongoRepository) GetIssuesById(ids []string) ([]*issue.Issue, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
// 	defer cancel()

// 	//get the collection
// 	collection := r.getIssuesCollection()

// 	//contruct the filter
// 	idPrimitive, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "repository.Issue.GetIssueById")
// 	}
// 	filter := bson.M{"_id": idPrimitive}

// 	//query the db
// 	issueObj := &issue.Issue{}
// 	err = collection.FindOne(ctx, filter).Decode(&issueObj)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil, errors.Wrap(issue.ErrIssueNotFound, "repository.Issue.GetIssueById")
// 		}
// 		return nil, errors.Wrap(err, "repository.Issue.GetIssueById")
// 	}
// 	return issueObj, nil
// }

//GetIssuesByProject returns all issues in from the given project
func (r *mongoRepository) GetIssuesByProject(projectId string) ([]*issue.Issue, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	//get the collection
	collection := r.getIssuesCollection()

	//construct the filter
	// idPrimitive, err := primitive.ObjectIDFromHex(projectId)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "repository.Issue.GetIssuesByProject")
	// }
	filter := bson.M{"project": projectId}

	//query the db
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(issue.ErrIssueNotFound, "repository.Issue.GetIssuesByProject")
		}
		return nil, errors.Wrap(err, "repository.Issue.GetIssuesByProject")
	}
	defer cur.Close(ctx)
	var issues []*issue.Issue
	if err = cur.All(ctx, &issues); err != nil {
		return nil, errors.Wrap(err, "repository.Issue.GetIssuesByProject")
	}

	return issues, nil
}

//GetIssuesByUser returns all issues in which the given user is present
func (r *mongoRepository) GetIssuesByUser(userId string) ([]*issue.Issue, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	//get the collection
	collection := r.getIssuesCollection()

	//construct the filter
	// idPrimitive, err := primitive.ObjectIDFromHex(userId)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "repository.Issue.GetIssuesByUser")
	// }
	filter := bson.M{"user": userId}

	//query the db
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Issue.GetIssuesByUser")
	}
	defer cur.Close(ctx)
	var issues []*issue.Issue
	if err = cur.All(ctx, &issues); err != nil {
		return nil, errors.Wrap(err, "repository.Issue.GetIssuesByUser")
	}

	return issues, nil
}

//CreateIssue adds a new issue to the database
func (r *mongoRepository) CreateIssue(issueStrut *issue.Issue) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	//get the collection
	collection := r.getIssuesCollection()

	//insert into the db
	_, err := collection.InsertOne(ctx, issueStrut)
	if err != nil {
		return errors.Wrap(err, "repository.Issue.CreateIssue")
	}

	return nil
}

//PutIssue updates an existing issue (completely replaces it)
func (r *mongoRepository) PutIssue(issueStrut *issue.Issue) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	//get the collection
	collection := r.getIssuesCollection()

	//construct the filter (to identify the issue to update)
	idPrimitive, err := primitive.ObjectIDFromHex(issueStrut.Id)
	if err != nil {
		return errors.Wrap(err, "repository.Issue.PutIssue")
	}
	filter := bson.M{"_id": idPrimitive}

	//don't include the Id in the object
	issueStrut.Id = ""

	//update the db
	_, err = collection.ReplaceOne(ctx, filter, issueStrut)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.Wrap(issue.ErrIssueNotFound, "repository.Issue.PutIssue")
		}
		return errors.Wrap(err, "repository.Issue.PutIssue")
	}
	return nil
}

//DeleteIssue removes a issue from the database
func (r *mongoRepository) DeleteIssue(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	//get the collection
	collection := r.getIssuesCollection()

	//construct the filter (to identify the issue to remove)
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Wrap(err, "repository.Issue.DeleteIssue")
	}
	filter := bson.M{"_id": idPrimitive}

	//remove from the db
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.Wrap(err, "repository.Issue.DeleteIssue")
	}
	return nil
}
