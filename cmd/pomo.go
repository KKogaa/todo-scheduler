package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pomoCmd = &cobra.Command{
	Use:   "pomo",
	Short: "test creating pomo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Test")
	},
}

func init() {
	rootCmd.AddCommand(pomoCmd)
}
