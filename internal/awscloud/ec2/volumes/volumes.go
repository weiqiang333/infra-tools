package volumes

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/weiqiang333/infra-tools/model"
)

var (
	sess = session.New()
	svc = ec2.New(sess, &aws.Config{
		Region:	aws.String(endpoints.CnNorth1RegionID),
	})
	volumes = []model.Volumes{}
)

// ReadVolume 过滤查询卷，读取卷信息
func ReadVolume(tags map[string]string, size []int64)  {
	filters := []*ec2.Filter{}
	filters = append(filters, &ec2.Filter{
		Name: aws.String("status"),
		Values: []*string{
			aws.String("in-use"),
		},
	})
	for k, v := range tags {
		filters = append(filters, &ec2.Filter{
			Name: aws.String(fmt.Sprintf("tag:%s", k)),
			Values: []*string{
				aws.String(v),
			},
		})
	}
	input := &ec2.DescribeVolumesInput{
		Filters: filters,
	}
	result, err := svc.DescribeVolumes(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, volume := range result.Volumes{
		var name = ""
		if len(size) == 1 {
			if *volume.Size != size[0] {
				continue
			}
		} else if size[0] > *volume.Size || *volume.Size > size[len(size)-1] {
			continue
		}

		for _, tag := range volume.Tags {
			if *tag.Key == "Name" {
				name = *tag.Value
			}
		}
		v := model.Volumes{}
		v.Name = name
		v.Size = *volume.Size
		v.VolumesID = *volume.VolumeId
		v.VolumeType = *volume.VolumeType
		volumes = append(volumes, v)
		fmt.Println(v)
	}
	return
}

// ModifyVolume 修改卷大小
func ModifyVolume(modifySize int64, modifySizeAdd int64)  {
	if modifySize == 0 {
		fmt.Println("将使用 modify-size-add 方法增加卷大小: ", modifySizeAdd)
	} else {
		fmt.Println("将使用 modify-size 方法修改卷大小: ", modifySize)
	}
	var input = &ec2.ModifyVolumeInput{}
	for i, volume := range volumes {
		fmt.Print(i, "\t", volume, "\n\t")
		if modifySize == 0 {
			input = &ec2.ModifyVolumeInput{
				VolumeId: aws.String(volume.VolumesID),
				Size: aws.Int64(volume.Size + modifySizeAdd),
			}
		} else {
			input = &ec2.ModifyVolumeInput{
				VolumeId: aws.String(volume.VolumesID),
				Size: aws.Int64(modifySize),
			}
		}

		// 修改卷大小
		result, err := svc.ModifyVolume(input)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("修改前：", *result.VolumeModification.OriginalSize, "\t修改后：", *result.VolumeModification.TargetSize)
	}
}

