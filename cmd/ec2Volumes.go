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
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"infra-tools/internal/ec2/volumes"
)

// ec2VolumesCmd represents the ec2Volumes command
var ec2VolumesCmd = &cobra.Command{
	Use:   "ec2Volumes",
	Short: "infra Tools for ec2Volumes",
	Long: `infra Tools for ec2Volumes:
	It's a volume management tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		tags, err := cmd.Flags().GetStringSlice("tag")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		tag_map := make(map[string]string)
		for _, tag := range tags {
			t := strings.Split(tag, ":")
			if len(t) != 2 {
				fmt.Println("tag's Flags incorrect")
				return
			}
			k, y := t[0], t[1]
			tag_map[k] = y
		}
		sizes, err := cmd.Flags().GetInt64Slice("size")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// 读取卷信息
		volumes.ReadVolume(tag_map, sizes)

		modify, err := cmd.Flags().GetBool("modify")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if modify == false {
			return
		}
		modifySize, err := cmd.Flags().GetInt64("modify-size")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		modifySizeAdd, err := cmd.Flags().GetInt64("modify-size-add")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if modifySize == 0 && modifySizeAdd == 0 {
			fmt.Println("Error: To modify a volume, you must have one of the parameters\n\t--modify-size\n\t--modify-size-add")
			os.Exit(79)
		}

		// 修改卷大小
		volumes.ModifyVolume(modifySize, modifySizeAdd)
	},
}

func init() {
	rootCmd.AddCommand(ec2VolumesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ec2VolumesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ec2VolumesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	ec2VolumesCmd.Flags().StringSliceP("tag","t", []string{"App:kubernetes"}, "Filter tag value")
	ec2VolumesCmd.Flags().Int64SliceP("size", "s", []int64{10,9999}, "Filter Volume size range\n" +
		"If the length of the parameter value is 1, Filter volumes of the same size")
	ec2VolumesCmd.Flags().BoolP("modify", "m", false, "Modify Volume (default false)")
	ec2VolumesCmd.Flags().Int64("modify-size",0, "Modify volume size")
	ec2VolumesCmd.Flags().Int64("modify-size-add", 0, "Modify the volume size incrementally\n" +
		"modify-size priority is higher")
}
