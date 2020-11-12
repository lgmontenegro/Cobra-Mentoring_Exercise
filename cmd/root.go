package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Several CMD tools",
	Short: "Several CMD tools Several CMD tools Several CMD tools",
	Long: `Several CMD tools Several CMD tools Several CMD tools
	Several CMD tools Several CMD tools Several CMD tools
	Several CMD tools Several CMD tools Several CMD tools`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Several CMD tools")
	},
}

// Execute Runs the main code
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
