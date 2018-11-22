#!/bin/sh
#export LD_LIBRARY_PATH=/opt/glibc2.14/lib:$LD_LIBRARY_PATH
#go build -a -o file_pool_ser main.go
kill -9 $(pidof /var/gowww/src/file_pool_service/file_pool_ser)
nohup /var/gowww/src/file_pool_service/file_pool_ser -c /var/gowww/src/file_pool_service/conf/conf.yaml > /var/gowww/src/file_pool_service/sys.log  2>&1 &
