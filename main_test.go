/*************************************************************************************************************************
* Name: Malik Freeman
* Date: 1/19/2025
* Course: SSE 662 - Design, Maintenance, and Quality
* Assignment: Module 1 Programming Task
* File Name: main_test.go
* Description: This file implements  unit test to test
*              - Add Task: Allow users to add new tasks with a brief description.
*              - List Tasks: Display all existing tasks with their status (Pending/Done).
*              - Mark Task as Done: Enable users to mark a specific task as done by its task number.
*              - Delete Task: Permit deletion of a task by its task number.
*
***********************************************************************************************************************/

package main

import (
	"testing"
)

// Test addTask function
func TestAddTask(t *testing.T) {
	// Clear existing tasks before testing
	tasks = []Task{}

	// Add a task
	addTask([]string{"Test task"})

	// Check if the task was added
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
	if tasks[0].Description != "Test task" {
		t.Errorf("Expected task description to be 'Test task', got '%s'", tasks[0].Description)
	}
	if tasks[0].IsDone != false {
		t.Errorf("Expected task status to be false (Pending), got true")
	}
}

// Test listTasks function
func TestListTasks(t *testing.T) {
	// Clear existing tasks before testing
	tasks = []Task{}

	// Add tasks
	addTask([]string{"Task 1"})
	addTask([]string{"Task 2"})

	// List tasks and check if the correct number of tasks are displayed
	// You would need to capture the output of the function for testing, for example using a custom writer.
	// Here's a basic example without capturing output.
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
}

// Test markTaskDone function
func TestMarkTaskDone(t *testing.T) {
	// Clear existing tasks before testing
	tasks = []Task{}

	// Add a task
	addTask([]string{"Task to be done"})

	// Mark task as done
	markTaskDone([]string{"1"})

	// Check if the task is marked as done
	if tasks[0].IsDone != true {
		t.Errorf("Expected task status to be true (Done), got false")
	}
}

// Test deleteTask function
func TestDeleteTask(t *testing.T) {
	// Clear existing tasks before testing
	tasks = []Task{}

	// Add tasks
	addTask([]string{"Task 1"})
	addTask([]string{"Task 2"})

	// Delete the first task
	deleteTask([]string{"1"})

	// Check if the task was deleted
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
	if tasks[0].Description != "Task 2" {
		t.Errorf("Expected remaining task to be 'Task 2', got '%s'", tasks[0].Description)
	}
}
