# infra-tools

System operation and maintenance basic tools

Currently supports the following features:
```
    fileserver: It supports breakpoint continuation and segment Download.
    ec2Volumes: It's a volume management tool
        Authentication uses default mode [`credentials files`](https://docs.aws.amazon.com/zh_cn/cli/latest/userguide/cli-configure-files.html)
```


# use

- infra-tools -h
```
    infra Tools for Systems.
        fileserver: It supports breakpoint continuation and segment Download.
        ec2Volumes: It's a volume management tool
    
    Usage:
      infra-tools [command] [flags]
    
    Available Commands:
      ec2Volumes  infra Tools for ec2Volumes
      fileserver  infra Tools for fileserver
      help        Help about any command
      version     Current version
    
    Flags:
          --config string   config file (default is $HOME/.infra-files-server.yaml)
      -h, --help            help for infra-tools
```


### Example
- Provide /data/download/ and /apps/svr/ directory downloads
```bash
infra-tools fileserver -d /data/download/ -d /apps/svr/
```

- View volume information
```bash
infra-tools ec2Volumes -t App:Consul -s 100,300
```

- Modify volume size
```bash
infra-tools ec2Volumes -t App:Kubernetes -t Env:Production -s 100,300 -m --modify-size-add 50    
    # Modify filter criteria Env=Production and App=Kubernetes Volume size range is 100,300 Increase size 50G
```
