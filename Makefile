# Binary output name
BINARY_NAME=tar-cli

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get

# Build the binary executable
build: 
	$(GOBUILD) -o $(BINARY_NAME) main.go

# Run the program
run: 
	$(GORUN) main.go

# Clean up binaries
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Get dependencies
deps:
	$(GOGET) github.com/atotto/clipboard
	$(GOGET) github.com/charmbracelet/bubbletea
	$(GOGET) github.com/charmbracelet/lipgloss
	$(GOGET) github.com/urfave/cli/v2

# Include all from above
all: deps build run
