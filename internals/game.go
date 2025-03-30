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
	Space *Space
	ScrollY float64
}

func GameInitAndRun() error{
	g := &Game{
		Player: &Player{
			Sprite: MustLoadImage(RAPTOR),
			XCoordinate: 250,
			YCoordinate: 200,
			LeftLaser: []*Laser{},
			RightLaser: []*Laser{},
			ShootSpeed: 2,
			MoveSpeed: 2.5,
			SpawnRate: 0.1,
		},
		WindowW: 640, 
		WindowH: 480, 
		Space: &Space{
			Sprite: MustLoadImage(SPACE_BACKGROUND_PURPLE),
			Meteors: []*Meteor{},
			SpawnRate: 60,
		},
	}
	ebiten.SetWindowSize(int(g.WindowW),int(g.WindowH))
	ebiten.SetWindowTitle("Beebop Shooter 2D")
	return ebiten.RunGame(g)

}

func (g *Game) Update() error {
	g.Space.scrollSpace()
	g.Player.move()
	// clamp player if goes out of bounds
	g.Player.clamp_player(g.WindowW, g.WindowH)

	// check laser and metor collision
	g.Player.checkLaserCollision(g.Space)
	g.Space.spawnMeteor(g.WindowW)
	if inpututil.IsKeyJustPressed(ebiten.KeyX){
		g.Player.generateLaser()
	}

	g.Space.moveMeteors()
	g.Player.shoot()

	g.Player.reloadLaser(g.WindowW, g.WindowH)

	// log to check if the laser was getting destroyed if out of bounds
	// if len(g.Player.LeftLaser)>0 && g.Player.LeftLaser[0] != nil{
	// 	fmt.Println("NO")
	// }
	// if len(g.Player.LeftLaser)>0 && g.Player.LeftLaser[0] == nil{
	// 	fmt.Println("YES")
	// }


	if ebiten.IsKeyPressed(ebiten.KeyQ){
		return ebiten.Termination
	}
	
	g.Space.destroyMeteor(g.WindowW,g.WindowH)
	
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw space and surroundings
	g.Space.drawSpace(screen)
	// Draw Player and Lasers
	g.Player.drawPlayer(screen)
	

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}