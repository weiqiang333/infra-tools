# infra-tools

System operation and maintenance basic tools

Currently supports the following features:
```
    fileserver: It supports breakpoint continuation and segment Download.
    awscloud: Management of AWS Cloud Platform Resources
        Authentication uses default mode [`credentials files`](https://docs.aws.amazon.com/zh_cn/cli/latest/userguide/cli-configure-files.html)
    alicloud: Management of Alibaba Cloud Platform Resources
```


# use

- infra-tools -h
```
Usage:
  infra-tools [command]

Available Commands:
  alicloud    infra Tools for alibaba cloud
  awscloud    infra Tools for aws cloud
  fileserver  infra Tools for fileserver
  help        Help about any command
  version     Current version

Flags:
      --config string   config file (default is $HOME/.infra-tools.yaml)
  -h, --help            help for infra-tools
```


### Example

##### fileserver
- Provide /data/download/ and /apps/svr/ directory downloads
```bash
infra-tools fileserver -d /data/download/ -d /apps/svr/
```

##### aws cloud
- View volume information
```bash
infra-tools awscloud ec2Volumes -t App:Consul -s 100,300
```

- Modify volume size
```bash
infra-tools awscloud ec2Volumes -t App:Kubernetes -t Env:Production -s 100,300 -m --modify-size-add 50    
    # Modify filter criteria Env=Production and App=Kubernetes Volume size range is 100,300 Increase size 50G
```
