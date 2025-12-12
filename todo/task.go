package todo

type Task struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func newTask(id uint, title string, completed bool) *Task {
	t := Task{ID: id, Title: title, Completed: completed}
	return &t
}
