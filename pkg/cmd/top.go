package cmd

import (
	"github.com/spf13/cobra"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

var topCmd = &cobra.Command{
	Use:   "top",
}

var topQueriesCmd = &cobra.Command{
	Use:   "queries",
	Run: func(cmd *cobra.Command, args []string) {
		var connString = "host=" + hostname + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"
		db, err := sql.Open("postgres", connString)

		defer db.Close()

		if err != nil {
			panic(err)
		}

		slowQueries, err := db.Query("SELECT * FROM pg_stat_statements JOIN pg_roles r ON r.oid = userid WHERE calls > 100 ORDER BY mean_time DESC")

		if err != nil {
			panic(err)
		}

		for slowQueries.Next() {
			var name string
			if err := slowQueries.Scan(&name); err != nil {
				panic(err)
			}
			fmt.Printf("this is something: %s\n", name)
		}

	},
}