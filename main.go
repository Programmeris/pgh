package pgh

import "github.com/spf13/cobra"

func main() {
	
}

func initCommandFlags()  {
	
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
	Use:   "pgh",
}