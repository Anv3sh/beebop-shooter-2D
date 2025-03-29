package internals

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct{
	Sprite *ebiten.Image
	XCoordinate float64
	YCoordinate float64
	LeftLaser []*Laser
	RightLaser []*Laser
	ShootSpeed float64
	MoveSpeed float64
	// Scale float64
}


func (p *Player) move(){
	if ebiten.IsKeyPressed(ebiten.KeyW){
		p.YCoordinate -= p.MoveSpeed
		
	}
	if ebiten.IsKeyPressed(ebiten.KeyS){
		p.YCoordinate += p.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.XCoordinate -= p.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.XCoordinate += p.MoveSpeed
	}
	
}

func (p *Player) clamp_player(windowW float64, windowH float64){
	if p.YCoordinate < 0 {
		p.YCoordinate = 0
	}

	if p.YCoordinate>windowH - float64(p.Sprite.Bounds().Dy()){
		p.YCoordinate = windowH - float64(p.Sprite.Bounds().Dy())
	}

	if p.XCoordinate < 0 {
		p.XCoordinate = 0
	}
	if p.XCoordinate > windowW- float64(p.Sprite.Bounds().Dx()) {
		p.XCoordinate = windowW - float64(p.Sprite.Bounds().Dx())
	}
}

func (p *Player) shoot(){
	for _, laser := range p.LeftLaser {
		laser.YCoordinate -= p.ShootSpeed
	}

	for _, laser := range p.RightLaser {
		laser.YCoordinate -= p.ShootSpeed
	}
}

func (p *Player) generateLaser(){
	p.LeftLaser = append(p.LeftLaser, &Laser{Sprite:MustLoadImage(LASER_BLUE_16), XCoordinate: p.XCoordinate, YCoordinate: p.YCoordinate,})
	p.RightLaser = append(p.RightLaser, &Laser{Sprite:MustLoadImage(LASER_BLUE_16), XCoordinate: p.XCoordinate + float64(p.Sprite.Bounds().Dx())-5, YCoordinate: p.YCoordinate,})
}

func (p *Player) reloadLaser(windowW float64, windowH float64){
	// Safe memory-cleaning version:
	newLeft := make([]*Laser, 0, len(p.LeftLaser))
	for _, laser := range p.LeftLaser {
		if !laser.isLaserOutOfBounds(windowW, windowH) {
			newLeft = append(newLeft, laser)
		}
	}
	p.LeftLaser = newLeft

	newRight := make([]*Laser, 0, len(p.RightLaser))
	for _, laser := range p.RightLaser {
		if !laser.isLaserOutOfBounds(windowW, windowH) {
			newRight = append(newRight, laser)
		}
	}
	p.RightLaser = newRight

	// Unsafe memory-cleaning version:
	// if len(p.LeftLaser)>0 && len(p.RightLaser)>0{
	// 	if p.LeftLaser[0].isLaserOutOfBounds(windowW, windowH){
	// 		p.LeftLaser[0] = nil
	// 		p.LeftLaser = p.LeftLaser[1:]
	// 	}

	// 	if p.RightLaser[0].isLaserOutOfBounds(windowW, windowH){
	// 		p.RightLaser[0] = nil
	// 		p.RightLaser = p.RightLaser[1:]
	// 	}
	// }
}

func (p *Player) drawPlayer(screen *ebiten.Image){
	for _, laser := range p.LeftLaser{
		laser.drawLaser(screen)
	}
	for _, laser := range p.RightLaser{
		laser.drawLaser(screen)
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.XCoordinate, p.YCoordinate)
	screen.DrawImage(p.Sprite, op)
}