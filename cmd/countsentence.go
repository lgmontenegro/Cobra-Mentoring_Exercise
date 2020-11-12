package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(countSentenceCmd)
}

var countSentenceCmd = &cobra.Command{
	Use:   "countsentence",
	Short: "countsentence counts a god damm sentence",
	Long: `countsentence counts a god damm sentence
	countsentence counts a god damm sentence
	countsentence counts a god damm sentence`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(len(args[0]))
	},
}
