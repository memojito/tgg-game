package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"log"
)

// Player of the game
type Player struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// PlayerSystem controls the Player
type PlayerSystem struct {
	world  *ecs.World
	player Player
}

func Gravity(position float32, g float32) float32 {
	position += g
	return position
}

var NullLvl float32 = 190
var g float32 = 0.6

// New is called when the system is added to the world.
func (ps *PlayerSystem) New(w *ecs.World) {
	ps.world = w
	log.Println("PlayerSystem was added to the Scene")

	ps.player.BasicEntity = ecs.NewBasic()
	ps.player.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{100, NullLvl},
		Width:    64,
		Height:   64,
	}

	texture, err := common.LoadedSprite("textures/main-char.png")
	if err != nil {
		log.Printf("failed to load texture: %v", err)
	}

	ps.player.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{1, 1},
	}

	for _, system := range ps.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&ps.player.BasicEntity, &ps.player.RenderComponent, &ps.player.SpaceComponent)
		}
	}
}

// Update is called each frame to update the system.
func (ps *PlayerSystem) Update(dt float32) {
	//if engo.Input.Button("MoveUp").Down() {
	//	ps.player.SpaceComponent.Position.Y -= 5
	//}
	//if engo.Input.Button("MoveDown").Down() {
	//	ps.player.SpaceComponent.Position.Y += 5
	//}
	if engo.Input.Button("MoveRight").Down() {
		ps.player.SpaceComponent.Position.X += 5
	}
	if engo.Input.Button("MoveLeft").Down() {
		ps.player.SpaceComponent.Position.X -= 5
	}
	if engo.Input.Button("Jump").JustPressed() {
		ps.player.SpaceComponent.Position.Y -= 80
	}
	if ps.player.SpaceComponent.Position.Y < NullLvl {
		g += 0.06 * g * g
		ps.player.SpaceComponent.Position.Y = Gravity(ps.player.SpaceComponent.Position.Y, g)
	}
	if ps.player.SpaceComponent.Position.Y >= NullLvl {
		g = 0.6
	}

	log.Println(g)
}

// Remove takes an enitty out of the system.
func (ps *PlayerSystem) Remove(entity ecs.BasicEntity) {

}
