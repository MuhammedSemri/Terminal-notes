package main

import (
	"log"

	"github.com/spf13/cobra"
)

var nt Note

var RootCmd = &cobra.Command{
	Use: "notes",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var TitleCmd = &cobra.Command{
	Use: "title",
	Aliases: []string{
		"t",
	},
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(args[0])
		nt.Title = args[0]

	},
}

var ContCmd = &cobra.Command{
	Use: "content",
	Aliases: []string{
		"c",
	},
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		nt.Cont = args[0]
		log.Println(nt.Cont)
		if DB == nil {
			log.Println("null")
		}
		nt.Save()
		log.Println("saved!!")
	},
}

var GetAllCmd = &cobra.Command{
	Use: "all",
	Aliases: []string{
		"a",
	},
	Run: func(cmd *cobra.Command, args []string) {
		nt.GetAll()
	},
}