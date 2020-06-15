package uart
/*
#include "myuart.h"
*/
import "C"

import (
	"unsafe"
)

func OpenMyuart(devicename string,baudrate int,databit int,stopbit string,parity byte)(int){
	result := C.OpenComPort(C.CString(devicename), C.int(baudrate), C.int(databit), C.CString(stopbit), C.char(parity));
	return int(result)
}


func WriteMyuart(writebuf string,size int)(int){
	result := C.WriteComPort(C.CString(writebuf),C.int(size))
	return int(result)
}

func ReadMyuart(readbuf *byte ,size int)(int){
	result := C.ReadComPort((*C.char)(unsafe.Pointer(readbuf)),C.int(size))
	return int(result)
}

func CloseMyuart()(int){
	result := C.CloseComPort()
	return int(result)
}

