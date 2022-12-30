package cmd

import (
	"fmt"

	"gin-todo-prc/migrations"
	"gin-todo-prc/util"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "database migrations tool",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new empty migration file",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println("Unable to read flag name", err.Error())
			return
		}

		if err := migrations.Create(name); err != nil {
			fmt.Println("Unable to create migration", err.Error())
			return
		}
	},
}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "run up migrations",
	Run: func(cmd *cobra.Command, args []string) {

		step, err := cmd.Flags().GetInt("step")
		if err != nil {
			fmt.Println("Unable to read flag `step`")
			return
		}

		db := util.SetupDb()

		migrator, err := migrations.Init(db)
		if err != nil {
			fmt.Println("Unable to fetch migrator")
			return
		}

		err = migrator.Up(step)
		if err != nil {
			panic(err)
			// fmt.Println("Unable to run `up` migrations")
			// return
		}

	},
}

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "run down migrations",
	Run: func(cmd *cobra.Command, args []string) {

		step, err := cmd.Flags().GetInt("step")
		if err != nil {
			fmt.Println("Unable to read flag `step`")
			return
		}

		db := util.SetupDb()

		migrator, err := migrations.Init(db)
		if err != nil {
			fmt.Println("Unable to fetch migrator")
			panic(err)
		}

		err = migrator.Down(step)
		if err != nil {
			fmt.Println("Unable to run `down` migrations")
			return
		}
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "display status of each migrations",
	Run: func(cmd *cobra.Command, args []string) {
		db := util.SetupDb()
		fmt.Println("util.SetupDb")

		migrator, err := migrations.Init(db)
		if err != nil {
			fmt.Println("Unable to fetch migrator")
			return
		}

		if err := migrator.MigrationStatus(); err != nil {
			fmt.Println("Unable to fetch migration status")
			return
		}

		return
	},
}

func init() {

	createCmd.Flags().StringP("name", "n", "", "Name for the migration")    //add "--name" or "-n" flag to the command
	upCmd.Flags().IntP("step", "s", 0, "Number of migrations to execute")   //add "--step" or "-s" flag to the command
	downCmd.Flags().IntP("step", "s", 0, "Number of migrations to execute") //add "--step" or "-s" flag to the command
	migrateCmd.AddCommand(upCmd, downCmd, createCmd, statusCmd)             // adding all commands to the "migrate" command
	rootCmd.AddCommand(migrateCmd)                                          //adding 'migrate' command to the root

}
