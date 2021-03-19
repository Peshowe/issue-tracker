package issue

type Issue struct {
	Id string `bson:"_id,omitempty"`

	//Name of issue between 2 and 40 characters
	Name string `bson:"name,omitempty" validate:"gte=2 & lte=40"`

	//Description (can be empty and shouldn't be more than 500 chars)
	Desc string `bson:"desc,omitempty" validate:"emtpy=true & lte=500"`

	//Type of the issue (can be one of "bug", "feature" or "adhoc")
	IssueType string `bson:"issue_type,omitempty" validate:"one_of=bug,feature,adhoc"`

	//Current status of the issue (can be one of "to do", "in progress" or "done")
	Status string `bson:"status,omitempty" validate:"one_of=to do,in progress,done"`

	//Optional error trace from the encountered bug
	BugTrace string `bson:"bug_trace,omitempty"`

	//Asignee of the issue
	User string `bson:"user,omitempty"`

	//Project to which the issue belongs
	Project string `bson:"project,omitempty"`

	CreatedOn int64 `bson:"created_on,omitempty"`

	LastModifiedOn int64 `bson:"last_modified_on,omitempty"`
}
