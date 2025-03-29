package internals

import (

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	// "math"
)

const speed = 3.0

type Game struct{
	Player *Player
	LeftLaser *Laser
	RightLaser *Laser
	WindowW float64
	WindowH float64
	Space *ebiten.Image
}

func (g *Game) Update() error {
	if g.LeftLaser != nil && g.RightLaser!=nil{
		g.LeftLaser.Move(speed)
		g.RightLaser.Move(speed)
	}
	
	g.Player.move(speed)

	if ebiten.IsKeyPressed(ebiten.KeyQ){
		return ebiten.Termination
	}
	// clamp player if goes out of bounds
	g.Player.clamp_player(g.WindowW, g.WindowH)

	if inpututil.IsKeyJustPressed(ebiten.KeyX){
		g.LeftLaser = &Laser{Sprite:MustLoadImage(LASER_BLUE_16), XCoordinate: g.Player.XCoordinate, YCoordinate: g.Player.YCoordinate}
		g.RightLaser = &Laser{Sprite:MustLoadImage(LASER_BLUE_16), XCoordinate: g.Player.XCoordinate + float64(g.Player.Sprite.Bounds().Dx()), YCoordinate: g.Player.YCoordinate}
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
	if g.LeftLaser != nil && g.RightLaser !=nil{
		opLeftLaser.GeoM.Scale(0.5,0.5)
		opLeftLaser.GeoM.Translate(g.LeftLaser.XCoordinate, g.LeftLaser.YCoordinate)

		opRightLaser.GeoM.Scale(0.5,0.5)
		opRightLaser.GeoM.Translate(g.RightLaser.XCoordinate, g.RightLaser.YCoordinate)
	}
	opPlayer.GeoM.Translate(g.Player.XCoordinate, g.Player.YCoordinate)

	// w, h := Space.Size()
	// screenW, screenH := 640.0, 480.0
	// fmt.Println("original back size:",w,h)
	// opPlayer.GeoM.Scale(g.Scale,g.Scale)
	// opBack.GeoM.Scale()
	// opPlayer.GeoM.Rotate(45.0 * math.Pi / 180.0)

	screen.DrawImage(g.Space, opBack)
	if g.LeftLaser != nil && g.RightLaser != nil{
		screen.DrawImage(g.LeftLaser.Sprite,opLeftLaser)
		screen.DrawImage(g.RightLaser.Sprite,opRightLaser)
	}
	screen.DrawImage(g.Player.Sprite, opPlayer)


}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}