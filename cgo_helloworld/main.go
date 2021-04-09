package main

//#include<stdio.h>
//void SayHello(_GoString_ s);
import "C"
import "fmt"

func main(){
	C.SayHello("Hello world")
}

//export SayHello
func SayHello(s string){
	fmt.Printf("Test: %s\n",s)
}
