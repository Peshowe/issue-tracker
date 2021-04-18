package project

type Project struct {
	Id string `bson:"_id,omitempty" json:"id,omitempty"`

	//Name of project between 2 and 40 characters
	Name string `bson:"name,omitempty" json:"name,omitempty" validate:"gte=2 & lte=40"`

	// //Issues inside a project
	// Issues []string `bson:"issues,omitempty" json:"issues,omitempty"`

	//Users assigned to the project
	Users []string `bson:"users,omitempty" json:"users,omitempty"`

	CreatedOn int64 `bson:"created_on,omitempty" json:"created_on,omitempty"`
}
