package main

import (
	"fmt"
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
		if args[1] == "content" || args[1] == "-c" && args[2] != ""{
			nt.Cont = args[2]		
		}else {
			fmt.Println("missing content argument (-c)")
		}
		nt.Save()

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