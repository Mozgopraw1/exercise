package main

import "fmt"

func main(){
	var x int
	_,err := fmt.Scan(x)
	if err != nil {
		fmt.Println(err)
	}
		
}
