package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func otherWordFunc(cmd *cobra.Command, args []string) {
	if len(args) != countFlag {
		fmt.Println("bad input dude")
		return
	}

	println(countFlag)
}
