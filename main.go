package main

// import necessary packages
import (
    "fmt"
    "os"

	clipboard "github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
	cli "github.com/urfave/cli/v2"
)

// define style
var style = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#FAFAFA")).
    Background(lipgloss.Color("#7D56F4")).
    PaddingTop(2).
    PaddingLeft(4).
    Width(100)


// define model
type model struct {
    choices  []string
    cursor   int
    selected int
}

// define our application’s initial state. 
// In this case, we’re defining a function to return our initial model
func initialModel() model {
    return model{
        choices: []string{"Create tar", "Extract tar", "List contents"},
        selected: -1,
    }
}

// initalization method
func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
    return nil
}

// `Cmd`s run asynchronously in a goroutine. 
// The `Msg`` they return is collected and sent to our update function for handling.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

	// Is it a key press?
    case tea.KeyMsg:
		// Cool, what was the actual key pressed?
        switch msg.String() {

		// The "up" key moves the cursor up
        case "up":
            if m.cursor > 0 {
                m.cursor--
            }

		// The "down" key moves the cursor down
        case "down":
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            }

		// The "enter" key toggles
        // the selected state for the item that the cursor is pointing at.
        case "enter":
            m.selected = m.cursor
            return m, tea.Quit
		
		// q or Ctrl+c exits.
        case "q", "ctrl+c":
            return m, tea.Quit
        }
    }

	// If we happen to get any other messages, don't do anything.
    return m, nil
}

// View method renders our UI
// We look at the model in its current state and use it to return a `string`
func (m model) View() string {
	// The header
	s := "A TUI for the tar command: generate tar commands with the bat of an eye\n\n"

	// Iterate over our choices
    for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor!
        }

		// Render the row
        s += fmt.Sprintf("%s %s\n", cursor, choice)
    }
	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
    return style.Render(s)
}

func runCLI() {
    app := &cli.App{
        Name: "tar-cli",
        Usage: "A TUI for the tar command: generate tar commands with the bat of an eye",
        Action: func(c *cli.Context) error {
			// We pass our initial model to `tea.NewProgram` and let it rip
            p := tea.NewProgram(initialModel())
            m, err := p.StartReturningModel()
            if err != nil {
                return err
            }

            model := m.(model)
            if model.selected != -1 {
                var cmd string
                switch model.choices[model.selected] {
                case "Create tar":
                    cmd = "tar -cvf archive.tar *"
                case "Extract tar":
                    cmd = "tar -xvf *.tar"
                case "List contents":
                    cmd = "tar -tvf *.tar"
                }

				// save command to clipboard
                clipboard.WriteAll(cmd)
                fmt.Println(style.Render("Command copied to clipboard:", cmd))
            }
            return nil
        },
    }

	// error handling
    err := app.Run(os.Args)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

func main() {
    runCLI()
}
