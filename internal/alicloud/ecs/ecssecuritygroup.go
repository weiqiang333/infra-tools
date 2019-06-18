package ecs

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

var (
	lookSG = map[string]string{}
)

// LookSecurityGroups 查看安全组信息
func LookSecurityGroups() {
	client, err := ecs.NewClientWithAccessKey(viper.GetString("alicloud.region-id"), viper.GetString("alicloud.access-key"), viper.GetString("alicloud.secret-key"))
	if err != nil {
		fmt.Println("ecs client err: ", err)
	}
	request := ecs.CreateDescribeSecurityGroupsRequest()
	request.SecurityGroupName = viper.GetString("groupname")
	request.SecurityGroupId = viper.GetString("groupid")
	response, err := client.DescribeSecurityGroups(request)
	if err != nil {
		fmt.Println("ecs response err: ", err.Error())
		return
	}
	for i, sg := range response.SecurityGroups.SecurityGroup {
		fmt.Println(i, sg.SecurityGroupName, sg.SecurityGroupId, sg.Description)
		lookSG[sg.SecurityGroupId] = sg.SecurityGroupName
	}
}

// AddSecurityGroupRules 修改安全组规则，增加入口规则策略
func AddSecurityGroupRules()  {
	client, err := ecs.NewClientWithAccessKey(viper.GetString("alicloud.region-id"), viper.GetString("alicloud.access-key"), viper.GetString("alicloud.secret-key"))
	if err != nil {
		fmt.Println("ecs client err: ", err)
	}
	for groupId, groupName := range lookSG {
		request := ecs.CreateAuthorizeSecurityGroupRequest()
		request.SecurityGroupId = groupId
		request.IpProtocol = viper.GetString("ipprotocol")
		request.PortRange = viper.GetString("portrange")
		request.SourceCidrIp = viper.GetString("sourcecidrip")
		request.Priority = viper.GetString("priority")
		request.Description = viper.GetString("description")
		resp, err := client.AuthorizeSecurityGroup(request)
		if err != nil {
			fmt.Println("ecs response err: ", err.Error())
			return
		}
		fmt.Println("Adding Rules in ", groupName, groupId, resp.GetHttpStatus(), resp.GetHttpContentString())
	}
}