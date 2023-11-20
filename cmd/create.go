package cmd

import (
	"time"

	"github.com/KKogaa/todo-scheduler/application"
	"github.com/KKogaa/todo-scheduler/core"
	"github.com/KKogaa/todo-scheduler/infra/repositories/local"
	// tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// type model struct {
// 	choices  []string         // items on the to-do list
// 	cursor   int              // which to-do list item our cursor is pointing at
// 	selected map[int]struct{} // which to-do items are selected
// }

// func initialModel() model {
// 	return model{
// 		// Our to-do list is a grocery list
// 		choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

// 		// A map which indicates which choices are selected. We're using
// 		// the  map like a mathematical set. The keys refer to the indexes
// 		// of the `choices` slice, above.
// 		selected: make(map[int]struct{}),
// 	}
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {

// 	// Is it a key press?
// 	case tea.KeyMsg:

// 		// Cool, what was the actual key pressed?
// 		switch msg.String() {

// 		// These keys should exit the program.
// 		case "ctrl+c", "q":
// 			return m, tea.Quit

// 		// The "up" and "k" keys move the cursor up
// 		case "up", "k":
// 			if m.cursor > 0 {
// 				m.cursor--
// 			}

// 		// The "down" and "j" keys move the cursor down
// 		case "down", "j":
// 			if m.cursor < len(m.choices)-1 {
// 				m.cursor++
// 			}

// 		// The "enter" key and the spacebar (a literal space) toggle
// 		// the selected state for the item that the cursor is pointing at.
// 		case "enter", " ":
// 			_, ok := m.selected[m.cursor]
// 			if ok {
// 				delete(m.selected, m.cursor)
// 			} else {
// 				m.selected[m.cursor] = struct{}{}
// 			}
// 		}
// 	}

// 	// Return the updated model to the Bubble Tea runtime for processing.
// 	// Note that we're not returning a command.
// 	return m, nil
// }

// func (m model) View() string {
// 	// The header
// 	s := "What should we buy at the market?\n\n"

// 	// Iterate over our choices
// 	for i, choice := range m.choices {

// 		// Is the cursor pointing at this choice?
// 		cursor := " " // no cursor
// 		if m.cursor == i {
// 			cursor = ">" // cursor!
// 		}

// 		// Render the row
// 		s += fmt.Sprintf("%s %s\n", cursor, choice)
// 	}

// 	// The footer
// 	s += "\nPress q to quit.\n"

// 	// Send the UI for rendering
// 	return s
// }

// func (m model) Init() tea.Cmd {
// 	// Just return `nil`, which means "no I/O right now, please."
// 	return nil
// }

var StringToTaskStatus = map[string]core.TaskStatus{
	"Pending":    core.Pending,
	"Doing":      core.Doing,
	"Done":       core.Done,
	"Terminated": core.Terminated,
	"Cancelled":  core.Cancelled,
}

var StringToTaskDifficulty = map[string]core.TaskDifficulty{
	"Easy":     core.Easy,
	"Medium":   core.Medium,
	"Hard":     core.Hard,
	"VeryHard": core.VeryHard,
	"Unknown":  core.Unknown,
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Add a new tast",
	Run: func(cmd *cobra.Command, args []string) {
		repo := local.NewTaskRepository()
		usecase := application.NewCreateUsecase(repo)

		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		status, _ := cmd.Flags().GetString("status")
		est, _ := cmd.Flags().GetInt8("est")
		difficulty, _ := cmd.Flags().GetString("diff")

		estimatedTime := time.Minute * time.Duration(est)

		usecase.Execute(description, name, StringToTaskStatus[status],
			estimatedTime, StringToTaskDifficulty[difficulty])

		// p := tea.NewProgram(initialModel())
		// if _, err := p.Run(); err != nil {
		// 	fmt.Printf("Alas, there's been an error: %v", err)
		// 	os.Exit(1)
		// }
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().String("name", "", "Name of the task")
	createCmd.PersistentFlags().String("desc", "", "Description of the task")
	createCmd.PersistentFlags().String("status", "", "Status of the task")
	createCmd.PersistentFlags().String("est", "", "Estimated time of the task")
	createCmd.PersistentFlags().String("diff", "", "Difficulty time of the task")
}
