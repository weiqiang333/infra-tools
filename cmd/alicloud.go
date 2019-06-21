/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/viper"
	"github.com/spf13/cobra"

	"github.com/weiqiang333/infra-tools/internal/alicloud/ecs"
)

// alicloudCmd represents the alicloud command
var alicloudCmd = &cobra.Command{
	Use:   "alicloud",
	Short: "infra Tools for alibaba cloud",
	Long: `infra Tools for alicloud:
	alicloud: Management of Alibaba Cloud Platform Resources`,
}

var ecsSecurityGroup = &cobra.Command{
	Use:   "ecsSecurityGroup",
	Short: "infra Tools for alicloud 's ecsSecurityGroup",
	Long: `infra Tools for alicloud:
	ecsSecurityGroup`,
	Run: func(cmd *cobra.Command, args []string) {
		ecs.LookSecurityGroups()
		addRules, _ := cmd.Flags().GetBool("AddRules")
		if addRules == true {
			fmt.Println("Adding security group rules")
			ecs.AddSecurityGroupRules()
		}
	},
}

func init() {
	rootCmd.AddCommand(alicloudCmd)
	alicloudCmd.AddCommand(ecsSecurityGroup)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// alicloudCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// alicloudCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	ecsSecurityGroup.Flags().String("GroupName","","ecs SecurityGroupName")
	ecsSecurityGroup.Flags().String("GroupId", "", "ecs SecurityGroupId")
	ecsSecurityGroup.Flags().BoolP("AddRules", "a",false, "Adding Security Group Rules")
	ecsSecurityGroup.Flags().String("IpProtocol", "TCP", "传输层协议,取值范围：\n TCP,UDP,ICMP,GRE,ALL")
	ecsSecurityGroup.Flags().String("PortRange", "9999/9999", "开放的传输层协议相关的端口范围")
	ecsSecurityGroup.Flags().String("SourceCidrIp", "", "源端 IPv4 CIDR 地址段")
	ecsSecurityGroup.Flags().String("Priority", "100", "安全组规则优先级。取值范围：1~100, 1最高")
	ecsSecurityGroup.Flags().String("Description", "Increased by sdk", "安全组规则的描述信息")
	viper.BindPFlags(ecsSecurityGroup.Flags())

}
