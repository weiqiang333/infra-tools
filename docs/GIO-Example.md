# GIO-Example

```
    适用于 GrowingIO EC2 卷扩容案例的参考
    AWS API 认证，请登录 jenkins0/infra0 等服务器 apps users 进行操作。
        API 的认证处于严格的监管，除特殊服务器之外的个人及服务器不得拥有认证配置。
    服务器磁盘的调整大小，请位于 jumpbox 机器进行操作。
```

### 扩容卷
- 案例一
```bash
# 对 hosts_list 中服务器的磁盘扩容，并过滤磁盘大小
cat > hosts_list << EOF
cnhd243
cnhd248
cnhd250
cnhd221
cnhd320
cnhd330
EOF

for host in $(cat hosts_list)
do
    infra-tools awscloud ec2Volumes -t Name:$host -s 500,700 -m --modify-size-add 100
done
```

- 案例二
```bash
infra-tools awscloud ec2Volumes -t App:Hadoop -t Env:Production -s 500,700 -m --modify-size-add 100
```

### 扩容磁盘
- 案例
```
    bash internal/awscloud/ec2/volumes/resize_disk.sh
```