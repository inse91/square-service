package task

type Task struct {
	ID          string   `json:"id" bson:"__id,omitempty"`
	Description string   `json:"description" bson:"description"`
	Tags        []string `json:"tags" bson:"tags"`
	Priority    string   `json:"priority" bson:"priority"`
	//DueDate     time.Time `json:"due"`
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
	ID:          "999",
	Description: "this is Task example",
	Tags:        []string{"example"},
	Priority:    "low",
}

func GetExampleTask() *Task {
	return &exampleTask
}
