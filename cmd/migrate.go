/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	migration "github.com/just-arun/micro-auth/migration"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate called")
		env, err := cmd.Flags().GetString("environment")
		if err != nil {
			panic("env is not provided")
		}
		ctx, err := cmd.Flags().GetString("environment-context")
		if err != nil {
			panic("environment-context is not provided")
		}

		migration.Run(ctx, env)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")


	migrateCmd.Flags().StringP("environment", "e", "dev", "env refers to the file which we call on run time ie if we pass env as `dev` then it will look for `.env.dev.yml` in case of `prod` it is `.env.prod.yml`")
	migrateCmd.Flags().StringP("environment-context", "c", ".", "environment context to find app configuration path")
}
