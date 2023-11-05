package main

import "github.com/spf13/cobra"

func main() {
	initCommandFlags()
	pghCmd.Execute()
}

func initCommandFlags()  {
	pghCmd.PersistentFlags().StringP("hostname", "a", "localhost", "")
	pghCmd.PersistentFlags().StringP("port", "p", "5432", "")
	pghCmd.PersistentFlags().StringP("username", "U", "", "")
	pghCmd.PersistentFlags().StringP("password", "P", "", "")

	getCmd.AddCommand(tablesCmd)
	pghCmd.AddCommand(getCmd)
}

var pghCmd = &cobra.Command{
	Use:   "pgh",
	Short: "Simple utility for interactive with PostgreSQL",
	Long:  "Simple utility for interactive with PostgreSQL",
}

var getCmd = &cobra.Command{
	Use:   "get",
}

var tablesCmd = &cobra.Command{
	Use:   "tables",
}
