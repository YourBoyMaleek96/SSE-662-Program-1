/*************************************************************************************************************************
* Name: Malik Freeman
* Date: 1/19/2025
* Course: SSE 662 - Design, Maintenance, and Quality
* Assignment: Module 1 Programming Task
* File Name: main.go
* Description: This file implements a command-line To-Do List Manager in Go. The program
*              allows users to:
*              - Add Task: Allow users to add new tasks with a brief description.
*              - List Tasks: Display all existing tasks with their status (Pending/Done).
*              - Mark Task as Done: Enable users to mark a specific task as done by its task number.
*              - Delete Task: Permit deletion of a task by its task number.
*              - Help/Usage: Provide a help menu that outlines how to use each command.
***********************************************************************************************************************/

package main

//imports use to run the program
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Struct that defines Description as a string and IsDone as a boolean
type Task struct {
	Description string
	IsDone      bool
}

// tasks is a slice that holds the list of all to-do tasks.
var tasks []Task

// Main function
func main() {
	reader := bufio.NewReader(os.Stdin)

	// To Do List header
	fmt.Println("To-Do List Manager")
	fmt.Println("-------------------")
	fmt.Println("Type 'help' for usage instructions.")

	// Main loop for handling user commands.
	for {
		fmt.Print("\n> ")
		input, _ := reader.ReadString('\n')                 // Read user input
		command := strings.Fields(strings.TrimSpace(input)) // Split input into command and arguments

		if len(command) == 0 { // Ignore empty input
			continue
		}

		// Handle commands based on user input
		switch command[0] {
		case "add":
			addTask(command[1:])
		case "list":
			listTasks()
		case "done":
			markTaskDone(command[1:])
		case "delete":
			deleteTask(command[1:])
		case "help":
			showHelp()
		case "quit":
			fmt.Println("The program is exit, rerun program to start")
			return
		default:
			fmt.Println("Invalid command. Type 'help' for a list of valid commands.")
		}
	}
}

// addTask adds a new task to the to-do list.
// Arguments:
// - args: The task description split into words.
func addTask(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: add <task_description>")
		return
	}

	// Combine words into a single description and append to tasks.
	description := strings.Join(args, " ")
	tasks = append(tasks, Task{Description: description, IsDone: false})
	fmt.Printf(" The following task was added: \"%s\"\n", description)
}

// listTasks displays all tasks in the to-do list with their status.
func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("There are Nno tasks found, type add to add one.") // Handle case with no tasks
		return
	}

	fmt.Println("Tasks:")
	for i, task := range tasks {
		status := "Pending"
		if task.IsDone {
			status = "Done"
		}
		// Print task number, status, and description
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Description)
	}
}

// markTaskDone marks a task as completed based on its task number.
// Arguments:
// - args: The task number provided by the user.
func markTaskDone(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: done <task_number>")
		return
	}

	// Make sure the task is in the list that is getting marked done
	taskNumber, err := strconv.Atoi(args[0])
	if err != nil || taskNumber < 1 || taskNumber > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	// Mark the specified task as done
	tasks[taskNumber-1].IsDone = true
	fmt.Printf("Task %d is done: \"%s\".\n", taskNumber, tasks[taskNumber-1].Description)
}

// deleteTask removes a task from the to-do list based on its task number.
// Arguments:
// - args: The task number provided by the user.
func deleteTask(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: delete <task_number>")
		return
	}

	// Make sure the task is in the list that is getting deleted
	taskNumber, err := strconv.Atoi(args[0])
	if err != nil || taskNumber < 1 || taskNumber > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	// Remove the specified task from the list
	deletedTask := tasks[taskNumber-1]
	tasks = append(tasks[:taskNumber-1], tasks[taskNumber:]...)
	fmt.Printf("Deleted task %d: \"%s\"\n", taskNumber, deletedTask.Description)
}

// showHelp displays usage instructions for all supported commands.
func showHelp() {
	fmt.Println("Usage:")
	fmt.Println("  add <task_description>  - Add a new task with the given description.")
	fmt.Println("  list                    - List all tasks with their status (Pending/Done).")
	fmt.Println("  done <task_number>      - Mark the task with the given number as done.")
	fmt.Println("  delete <task_number>    - Delete the task with the given number.")
	fmt.Println("  help                    - Show this help message.")
	fmt.Println("  quit                    - Quit the application.")
}
