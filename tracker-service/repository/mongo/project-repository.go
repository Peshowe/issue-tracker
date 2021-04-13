package mongo

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Peshowe/issue-tracker/tracker-service/tracker/project"
	
)

//getProjectsCollection gets a reference to the projects collection in the db
func (r *mongoRepository) getProjectsCollection() *mongo.Collection {
	return r.client.Database(r.database).Collection("projects")
}

//GetProjectsAll returns all projects in the database
func (r *mongoRepository) GetProjectsAll() ([]*project.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	//get the collection
	collection := r.getProjectsCollection()

	//query everything
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.Wrap(err, "repository.Project.GetProjectsAll")
	}
	defer cur.Close(ctx)
	var projects []*project.Project
	if err = cur.All(ctx, &projects); err != nil {
		return nil, errors.Wrap(err, "repository.Project.GetProjectsAll")
	}

	return projects, nil
}

//GetProjectById returns a single project that matches the given id
func (r *mongoRepository) GetProjectById(id string) (*project.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	//get the collection
	collection := r.getProjectsCollection()

	//contruct the filter
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Project.GetProjectById")
	}
	filter := bson.M{"_id": idPrimitive}

	//query the db
	projectObj := &project.Project{}
	err = collection.FindOne(ctx, filter).Decode(&projectObj)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(project.ErrProjectNotFound, "repository.Project.GetProjectById")
		}
		return nil, errors.Wrap(err, "repository.Project.GetProjectById")
	}
	return projectObj, nil
}

//GetProjectsByUser returns all projects in which the given user is present
func (r *mongoRepository) GetProjectsByUser(userId string) ([]*project.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	//get the collection
	collection := r.getProjectsCollection()

	//construct the filter
	// idPrimitive, err := primitive.ObjectIDFromHex(userId)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "repository.Project.GetProjectsByUser")
	// }
	filter := bson.D{
		{"users", bson.D{{"$all", bson.A{userId}}}},
	}

	//query the db
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Project.GetProjectsByUser")
	}
	defer cur.Close(ctx)
	var projects []*project.Project
	if err = cur.All(ctx, &projects); err != nil {
		return nil, errors.Wrap(err, "repository.Project.GetProjectsByUser")
	}

	return projects, nil
}

//CreateProject adds a new project to the database
func (r *mongoRepository) CreateProject(projectStrut *project.Project) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	//get the collection
	collection := r.getProjectsCollection()

	//insert into the db
	_, err := collection.InsertOne(ctx, projectStrut)
	if err != nil {
		return errors.Wrap(err, "repository.Project.CreateProject")
	}
	return nil
}

//DeleteProject removes a project from the database
func (r *mongoRepository) DeleteProject(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	//get the collection
	collection := r.getProjectsCollection()

	//construct the filter (to identify the project to remove)
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Wrap(err, "repository.Project.DeleteProject")
	}
	filter := bson.M{"_id": idPrimitive}

	//remove from the db
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.Wrap(project.ErrProjectNotFound, "repository.Project.DeleteProject")
		}
		return errors.Wrap(err, "repository.Project.DeleteProject")
	}
	return nil
}

//AddIssue adds an issue to the given project
func (r *mongoRepository) AddIssue(projectId string, issueId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	//get the collection
	collection := r.getProjectsCollection()

	//construct the filter (to identify the project to update)
	idPrimitive, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return errors.Wrap(err, "repository.Project.AddIssue")
	}
	filter := bson.M{"_id": idPrimitive}

	//construct the update statement
	// idPrimitive, err = primitive.ObjectIDFromHex(issueId)
	// if err != nil {
	// 	return errors.Wrap(err, "repository.Project.AddIssue")
	// }
	update := bson.D{
		//we'll use addToSet here, although push should work just as well
		{"$addToSet", bson.D{{"issues", issueId}}},
	}

	//update the db
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.Wrap(project.ErrProjectNotFound, "repository.Project.UpdateOne")
		}
		return errors.Wrap(err, "repository.Project.AddIssue")
	}
	return nil
}

//RemoveIssue removes an issue from the given project
func (r *mongoRepository) RemoveIssue(projectId string, issueId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	//get the collection
	collection := r.getProjectsCollection()

	//construct the filter (to identify the project to update)
	idPrimitive, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return errors.Wrap(err, "repository.Project.RemoveIssue")
	}
	filter := bson.M{"_id": idPrimitive}

	//construct the update statement
	idPrimitive, err = primitive.ObjectIDFromHex(issueId)
	if err != nil {
		return errors.Wrap(err, "repository.Project.RemoveIssue")
	}
	update := bson.D{
		{"$pull", bson.D{{"issues", idPrimitive}}},
	}

	//update the db
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.Wrap(project.ErrProjectNotFound, "repository.Project.RemoveIssue")
		}
		return errors.Wrap(err, "repository.Project.RemoveIssue")
	}
	return nil
}

//AddUser adds a user to the given project
func (r *mongoRepository) AddUser(projectId string, userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	//get the collection
	collection := r.getProjectsCollection()

	//construct the filter (to identify the project to update)
	idPrimitive, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return errors.Wrap(err, "repository.Project.AddUser")
	}
	filter := bson.M{"_id": idPrimitive}

	//construct the update statement
	idPrimitive, err = primitive.ObjectIDFromHex(userId)
	if err != nil {
		return errors.Wrap(err, "repository.Project.AddUser")
	}
	update := bson.D{
		//we'll use addToSet here, although push should work just as well
		{"$addToSet", bson.D{{"users", idPrimitive}}},
	}

	//update the db
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.Wrap(project.ErrProjectNotFound, "repository.Project.AddUser")
		}
		return errors.Wrap(err, "repository.Project.AddUser")
	}
	return nil
}

//RemoveUser removes a user from the given project
func (r *mongoRepository) RemoveUser(projectId string, userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	//get the collection
	collection := r.getProjectsCollection()

	//construct the filter (to identify the project to update)
	idPrimitive, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return errors.Wrap(err, "repository.Project.RemoveUser")
	}
	filter := bson.M{"_id": idPrimitive}

	//construct the update statement
	idPrimitive, err = primitive.ObjectIDFromHex(userId)
	if err != nil {
		return errors.Wrap(err, "repository.Project.RemoveUser")
	}
	update := bson.D{
		{"$pull", bson.D{{"users", idPrimitive}}},
	}

	//update the db
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.Wrap(project.ErrProjectNotFound, "repository.Project.RemoveUser")
		}
		return errors.Wrap(err, "repository.Project.RemoveUser")
	}
	return nil
}
