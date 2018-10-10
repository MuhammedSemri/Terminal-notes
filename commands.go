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
	Args: cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		nt.Title = args[0]
		if args[1] == "content" || args[1] == "c" && args[2] != ""{
			nt.Cont = args[2]
			fmt.Println("Saved!!!")		
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
		nt.Save()
		log.Println("saved!!")
	},
}

var DeleteIDCmd = &cobra.Command{
	Use: "idelete",
	Aliases: []string{
		"id",
	},
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		stmt, err := DB.Prepare("delete from Notes where id=?")
		if err != nil {
			log.Println(err)
			return
		}
		defer stmt.Close()
		_,err = stmt.Exec(id)
		if err != nil{
			log.Println(err)
			return
		}	
		
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

var DelteAllCmd = &cobra.Command{
	Use: "ad",
	Run: func(cmd *cobra.Command, args []string) {
		DeleteAll()
		fmt.Println("Deleted all your notes :)")
	},
}