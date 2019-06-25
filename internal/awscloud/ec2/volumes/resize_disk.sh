#!/usr/bin/env bash

: '''
    AWS 云服务 EC2，支持动态扩容。
    调用 EC2 API 扩容之后，依然需要在此对服务器文件系统进行加载扩容。
'''

hosts_file="hosts_list"

for host in $(cat $hosts_file)
do
    echo $host
    ssh $host "
        for dev in \$(df -h | grep appsdata | awk '{print \$1}')
        do
            sudo resize2fs \$dev
        done
    "
done
