/*
Copyright © 2022 Jordan Garrison <dev@jordangarrison.dev>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/jordangarrison/gcp-iam-compare/lib"
	"github.com/spf13/cobra"
)

// compareRolesCmd represents the compareRoles command
var compareRolesCmd = &cobra.Command{
	Use:   "compare",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			panic("Not enough args")
		}
		role0 := args[0]
		role1 := args[1]
		fmt.Println("Comparing roles", role0, role1)
		ctx := context.Background()
		svc, err := lib.NewService(ctx)
		if err != nil {
			panic(err)
		}
		role0Resp, err := svc.Roles.Get(role0).Context(ctx).Do()
		if err != nil {
			panic(err)
		}
		role0Permissions := role0Resp.IncludedPermissions
		role1Resp, err := svc.Roles.Get(role1).Context(ctx).Do()
		if err != nil {
			panic(err)
		}
		role1Permissions := role1Resp.IncludedPermissions
		allPermissions := append(role0Permissions, role1Permissions...)
		mapCounts := make(map[string]int)
		for _, v := range allPermissions {
			if _, ok := mapCounts[v]; ok {
				mapCounts[v]++
			} else {
				mapCounts[v] = 1
			}
		}
		for k, v := range mapCounts {
			if v > 1 {
				fmt.Println(k)
			}
		}
	},
}

func init() {
	rolesCmd.AddCommand(compareRolesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compareRolesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compareRolesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}