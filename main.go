package main

import (
	"fmt"
	"github.com/Anv3sh/test_game/internals"
)


func main(){
	fmt.Println("New game!")
	err := internals.GameInitAndRun()
	if err!=nil{
		panic(err)
	}
}