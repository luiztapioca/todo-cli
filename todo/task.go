package todo

type Task struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (t Task) Status() string {
	if t.Completed {
		return "Completed"
	} else {
		return "Not completed"
	}
}
