/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

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
		fmt.Println("decrypt called")
		key, err := cmd.Flags().GetString("privateKey")
		if err != nil {
			panic("private key")
		}
		str, err := cmd.Flags().GetString("value")
		if err != nil {
			panic("value")
		}

		priK, err := util.Rsa2().PrivateKeyFrom64(key)

		if err != nil {
			log.Fatalf("ERROR: (3) %v", err)
		}

		decStr, err := util.Rsa2().PrivateDecryptWithBase64String(priK, str)

		if err != nil {
			log.Fatalf("ERROR: (4) %v", err)
		}

		fmt.Println("\n\nDECODED STRING: \n-------------------------\n\n", string(decStr))
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
