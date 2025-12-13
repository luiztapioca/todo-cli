package todo

import "fmt"

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

func (t Task) String() string {
	return fmt.Sprintf("%d -- %s [%s]", t.ID, t.Title, t.Status())
}
