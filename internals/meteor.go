package internals

import (
	// "math"
	"math/rand"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)


type Meteor struct{
	Sprite *ebiten.Image
	XCoordinate float64
	YCoordinate float64
	Speed float64
	Type string
	HitTimer int
	Destroyed bool
}

func generateMeteor(windowW float64) *Meteor{
	return &Meteor{
		Sprite: MustLoadImage(METEOR_MED), 
		XCoordinate: float64(rand.Intn(int(windowW))), 
		YCoordinate: -10.0,
		Speed: 2,
	}
}

func (m *Meteor) updateMeteor() bool{
	if m.Destroyed && m.HitTimer <=0{
		return true
	}
	m.YCoordinate += m.Speed
	m.HitTimer--
	return false
}

func (m *Meteor) isMeteorOutOfBounds(windowW float64, windowH float64) bool{
	// if m.YCoordinate < 0 {
	// 	return true
	// }
	if m.YCoordinate>windowH - float64(m.Sprite.Bounds().Dy()){
		return true
	}

	if m.XCoordinate < 0 {
		return true
	}
	if m.XCoordinate > windowW- float64(m.Sprite.Bounds().Dx()) {
		return true
	}
	return false
}

func (m *Meteor) drawMeteor(screen *ebiten.Image){
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.XCoordinate,m.YCoordinate)
	// op.GeoM.Rotate(45.0 * math.Pi / 180.0)
	if m.HitTimer > 0 {
		op.ColorScale.ScaleWithColor(color.RGBA{255, 255, 255, 0}) // reddish flash
		m.HitTimer--
	}
	screen.DrawImage(m.Sprite, op)
}