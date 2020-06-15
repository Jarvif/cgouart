/*
**************************************************************************
*	 File Name: myuart.h
*	 Author: jarvif
*	 Mail: liuzw@sunmnet.com 
*	 Created Time: Thu 11 Jun 2020 00:04:05 PDT
*
*	 COPYRIGHT(c) Sunmnet Technology Co. Ltd
**************************************************************************
*/
#ifndef MYUART_H
#define MYUART_H

#include <termios.h>            /* tcgetattr, tcsetattr */
#include <stdio.h>              /* perror, printf, puts, fprintf, fputs */
#include <unistd.h>             /* read, write, close */
#include <fcntl.h>              /* open */
#include <sys/signal.h>
#include <sys/types.h>
#include <string.h>             /* bzero, memcpy */
#include <limits.h>             /* CHAR_MAX */

typedef int     INT32;
typedef short   INT16;
typedef char    INT8;
typedef unsigned int UNIT32;
typedef unsigned short UINT16;
typedef unsigned char UINT8;

/* myuart.c */
INT32 OpenComPort (char *devicename, INT32 baudrate, INT32 databit,
										char *stopbit, char parity);
INT32 CloseComPort (void);

INT32 ReadComPort (char *data, INT32 datalength);
INT32 WriteComPort (char * data, INT32 datalength);

/**
 * export serial fd to other program to perform
 * directly read, write to serial.
 * 
 * @return serial's file description 
 */
int getPortFd();

#endif
