package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/luiztapioca/todo-cli/todo"
)

func main() {
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	listAll := listCmd.Bool("all", false, "List all tasks")
	listPendent := listCmd.Bool("pendent", false, "List all pendent tasks")
	listCompleted := listCmd.Bool("completed", false, "List all completed tasks")

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	addTitle := addCmd.String("title", "", "Task title")
	addID := addCmd.Uint("id", 0, "Task ID")
	addStatus := addCmd.Bool("completed", false, "Task status")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	deleteID := deleteCmd.Uint("id", 0, "Delete a task")

	completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)

	completeID := completeCmd.Uint("id", 0, "Complete a task")

	if len(os.Args) < 2 {
		log.Fatal("Esperando argumentos 'add', 'list', 'delete'...")
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if *addTitle == "" {
			fmt.Println("You have to name the task...")
			addCmd.PrintDefaults()
			return
		}

		if *addID == 0 {
			fmt.Println("You have to give an valid id...")
			addCmd.PrintDefaults()
			return
		}

		newTask := todo.Task{
			ID:        *addID,
			Title:     *addTitle,
			Completed: *addStatus,
		}

		fmt.Printf("Adding task '%s' on id %d with status %s\n", newTask.Title, newTask.ID, newTask.Status())
		task, err := todo.SaveTask(newTask)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Task '%s' saved successfully!\n", task.Title)

	case "list":
		listCmd.Parse(os.Args[2:])
		if *listAll {
			fmt.Println("Listing all tasks...")
			tasks, err := todo.GetTasks()

			if err != nil {
				log.Fatal(err)
			}

			for _, task := range tasks {
				fmt.Println(task.String())
			}

		} else if *listPendent {
			fmt.Println("Listing pendent tasks...")
			tasks, err := todo.GetTasks()

			if err != nil {
				log.Fatal(err)
			}

			for _, task := range tasks {
				if !task.Completed {
					fmt.Println(task.String())
				}
			}
		} else if *listCompleted {
			fmt.Println("Listing completed tasks...")
			tasks, err := todo.GetTasks()

			if err != nil {
				log.Fatal(err)
			}

			for _, task := range tasks {
				if task.Completed {
					fmt.Println(task.String())
				}
			}
		}
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if *deleteID == 0 {
			fmt.Println("You have to inform a valid task id...")
			addCmd.PrintDefaults()
			return
		}

		fmt.Printf("Deleting task %d...\n", *deleteID)
		if err := todo.DeleteTask(*deleteID); err != nil {
			log.Fatalf("It was not possible to delete task %d.", *deleteID)
		}
		fmt.Printf("Task %d deleted successfully!\n", *deleteID)
	case "complete":
		completeCmd.Parse(os.Args[2:])
		if *completeID == 0 {
			fmt.Println("You have to inform a valid task id...")
			completeCmd.PrintDefaults()
			return
		}

		fmt.Printf("Completing task %d...\n", *completeID)
		tasks, err := todo.GetTasks()

		if err != nil {
			log.Fatal(err)
		}

		index := slices.IndexFunc(tasks, func(t todo.Task) bool {
			return t.ID == *completeID
		})

		if index != -1 {
			task, err := todo.SaveTask(todo.Task{
				ID:        *completeID,
				Title:     tasks[index].Title,
				Completed: true,
			})

			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Task '%s' completed successfully!\n", task.Title)
			return
		}

		log.Fatalf("Task with id %d not found...", *completeID)
	}
}
