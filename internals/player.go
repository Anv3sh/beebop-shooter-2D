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
	SpawnTick float64
	SpawnRate float64
	Score int
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
	leftfiltered := []*Laser{}
	for _, laser := range p.LeftLaser {
		hit := laser.Update(p.ShootSpeed)
		if !hit{
			leftfiltered = append(leftfiltered, laser)
		}
	}
	p.LeftLaser = leftfiltered
	rightfiltered := []*Laser{}
	for _, laser := range p.RightLaser {
		hit := laser.Update(p.ShootSpeed)
		if !hit{
			rightfiltered = append(rightfiltered, laser)
		}
	}
	p.RightLaser = rightfiltered
	//fmt.Println()
}

func (p *Player) generateLaser(){
	if p.SpawnTick>0{
		p.SpawnTick--
		return
	}
	p.LeftLaser = append(p.LeftLaser, &Laser{Sprite:MustLoadImage(LASER_BLUE_16), XCoordinate: p.XCoordinate, YCoordinate: p.YCoordinate,})
	p.RightLaser = append(p.RightLaser, &Laser{Sprite:MustLoadImage(LASER_BLUE_16), XCoordinate: p.XCoordinate + float64(p.Sprite.Bounds().Dx())-30, YCoordinate: p.YCoordinate,})
	p.SpawnTick = p.SpawnRate
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
	op.GeoM.Scale(0.75,0.75)
	op.GeoM.Translate(p.XCoordinate, p.YCoordinate)
	screen.DrawImage(p.Sprite, op)
}

func (p *Player) checkPlayerCollision(space *Space) bool{
	for _, meteor := range space.Meteors{
		if meteor.Destroyed{
			continue
		}
		if isColliding(
			p.XCoordinate,
			p.YCoordinate,
			float64(p.Sprite.Bounds().Dx()),
			float64(p.Sprite.Bounds().Dy()),
			meteor.XCoordinate,
			meteor.YCoordinate,
			float64(meteor.Sprite.Bounds().Dx()),
			float64(meteor.Sprite.Bounds().Dy()),
		){
			return true
		}
	}
	return false
}

func (p *Player) checkLaserCollision(space *Space) {
	// newLeftLasers := []*Laser{}
	// newRightLasers := []*Laser{}
	// newMeteors := []*Meteor{}

	// track which meteors got hit
	// hitMap := make(map[*Meteor]bool)

	// check left lasers
	for _, laser := range p.LeftLaser {
		if laser.Hit{
			continue
		}
		// hit := false
		for _, meteor := range space.Meteors {
			if meteor.Destroyed {
				continue // already hit
			}
			if isColliding(
				laser.XCoordinate, laser.YCoordinate,
				float64(laser.Sprite.Bounds().Dx()), float64(laser.Sprite.Bounds().Dy()),
				meteor.XCoordinate, meteor.YCoordinate,
				float64(meteor.Sprite.Bounds().Dx()), float64(meteor.Sprite.Bounds().Dy()),
			) {
				// hitMap[meteor] = true
				// hit = true
				p.updateScore(METEOR)
				laser.Sprite=MustLoadImage(LASER_BLUE_COLLIDED)
				laser.Hit = true
				laser.HitTimer = 2
				meteor.HitTimer = 3
				meteor.Destroyed = true
				go MustPlay(LASER_HIT_SOUND)
				break
			}
		}
		// if !hit {
		// 	newLeftLasers = append(newLeftLasers, laser)
		// }
	}

	// check right lasers
	for _, laser := range p.RightLaser {
		if laser.Hit{
			continue
		}
		// hit := false
		for _, meteor := range space.Meteors {
			if meteor.Destroyed {
				continue
			}
			if isColliding(
				laser.XCoordinate, laser.YCoordinate,
				float64(laser.Sprite.Bounds().Dx()), float64(laser.Sprite.Bounds().Dy()),
				meteor.XCoordinate, meteor.YCoordinate,
				float64(meteor.Sprite.Bounds().Dx()), float64(meteor.Sprite.Bounds().Dy()),
			) {
				// hitMap[meteor] = true
				// hit = true
				p.updateScore(METEOR)
				laser.Sprite=MustLoadImage(LASER_BLUE_COLLIDED)
				laser.Hit = true
				laser.HitTimer = 2
				meteor.HitTimer = 3
				meteor.Destroyed = true
				go MustPlay(LASER_HIT_SOUND)
				break
			}
		}
		// if !hit {
		// 	newRightLasers = append(newRightLasers, laser)
		// }
	}

	// rebuild the meteor list (only non-hit ones)
	// for _, meteor := range space.Meteors {
	// 	if !hitMap[meteor] {
	// 		newMeteors = append(newMeteors, meteor)
	// 	}
	// }

	// assign updated slices
	// p.LeftLaser = newLeftLasers
	// p.RightLaser = newRightLasers
	// space.Meteors = newMeteors
}


func (p *Player) updateScore(enemyType string){
	if enemyType == METEOR{
		p.Score+=1
	}
}
