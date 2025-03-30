package internals

import "github.com/hajimehoshi/ebiten/v2"



type Space struct{
	Sprite *ebiten.Image
	ScrollY float64
	Meteors []*Meteor
	SpawnTick int
	SpawnRate int //rate of spawn say after x frames
	// TODO: add enemies
	// Enemies *[]Enemy
}

func (s *Space) scrollSpace(){
	s.ScrollY += 1 // speed of scrolling
	h := s.Sprite.Bounds().Dy()
	if s.ScrollY >= float64(h) {
		s.ScrollY = 0 // reset to loop
	}

}

func (s *Space) drawSpace(screen *ebiten.Image){
	h := s.Sprite.Bounds().Dy()
	hf := float64(h)

	// First background
	opBack1 := &ebiten.DrawImageOptions{}
	opBack1.GeoM.Scale(2.5,2)
	opBack1.GeoM.Translate(0, s.ScrollY)
	screen.DrawImage(s.Sprite, opBack1)

	// Second background (above it)
	opBack2 := &ebiten.DrawImageOptions{}
	opBack2.GeoM.Scale(2.5,2)
	opBack2.GeoM.Translate(0, s.ScrollY - hf)
	screen.DrawImage(s.Sprite, opBack2)

	for _, meteor := range s.Meteors{
		meteor.drawMeteor(screen)
	}
}

func (s *Space) moveMeteors(){
	for _, meteor := range s.Meteors{
		meteor.moveMeteor()
	}
}


func (s *Space) spawnMeteor(windowW float64){
	if s.SpawnTick >0 {
		s.SpawnTick--
		return
	}
	s.Meteors = append(s.Meteors,generateMeteor(windowW))

	s.SpawnTick = s.SpawnRate
}

func (s *Space) destroyMeteor(windowW float64, windowH float64){

	newMeteors := make([]*Meteor, 0, len(s.Meteors))
	for _, meteor := range s.Meteors{
		if !meteor.isMeteorOutOfBounds(windowW, windowH){
			newMeteors = append(newMeteors, meteor)
		}
	}
	s.Meteors = newMeteors
}