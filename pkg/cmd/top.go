package cmd

import (
	"database/sql"
	"pgh/pkg/util"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var topCmd = &cobra.Command{
	Use:   "top",
}

var topQueriesCmd = &cobra.Command{
	Use:   "queries",
	Run: func(cmd *cobra.Command, args []string) {
		var connString = "host=" + hostname + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"
		db, err := sql.Open("postgres", connString)

		if err != nil {
			panic(err)
		}

		defer db.Close()

		slowQueries, err := db.Query("SELECT * FROM pg_stat_statements JOIN pg_roles r ON r.oid = userid WHERE calls > 100 ORDER BY mean_time DESC")

		if err != nil {
			panic(err)
		}

		util.PrintRows(slowQueries, "Slow Queries")

	},
}
