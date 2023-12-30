BINARY_NAME=tar-cli

# golang parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get

# go build the binary executable tar-cli
build: 
	$(GOBUILD) -o $(BINARY_NAME) main.go

# go run program
run: 
	$(GORUN) main.go

# go clean binaries
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# go get dependencies
deps:
	$(GOGET) github.com/atotto/clipboard
	$(GOGET) github.com/charmbracelet/bubbletea
	$(GOGET) github.com/charmbracelet/lipgloss
	$(GOGET) github.com/urfave/cli/v2

# do all of the above
all: deps build run
