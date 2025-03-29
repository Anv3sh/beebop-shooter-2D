package internals

import "github.com/hajimehoshi/ebiten/v2"

type Laser struct{
	Sprite *ebiten.Image
	XCoordinate float64
	YCoordinate float64
}


func (l *Laser) Move(speed float64){
	l.YCoordinate -= speed
}