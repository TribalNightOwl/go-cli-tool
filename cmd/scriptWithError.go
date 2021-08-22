/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/spf13/cobra"
)

// scriptWithErrorCmd represents the scriptWithError command
var scriptWithErrorCmd = &cobra.Command{
	Use:   "scriptWithError",
	Short: "Execute script with an exit error",
	Run: func(cmd *cobra.Command, args []string) {
		command := exec.Command("./exit-error.sh", strconv.Itoa(desiredExitCode))
		if err := command.Run(); err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				exitCode := exitError.ExitCode()
				fmt.Println(exitCode)
			}
		}
	},
}

var desiredExitCode int

func init() {
	rootCmd.AddCommand(scriptWithErrorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scriptWithErrorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scriptWithErrorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	scriptWithErrorCmd.Flags().IntVarP(&desiredExitCode, "exit-code", "e", 18, "desired exit code for the script")

}
