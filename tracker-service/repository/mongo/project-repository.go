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
func (r *mongoRepository) GetProjectsAll(ctx context.Context) ([]*project.Project, error) {
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
func (r *mongoRepository) GetProjectById(ctx context.Context, id string) (*project.Project, error) {
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
func (r *mongoRepository) GetProjectsByUser(ctx context.Context, userId string) ([]*project.Project, error) {
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
func (r *mongoRepository) CreateProject(ctx context.Context, projectStrut *project.Project) error {
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
func (r *mongoRepository) DeleteProject(ctx context.Context, id string) error {
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

//AddUser adds a user to the given project
func (r *mongoRepository) AddUser(ctx context.Context, projectId string, userId string) error {
	//get the collection
	collection := r.getProjectsCollection()

	//construct the filter (to identify the project to update)
	idPrimitive, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return errors.Wrap(err, "repository.Project.AddUser")
	}
	filter := bson.M{"_id": idPrimitive}

	//construct the update statement
	update := bson.D{
		//we'll use addToSet here, although push should work just as well
		{"$addToSet", bson.D{{"users", userId}}},
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
func (r *mongoRepository) RemoveUser(ctx context.Context, projectId string, userId string) error {
	//get the collection
	collection := r.getProjectsCollection()

	//construct the filter (to identify the project to update)
	idPrimitive, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return errors.Wrap(err, "repository.Project.RemoveUser")
	}
	filter := bson.M{"_id": idPrimitive}

	//construct the update statement
	update := bson.D{
		{"$pull", bson.D{{"users", userId}}},
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
