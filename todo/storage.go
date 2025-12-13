package todo

import (
	"encoding/json"
	"io"
	"os"
	"slices"
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

func persistTasks(tasks []Task) error {
	file, err := os.OpenFile("tasks.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(&tasks)
}

func SaveTask(task Task) (Task, error) {
	tasks, err := GetTasks()

	if err != nil {
		return task, err
	}

	index := slices.IndexFunc(tasks, func(t Task) bool {
		return t.ID == task.ID
	})

	if index == -1 {
		tasks = append(tasks, task)
	} else {
		tasks[index] = task
	}

	return task, persistTasks(tasks)
}

func DeleteTask(id uint) error {
	tasks, err := GetTasks()

	if err != nil {
		return err
	}

	index := slices.IndexFunc(tasks, func(t Task) bool {
		return t.ID == id
	})

	tasks = slices.Delete(tasks, index, index+1)

	return persistTasks(tasks)
}
