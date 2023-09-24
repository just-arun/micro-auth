/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/just-arun/micro-auth/util"
	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, err := cmd.Flags().GetString("privateKey")
		if err != nil {
			panic("private key")
		}
		str, err := cmd.Flags().GetString("value")
		if err != nil {
			panic("value")
		}

		value, err := util.Rsa().Decrypt(key, str)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("VALUE: \n-----------------\n\n%v\n\n", value)
		fmt.Println("decrypt called")
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	decryptCmd.Flags().StringP("privateKey", "k", "", "public key")
	decryptCmd.Flags().StringP("value", "v", "", "string to encrypt")
}
