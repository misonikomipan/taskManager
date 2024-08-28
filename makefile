# OS-specific settings
ifeq ($(OS),Windows_NT)
    EXT = .exe
else
    EXT =
endif

# Binary name
BINARY = taskmanager$(EXT)

# Directories
CMD_DIR = ./cmd/taskmanager

# Default target
all: build

# Build the binary
build:
	go build -o $(BINARY) $(CMD_DIR)

# Run the application
run: build
	./$(BINARY)

# Run the tests
test:
	go test ./cmd/taskmanager/...

# Add a new task
add: build
	./$(BINARY) add "$(task)"

# List all tasks
list: build
	./$(BINARY) list

# Complete a task
complete: build
	./$(BINARY) complete $(id)

# Delete a task
delete: build
	./$(BINARY) delete $(id)

# Clean up generated files
.PHONY: clean
clean:
	rm -f $(BINARY)