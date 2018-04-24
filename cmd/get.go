// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/Vico1993/tbot/tbot"
	"github.com/spf13/cobra"
)

var Specific string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [OPTIONS]",
	Short: "Get will you return all information from twitter",
	Run: func(cmd *cobra.Command, args []string) {
		if Specific == "" {
			tbot.GetTopTrends(49.246292, -123.116226)
		} else {
			tbot.GetSpecificStream(Specific)
		}

	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	getCmd.Flags().StringVarP(&Specific, "specific", "s", "", "This flags will you return a stream of a specific keyword")
}
