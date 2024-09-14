package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID     int
	Name   string
	Status string
}

var tasks []Task
var nextID = 1

func main() {
	loadTasks()
	
	fmt.Println("Simple Task Manager")
	fmt.Println("Commands: add, list, remove, quit")
	
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		
		input := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(input)
		
		if len(parts) == 0 {
			continue
		}
		
		command := strings.ToLower(parts[0])
		
		switch command {
		case "add":
			if len(parts) > 1 {
				taskName := strings.Join(parts[1:], " ")
				addTask(taskName)
			} else {
				fmt.Println("Usage: add <task description>")
			}
			
		case "list":
			listTasks()
			
		case "remove":
			if len(parts) > 1 {
				id, err := strconv.Atoi(parts[1])
				if err != nil {
					fmt.Println("Invalid ID. Please provide a number.")
				} else {
					removeTask(id)
				}
			} else {
				fmt.Println("Usage: remove <task_id>")
			}
			
		case "quit", "exit":
			saveTasks()
			fmt.Println("Goodbye!")
			return
			
		default:
			fmt.Println("Unknown command. Available commands: add, list, remove, quit")
		}
	}
}

func addTask(name string) {
	task := Task{
		ID:     nextID,
		Name:   name,
		Status: "Pending",
	}
	tasks = append(tasks, task)
	nextID++
	fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	
	fmt.Println("\nTasks:")
	fmt.Println("ID\tStatus\t\tDescription")
	fmt.Println("--\t------\t\t-----------")
	for _, task := range tasks {
		fmt.Printf("%d\t%-10s\t%s\n", task.ID, task.Status, task.Name)
	}
	fmt.Println()
}

func removeTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Task %d removed successfully\n", id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found\n", id)
}

func saveTasks() {
	// Simple file-based persistence
	file, err := os.Create("tasks.txt")
	if err != nil {
		fmt.Println("Warning: Could not save tasks to file")
		return
	}
	defer file.Close()
	
	for _, task := range tasks {
		line := fmt.Sprintf("%d|%s|%s\n", task.ID, task.Name, task.Status)
		file.WriteString(line)
	}
}

func loadTasks() {
	file, err := os.Open("tasks.txt")
	if err != nil {
		// File doesn't exist, start fresh
		return
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) >= 3 {
			id, err := strconv.Atoi(parts[0])
			if err != nil {
				continue
			}
			
			task := Task{
				ID:     id,
				Name:   parts[1],
				Status: parts[2],
			}
			tasks = append(tasks, task)
			
			if id >= nextID {
				nextID = id + 1
			}
		}
	}
}