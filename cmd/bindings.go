/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"

	"github.com/jordangarrison/gcp-iam-compare/lib"
	"github.com/spf13/cobra"
	"google.golang.org/api/cloudresourcemanager/v1"
)

// bindingsCmd represents the bindings command
var bindingsCmd = &cobra.Command{
	Use:   "bindings",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		rmService, err := lib.NewRMService(ctx)
		if err != nil {
			panic(err)
		}
		req := cloudresourcemanager.GetIamPolicyRequest{}
		getCall, err := rmService.Projects.GetIamPolicy("flosports-174016", &req).Do()
		if err != nil {
			panic(err)
		}

		var bindings []iamBinding
		for _, p := range getCall.Bindings {
			bindings = append(bindings, iamBinding{
				Members: p.Members,
				Role:    p.Role,
			})
		}
		fmt.Printf("%+v", bindings)
	},
}

type iamBinding struct {
	Members []string
	Role    string
}

func init() {
	rootCmd.AddCommand(bindingsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bindingsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bindingsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
