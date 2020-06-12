package main

import "fmt"

var exp = 0
var level = 1

func UpLevel(e int){
	exp += e
	level = (exp / 10) + 1
}

func main(){
	UpLevel(10)
	fmt.Printf("Level:%d, exp:%d",level,exp)
}