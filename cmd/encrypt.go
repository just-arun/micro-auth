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

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, err := cmd.Flags().GetString("publicKey")
		if err != nil {
			panic("public error")
		}
		str, err := cmd.Flags().GetString("value")
		if err != nil {
			panic("value error")
		}

		pubk, err := util.Rsa2().PublicKeyFrom64(key)

		if err != nil {
			log.Fatalf("ERROR: (0) %v", err)
		}

		encStr, err := util.Rsa2().PublicEncryptReturnsBase64String(pubk, []byte(str))

		if err != nil {
			log.Fatalf("ERROR: (1) %v", err)
		}

		fmt.Println("\n\nENCRYPTED STRING: \n-------------------------\n\n", encStr)

	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	encryptCmd.Flags().StringP("publicKey", "k", "", "public key in base64 formate")
	encryptCmd.Flags().StringP("value", "v", "", "string to encrypt")
}
