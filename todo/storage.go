package todo

import (
	"encoding/json"
	"io"
	"os"
)

func GetTasks() ([]Task, error) {
	file, err := os.Open("tasks.json")

	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		if err == io.EOF {
			return []Task{}, nil
		}
		return nil, err
	}

	return tasks, nil
}

func SaveTask(task Task) (Task, error) {
	tasks, err := GetTasks()

	if err != nil {
		return task, err
	}

	tasks = append(tasks, task)

	file, err := os.OpenFile("tasks.json", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return task, err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(&tasks); err != nil {
		return task, err
	}

	return task, nil
}
