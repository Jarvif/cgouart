#########################################################################
# File Name: build.sh
# Author:jarvif 
# mail:liuzw@sunmnet.com 
# Created Time: Fri 12 Jun 2020 02:06:33 PDT
#########################################################################
#!/bin/bash
CGO_ENABLED=1 GOOS=linux GOARCH=arm CC=arm-none-linux-gnueabi-gcc go build
