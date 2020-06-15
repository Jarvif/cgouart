一、项目概述

	本项目是golang嵌入式开发方法，主要是为嵌入式开发平台打造，基于cgo实现了golang对于/dev目录下的uart设备控制，读写等操作。cgo相关部分已封装成golang包，在golang业务代码中使用golang调用即可。

二、编译方法与执行

1、编译方法

x86平台:go build

arm平台：CGO_ENABLED=1 GOOS=linux GOARCH=arm CC=arm-none-linux-gnueabi-gcc go build (注意：CC是你的交叉编译工具名称)

2、执行程序方法

两种平台都是: 

./cgouart

三、本项目已在x86、IMX6Q平台上测试通过

四、移植说明

	在移植过程中， 需要更改main.go执行程序的相应串口设备名称,即可移植使用，例如6Q的: /dev/ttymxc1或者x86的:ttyUSB0。这个需要自己根据自己平台串口的名称进行更改移植

五、意义

	对于做嵌入式开发来说，golang的网络支持是非常好用的，但是往往嵌入式平台下，经常需要用到一些外设，比如gpio，uart、i2c，spi等等一些/dev目录下的一些设备。所以，使用cgo封装成go包，对于要求需要网络，同时也需要基础硬件控制的平台来说，将会非常方便。

六、附录

	其实，就算不用cgo也能使用golang里的syscall做设备控制，只是基于很多的设备控制的例子都是C语言的，所以，这样会相对比较省事。

七、如有其它疑问，请e-mail:474418595@qq.com
