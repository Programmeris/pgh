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
		var connString = "host=" + hostname + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"
		db, err := sql.Open("postgres", connString)

		defer db.Close()

		if err != nil {
			panic(err)
		}

		tables, err := db.Query("SELECT tablename FROM pg_catalog.pg_tables")

		if err != nil {
			panic(err)
		}

		util.PrintRows(tables, "Tables")

	},
}

var seqCmd = &cobra.Command{
	Use:   "sequences",
	Run: func(cmd *cobra.Command, args []string) {
		var connString = "host=" + hostname + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"
		db, err := sql.Open("postgres", connString)

		defer db.Close()

		if err != nil {
			panic(err)
		}

		sequences, err := db.Query("SELECT sequencename FROM pg_catalog.pg_sequences")

		if err != nil {
			panic(err)
		}

		util.PrintRows(sequences, "Sequences")

	},
}

var indexCmd = &cobra.Command{
	Use:   "indexes",
	Run: func(cmd *cobra.Command, args []string) {
		var connString = "host=" + hostname + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"
		db, err := sql.Open("postgres", connString)

		defer db.Close()

		if err != nil {
			panic(err)
		}

		indexes, err := db.Query("SELECT indexname FROM pg_catalog.pg_indexes")

		if err != nil {
			panic(err)
		}

		util.PrintRows(indexes, "Indexes")

	},
}

var viewsCmd = &cobra.Command{
	Use:   "views",
	Run: func(cmd *cobra.Command, args []string) {
		var connString = "host=" + hostname + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"
		db, err := sql.Open("postgres", connString)

		defer db.Close()

		if err != nil {
			panic(err)
		}

		views, err := db.Query("SELECT viewname FROM pg_catalog.pg_views")

		if err != nil {
			panic(err)
		}

		util.PrintRows(views, "Views")

	},
}

var matViewsCmd = &cobra.Command{
	Use:   "matviews",
	Run: func(cmd *cobra.Command, args []string) {
		var connString = "host=" + hostname + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"
		db, err := sql.Open("postgres", connString)

		defer db.Close()

		if err != nil {
			panic(err)
		}

		matviews, err := db.Query("SELECT matviewname FROM pg_catalog.pg_matviews")

		if err != nil {
			panic(err)
		}

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