package internals

import "github.com/hajimehoshi/ebiten/v2"

type Laser struct{
	Sprite *ebiten.Image
	XCoordinate float64
	YCoordinate float64
	Hit bool
	HitTimer int
}


func (l *Laser) Update(speed float64) bool{
	if l.Hit && l.HitTimer<=0{
		return true
	}
	l.YCoordinate -= speed
	l.HitTimer--
	return false
}

func (l *Laser) isLaserOutOfBounds(windowW float64, windowH float64) bool {
	if l.YCoordinate < 0 {
		return true
	}

	if l.YCoordinate>windowH - float64(l.Sprite.Bounds().Dy()){
		return true
	}

	if l.XCoordinate < 0 {
		return true
	}
	if l.XCoordinate > windowW- float64(l.Sprite.Bounds().Dx()) {
		return true
	}
	return false
}

func (l *Laser) drawLaser(screen *ebiten.Image){
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5,0.5)
	op.GeoM.Translate(l.XCoordinate, l.YCoordinate)
	screen.DrawImage(l.Sprite,op)
}