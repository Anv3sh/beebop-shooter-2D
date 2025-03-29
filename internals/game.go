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
}

func (g *Game) Update() error {
	
	g.Player.shoot(speed)
	g.Player.move(speed)

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
	opBack:= &ebiten.DrawImageOptions{}
	opPlayer := &ebiten.DrawImageOptions{}
	opLeftLaser := &ebiten.DrawImageOptions{}
	opRightLaser := &ebiten.DrawImageOptions{}

	opBack.GeoM.Scale(2.5,2)

	// lasers
	if g.Player.LeftLaser != nil && g.Player.RightLaser !=nil{
		opLeftLaser.GeoM.Scale(0.5,0.5)
		opLeftLaser.GeoM.Translate(g.Player.LeftLaser.XCoordinate, g.Player.LeftLaser.YCoordinate)

		opRightLaser.GeoM.Scale(0.5,0.5)
		opRightLaser.GeoM.Translate(g.Player.RightLaser.XCoordinate, g.Player.RightLaser.YCoordinate)
	}
	opPlayer.GeoM.Translate(g.Player.XCoordinate, g.Player.YCoordinate)

	// w, h := Space.Size()
	// screenW, screenH := 640.0, 480.0
	// fmt.Println("original back size:",w,h)
	// opPlayer.GeoM.Scale(g.Scale,g.Scale)
	// opBack.GeoM.Scale()
	// opPlayer.GeoM.Rotate(45.0 * math.Pi / 180.0)

	screen.DrawImage(g.Space, opBack)
	if g.Player.LeftLaser != nil && g.Player.RightLaser != nil{
		screen.DrawImage(g.Player.LeftLaser.Sprite,opLeftLaser)
		screen.DrawImage(g.Player.RightLaser.Sprite,opRightLaser)
	}
	screen.DrawImage(g.Player.Sprite, opPlayer)


}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}