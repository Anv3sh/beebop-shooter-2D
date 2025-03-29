package internals

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	// "math"
)

const speed = 3.0

type Game struct{
	Player *Player
	WindowW float64
	WindowH float64
	Space *ebiten.Image
	ScrollY float64
}

func GameInitAndRun() error{
	g := &Game{
		Player: &Player{
			Sprite: MustLoadImage(RAPTOR),
			XCoordinate: 250,
			YCoordinate: 200,
			LeftLaser: make([]*Laser, 0, 1),
			RightLaser: make([]*Laser, 0, 1),
			ShootSpeed: 2,
			MoveSpeed: 2.5,
		},
		WindowW: 640, 
		WindowH: 480, 
		Space: MustLoadImage(SPACE_BACKGROUND_PURPLE),
	}
	ebiten.SetWindowSize(int(g.WindowW),int(g.WindowH))
	ebiten.SetWindowTitle("Beebop Shooter 2D")
	return ebiten.RunGame(g)

}

func (g *Game) Update() error {
	g.ScrollY += 1 // speed of scrolling
	h := g.Space.Bounds().Dy()
	if g.ScrollY >= float64(h) {
		g.ScrollY = 0 // reset to loop
	}
	g.Player.reloadLaser(g.WindowW, g.WindowH)

	// log to check if the laser was getting destroyed if out of bounds
	// if len(g.Player.LeftLaser)>0 && g.Player.LeftLaser[0] != nil{
	// 	fmt.Println("NO")
	// }
	// if len(g.Player.LeftLaser)>0 && g.Player.LeftLaser[0] == nil{
	// 	fmt.Println("YES")
	// }

	g.Player.shoot()
	g.Player.move()

	if ebiten.IsKeyPressed(ebiten.KeyQ){
		return ebiten.Termination
	}
	// clamp player if goes out of bounds
	g.Player.clamp_player(g.WindowW, g.WindowH)

	if inpututil.IsKeyJustPressed(ebiten.KeyX){
		g.Player.generateLaser()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	h := g.Space.Bounds().Dy()
	hf := float64(h)

	// First background
	opBack1 := &ebiten.DrawImageOptions{}
	opBack1.GeoM.Scale(2.5,2)
	opBack1.GeoM.Translate(0, g.ScrollY)
	screen.DrawImage(g.Space, opBack1)

	// Second background (above it)
	opBack2 := &ebiten.DrawImageOptions{}
	opBack2.GeoM.Scale(2.5,2)
	opBack2.GeoM.Translate(0, g.ScrollY - hf)
	screen.DrawImage(g.Space, opBack2)

	// Draw Player and Lasers
	g.Player.drawPlayer(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}