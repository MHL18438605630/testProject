package main

/*
// C 标志io头文件，你也可以使用里面提供的函数
#include <stdlib.h>
#include <stdio.h>
#include <stdin.h>


void hello(string msg){
	printf("hello %s\n",msg);
}
*/
import "C"

func main() {

	C.hello("world")

}
