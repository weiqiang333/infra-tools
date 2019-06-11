// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"github.com/spf13/cobra"

	"infra-tools/web"
	"infra-tools/model"
)

// fileserverCmd represents the fileserver command
var fileserverCmd = &cobra.Command{
	Use:   "fileserver",
	Short: "infra Tools for fileserver",
	Long: `infra Tools for fileserver:
	It supports breakpoint continuation and segment Download.`,
	Run: func(cmd *cobra.Command, args []string) {
		web.Web()
	},
}

func init() {
	rootCmd.AddCommand(fileserverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fileserverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fileserverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	fileserverCmd.Flags().StringSliceVarP(&model.Dir, "dir","d", []string{"/tmp/"}, "Absolute path: /data/,/apps/svr/")
}
