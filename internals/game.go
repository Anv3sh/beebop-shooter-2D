package internals

import (
	"fmt"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	// "github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	// "math"
)

const speed = 3.0

const defaultWindowW = 640
const defaultWindowH = 480


type Game struct{
	Player *Player
	WindowW float64
	WindowH float64
	Space *Space
	ScrollY float64
	GameOver bool
}

func GameInitAndRun() error{
	g := &Game{
		Player: &Player{
			Sprite: MustLoadImage(RAPTOR),
			XCoordinate: float64(defaultWindowW/2 - MustLoadImage(RAPTOR).Bounds().Dx()),
			YCoordinate: defaultWindowH-30,
			LeftLaser: []*Laser{},
			RightLaser: []*Laser{},
			ShootSpeed: 2,
			MoveSpeed: 2.5,
			SpawnRate: 0.1,
		},
		WindowW: defaultWindowW, 
		WindowH: defaultWindowH, 
		Space: &Space{
			Sprite: MustLoadImage(SPACE_BACKGROUND_PURPLE),
			Meteors: []*Meteor{},
			SpawnRate: 60,
		},
	}
	ebiten.SetWindowSize(int(g.WindowW),int(g.WindowH))
	ebiten.SetWindowTitle("Bebop Shooter 2D")
	return ebiten.RunGame(g)

}

func (g *Game) Update() error {
	if g.GameOver && ebiten.IsKeyPressed(ebiten.KeyR){
		g.resetGame()
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ){
		return ebiten.Termination
	}
	if g.Player.checkPlayerCollision(g.Space){
		g.GameOver = true

		return nil
	}
	g.Space.scrollSpace()
	g.Player.move()
	fmt.Println("hello")
	// clamp player if goes out of bounds
	g.Player.clamp_player(g.WindowW, g.WindowH)

	// check laser and metor collision
	g.Player.checkLaserCollision(g.Space)
	g.Space.spawnMeteor(g.WindowW)
	if inpututil.IsKeyJustPressed(ebiten.KeyX){
		g.Player.generateLaser()
	}

	g.Space.updateMeteors()
	g.Player.shoot()

	g.Player.reloadLaser(g.WindowW, g.WindowH)

	// log to check if the laser was getting destroyed if out of bounds
	// if len(g.Player.LeftLaser)>0 && g.Player.LeftLaser[0] != nil{
	// 	fmt.Println("NO")
	// }
	// if len(g.Player.LeftLaser)>0 && g.Player.LeftLaser[0] == nil{
	// 	fmt.Println("YES")
	// }
	g.Space.destroyMeteor(g.WindowW,g.WindowH)
	
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.GameOver {
        g.drawFinalScreen(screen)
        return
    }
	// Draw space and surroundings
	g.Space.drawSpace(screen)
	// Draw Player and Lasers
	g.Player.drawPlayer(screen)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint(g.Player.Score), int(g.WindowW)-30,1)
}

func(g *Game) drawFinalScreen(screen *ebiten.Image){
	gameOverText := "GAME OVER!"
	textWidth, textHeight := getTextWidthAndHeight(gameOverText)

	var x, y = int(g.WindowW/2), int(g.WindowH/2)
	var w, h = 60, 20
	red := color.RGBA{255, 0, 0, 255}
	vector.DrawFilledRect(screen, float32(x-textWidth/2), float32(y-textHeight), float32(w), float32(h), red, false)
	ebitenutil.DebugPrintAt(screen, gameOverText, x-textWidth/2, y-textHeight)
	restartText := "PRESS R to restart."
	textWidth, textHeight = getTextWidthAndHeight(restartText)
	ebitenutil.DebugPrintAt(screen, restartText, x-textWidth/2,int(g.WindowH-20))
}

func(g *Game) resetGame(){
	g.GameOver = false
	g.Space.Meteors = []*Meteor{}
	g.Player.LeftLaser = []*Laser{}
	g.Player.RightLaser = []*Laser{}
	g.Player.Score = 0
	g.Player.XCoordinate = float64(defaultWindowW/2 - g.Player.Sprite.Bounds().Dx())
	g.Player.YCoordinate = defaultWindowH-30
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}