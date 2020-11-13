package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	otherWordsCmd.PersistentFlags().IntVar(&countFlag, "count", 1, "I dont know")
	rootCmd.AddCommand(otherWordsCmd)
}

var countFlag int

var otherWordsCmd = &cobra.Command{
	Use:   "otherwords",
	Short: "other words I don't know what to do",
	Long: `Normally, both your asses would be dead 
	as fucking fried chicken, but you happen to pull 
	this shit while I'm in a transitional period so I 
	don't wanna kill you, I wanna help you. But I can't 
	give you this case, it don't belong to me. Besides, 
	I've already been through too much shit this morning 
	over this case to hand it over to your dumb ass.`,
	Run: otherWordFunc,
}
