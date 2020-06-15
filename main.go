package main

import (
	"os"
	"time"
	"fmt"
	"cgouart/uart"
	"unsafe"
)

func ProcessMyuart(readbuf [255]byte,length int){
		for i := 0; i < length; i++ { //looping from 0 to the length of the array
			fmt.Printf("%d ",readbuf[i])
		}
		fmt.Printf("\n")
}


func main() {
	var length int 
	var ret int
	readbuf := [255]byte{0}
	
	ret = uart.OpenMyuart("/dev/ttymxc1",9600, 8, "1", 'N');
	if ret < 0{
		fmt.Println("Open  myuart error\n")
		os.Exit(1)
	}

	length = uart.WriteMyuart("123456",6)
	if length <= 0 {
		fmt.Println("Write myuart error\n")
	}
	fmt.Printf("Write size %d data\n",length)

	for{
		length = uart.ReadMyuart((*byte)(unsafe.Pointer(&readbuf)),255)
		if length < 0 {
			fmt.Println("Read myuart error\n")
		}else if length == 0{
			continue
		}else{
			fmt.Printf("Read size %d data\n",length)
	/*	
			for i := 0; i < length; i++ { //looping from 0 to the length of the array
				fmt.Printf("%d ",readbuf[i])
			}
			fmt.Printf("\n")
	*/	
		go ProcessMyuart(readbuf,length)
		}
	}

	time.Sleep(time.Duration(2)*time.Second)
	uart.CloseMyuart()
}

