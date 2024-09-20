## Task Manager - Command Line Tool

A simple command-line task manager written in Go that allows you to add, remove, and list tasks.

### Installation

1. Make sure you have Go installed on your system
2. Clone or download the script files
3. Navigate to the directory containing the files
4. Build the application:
   ```bash
   go build -o task main.go
   ```

### Usage

The application provides the following commands:

#### Add a Task
```bash
./task add "Buy groceries"
./task add Complete project report
```
Adds a new task with the specified description.

#### List Tasks
```bash
./task list
```
Displays all current tasks with their IDs and descriptions.

#### Remove a Task
```bash
./task remove 1
```
Removes the task with the specified ID (use the ID shown in the list command).

#### Show Help
```bash
./task help
```
Displays usage instructions.

### Features

- **Persistent Storage**: Tasks are automatically saved to `tasks.txt` file
- **Automatic ID Generation**: Each task gets a unique incremental ID
- **Simple Interface**: Easy-to-use command-line interface
- **Error Handling**: Provides helpful error messages for invalid inputs

### File Structure

- `main.go` - Main application code
- `go.mod` - Go module definition
- `tasks.txt` - Auto-generated file storing tasks (created automatically)

### Example Session

```bash
# Add some tasks
$ ./task add "Learn Go programming"
Added task: [1] Learn Go programming

$ ./task add "Write documentation"
Added task: [2] Write documentation

# List tasks
$ ./task list
Tasks:
[1] Learn Go programming
[2] Write documentation

# Remove a task
$ ./task remove 1
Removed task: [1] Learn Go programming

# List again
$ ./task list
Tasks:
[2] Write documentation
```

### Notes

- Tasks are stored in a simple text file (`tasks.txt`) in the same directory
- The application automatically loads existing tasks on startup
- Task IDs are automatically assigned and incremented
- The application must be built before use (see Installation section)