package main

import (
	"fmt"
	"github.com/Anv3sh/bebop-shooter-2D/internals"
)


func main(){
	fmt.Println("New game!")
	err := internals.GameInitAndRun()
	if err!=nil{
		panic(err)
	}
}