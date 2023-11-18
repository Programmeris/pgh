package cmd

import (
	"github.com/spf13/cobra"
	_ "github.com/lib/pq"
	"fmt"
)

var ( 
	hostname string
	port string
	username string
	password string
	database string
)

var pghCmd = &cobra.Command{
	Use:   "pgh",
	Short: "Simple utility for interactive with PostgreSQL",
	Long:  "Simple utility for interactive with PostgreSQL",
}

func init(){

	pghCmd.PersistentFlags().StringVar(&hostname, "a", "localhost", "")
	pghCmd.PersistentFlags().StringVar(&port, "p", "5432", "")
	pghCmd.PersistentFlags().StringVar(&username, "U", "", "")
	pghCmd.PersistentFlags().StringVar(&password, "P", "", "")
	pghCmd.PersistentFlags().StringVar(&database, "d", "", "")

	topCmd.AddCommand(topQueriesCmd)
	getCmd.AddCommand(matViewsCmd)
	getCmd.AddCommand(tablesCmd)
	getCmd.AddCommand(indexCmd)
	getCmd.AddCommand(viewsCmd)
	getCmd.AddCommand(locksCmd)
	getCmd.AddCommand(seqCmd)
	pghCmd.AddCommand(getCmd)
	pghCmd.AddCommand(topCmd)

}

func Execute() {
	if err := pghCmd.Execute(); err != nil {
	  fmt.Println(err)
	}
  }