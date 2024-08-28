# TaskManager

TaskManager is a simple CLI-based task management application written in Go. This tool allows you to add, list, complete, and delete tasks directly from the command line.

## Features

- **Add Tasks**: Easily add tasks with a name and automatically assigned unique ID.
- **List Tasks**: View all your tasks, with their status (completed or not).
- **Complete Tasks**: Mark tasks as completed by specifying their ID.
- **Delete Tasks**: Remove tasks from your list by specifying their ID.
- **Makefile Shortcuts**: Easily run commands using the provided `Makefile`.

## Project Structure

```
taskmanager/
├── cmd/
│   └── taskmanager/
│       ├── main.go                # Entry point for the CLI application
│       ├── handler.go             # Contains the logic for handling CLI commands
│       ├── storage.go             # Handles loading and saving tasks to a JSON file
│       ├── task.go                # Defines the Task struct and related functionality
│       ├── handler_test.go        # Unit tests for handler.go
│       ├── storage_test.go        # Unit tests for storage.go
│       ├── task_test.go           # Unit tests for task.go
├── .gitignore
├── go.mod                         # Go module file
├── README.md                      # This file
├── Makefile                       # Makefile with shortcuts for common tasks
└── tasks.json                     # File where tasks are stored (generated after first run)
```

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/yourusername/taskmanager.git
   cd taskmanager/cmd/taskmanager
   ```

2. **Install dependencies**:

   Go will automatically handle dependencies based on the `go.mod` file. Just run your commands, and Go will fetch the necessary packages.

## Usage

### Using the Makefile Shortcuts

The `Makefile` provides several shortcuts to make common operations easier:

#### 1. Add a Task

Add a new task by providing a task name:

```bash
make add task="Complete project documentation"
```

#### 2. List All Tasks

View all tasks:

```bash
make list
```

#### 3. Complete a Task

Mark a task as completed by specifying its ID:

```bash
make complete id=1
```

#### 4. Delete a Task

Delete a task by specifying its ID:

```bash
make delete id=1
```

#### 5. Build the Application

If you need to build the application manually:

```bash
make build
```

#### 6. Run the Application

Run the application after building it:

```bash
make run
```

#### 7. Clean Up

Remove the generated binary:

```bash
make clean
```

### Running the Application Directly

Alternatively, you can run the application directly without using the `Makefile`:

#### 1. Add a Task

```bash
go run main.go add "Complete project documentation"
```

#### 2. List All Tasks

```bash
go run main.go list
```

#### 3. Complete a Task

```bash
go run main.go complete 1
```

#### 4. Delete a Task

```bash
go run main.go delete 1
```

## Testing

To run the tests, navigate to the `cmd/taskmanager` directory and run:

```bash
make test
```

This will run all unit tests for the project.

## Author

- **misonikomipan** - [GitHub](https://github.com/misonikomipan)
