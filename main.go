package main

import (
	"fmt"
	"github.com/Anv3sh/test_game/internals"
	"github.com/hajimehoshi/ebiten/v2"
)


func main(){
	fmt.Println("New game!")
	g := &internals.Game{
		Player: &internals.Player{
			Sprite: internals.MustLoadImage(internals.RAPTOR),
			XCoordinate: 250,
			YCoordinate: 200,
		},
		WindowW: 640, 
		WindowH: 480, 
		Space: internals.MustLoadImage(internals.SPACE_BACKGROUND_PURPLE),
	}
	ebiten.SetWindowSize(int(g.WindowW),int(g.WindowH))
	ebiten.SetWindowTitle("Space Shooter 2D")
	err := ebiten.RunGame(g)
	if err!=nil{
		panic(err)
	}
}