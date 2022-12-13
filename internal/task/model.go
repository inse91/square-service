package task

type Task struct {
	ID          string   `json:"id" bson:"_id,omitempty"`
	Description string   `json:"description" bson:"description"`
	Tags        []string `json:"tags" bson:"tags"`
	Priority    string   `json:"priority" bson:"priority"`
	// TODO DueDate     time.Time `json:"due"`
}

type CreateTaskDTO struct {
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Priority    string   `json:"priority"`
}

func NewTask() *Task {
	return &Task{}
}

var exampleTask = Task{
	Description: "this is Task example2",
	Tags:        []string{"example"},
	Priority:    "low",
}

func GetExampleTask() *Task {
	return &exampleTask
}
