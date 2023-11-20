package cmd

import (
	"database/sql"
	"pgh/pkg/util"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
}

var tablesCmd = &cobra.Command{
	Use:   "tables",
	Run: func(cmd *cobra.Command, args []string) {
		var connectionOptions = util.NewConnectionOptions(hostname, port, username, password, database)
		var tables = util.GetRowsFromTable(*connectionOptions, "tablename", "pg_catalog.pg_tables")

		util.PrintRows(tables, "Tables")

	},
}

var seqCmd = &cobra.Command{
	Use:   "sequences",
	Run: func(cmd *cobra.Command, args []string) {
		var connectionOptions = util.NewConnectionOptions(hostname, port, username, password, database)
		var sequences = util.GetRowsFromTable(*connectionOptions, "tablename", "pg_catalog.pg_sequences")

		util.PrintRows(sequences, "Sequences")

	},
}

var indexCmd = &cobra.Command{
	Use:   "indexes",
	Run: func(cmd *cobra.Command, args []string) {
		var connectionOptions = util.NewConnectionOptions(hostname, port, username, password, database)
		var indexes = util.GetRowsFromTable(*connectionOptions, "indexname", "pg_catalog.pg_indexes")

		util.PrintRows(indexes, "Indexes")

	},
}

var viewsCmd = &cobra.Command{
	Use:   "views",
	Run: func(cmd *cobra.Command, args []string) {
		var connectionOptions = util.NewConnectionOptions(hostname, port, username, password, database)
		var views = util.GetRowsFromTable(*connectionOptions, "viewname", "pg_catalog.pg_views")

		util.PrintRows(views, "Views")

	},
}

var matViewsCmd = &cobra.Command{
	Use:   "matviews",
	Run: func(cmd *cobra.Command, args []string) {
		var connectionOptions = util.NewConnectionOptions(hostname, port, username, password, database)
		var matviews = util.GetRowsFromTable(*connectionOptions, "matviewname", "pg_catalog.pg_matviews")

		util.PrintRows(matviews, "Materialized views")

	},
}

var locksCmd = &cobra.Command{
	Use:   "locks",
	Run: func(cmd *cobra.Command, args []string) {
		var connString = "host=" + hostname + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"
		db, err := sql.Open("postgres", connString)

		defer db.Close()

		if err != nil {
			panic(err)
		}

		locks, err := db.Query("SELECT locktype FROM pg_catalog.pg_locks")

		if err != nil {
			panic(err)
		}

		util.PrintRows(locks, "Locks")

	},
}
