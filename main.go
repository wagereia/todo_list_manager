package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID   int
	Name string
}

var tasks []Task
var nextID = 1

func main() {
	loadTasks()
	
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task description")
			return
		}
		taskName := strings.Join(os.Args[2:], " ")
		addTask(taskName)
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task ID to remove")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid task ID")
			return
		}
		removeTask(id)
	case "list":
		listTasks()
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
	}

	saveTasks()
}

func addTask(name string) {
	task := Task{
		ID:   nextID,
		Name: name,
	}
	tasks = append(tasks, task)
	nextID++
	fmt.Printf("Added task: [%d] %s\n", task.ID, task.Name)
}

func removeTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Removed task: [%d] %s\n", task.ID, task.Name)
			return
		}
	}
	fmt.Printf("Task with ID %d not found\n", id)
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Println("Tasks:")
	for _, task := range tasks {
		fmt.Printf("[%d] %s\n", task.ID, task.Name)
	}
}

func printUsage() {
	fmt.Println("Task Manager - Command Line Tool")
	fmt.Println("Usage:")
	fmt.Println("  task add <task description>    - Add a new task")
	fmt.Println("  task remove <task ID>          - Remove a task by ID")
	fmt.Println("  task list                      - List all tasks")
	fmt.Println("  task help                      - Show this help message")
}

func saveTasks() {
	file, err := os.Create("tasks.txt")
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, task := range tasks {
		_, err := writer.WriteString(fmt.Sprintf("%d|%s\n", task.ID, task.Name))
		if err != nil {
			fmt.Printf("Error writing task: %v\n", err)
			return
		}
	}
	writer.Flush()
}

func loadTasks() {
	file, err := os.Open("tasks.txt")
	if err != nil {
		if os.IsNotExist(err) {
			return // File doesn't exist yet, that's OK
		}
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) == 2 {
			id, err := strconv.Atoi(parts[0])
			if err != nil {
				continue
			}
			task := Task{
				ID:   id,
				Name: parts[1],
			}
			tasks = append(tasks, task)
			if id >= nextID {
				nextID = id + 1
			}
		}
	}
}