package main

import (
	"fmt"
	"container/list"
)

func main(){

	var myList list.List

	myList.PushBack("apple")
	myList.PushBack("papaya")
	myList.PushBack("orange")
	myList.PushBack("grapes")

	for ele:=myList.Front();ele!=nil;ele=ele.Next(){
		fmt.Println(ele.Value)
	}


}