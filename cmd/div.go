/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// divCmd represents the div command
var divCmd = &cobra.Command{
	Use:   "div",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		num1, _ := strconv.ParseUint(args[0], 10, 32)
		num2, _ := strconv.ParseUint(args[1], 10, 32)

		//for when the denominator is zero answer is undefined
		if num2 == 0 {
			fmt.Println("Answer is Undefined")

		} else {
			ans := float64(num1) / float64(num2)
			fmt.Println(ans)
		}

	},
}

func init() {
	rootCmd.AddCommand(divCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// divCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// divCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}