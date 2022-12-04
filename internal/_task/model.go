package task

type task struct {
	ID          int      `json:"id"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Priority    string   `json:"priority"`
	//DueDate     time.Time `json:"due"`
}

func NewTask() *task {
	return &task{}
}

var exampleTask = task{
	ID:          999,
	Description: "this is _task example",
	Tags:        []string{"example"},
	Priority:    "low",
}

func GetExampleTask() *task {
	return &exampleTask
}
