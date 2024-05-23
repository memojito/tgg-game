package systems

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

// Player of the game
type Player struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	common.CollisionComponent
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

var NullLvl float32 = 0
var g float32 = 0.6

// New is the initialization of the system.
func (ps *PlayerSystem) New(w *ecs.World) {
	ps.world = w
	log.Println("PlayerSystem was added to the Scene")

	ps.player.BasicEntity = ecs.NewBasic()
	ps.player.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 100, Y: 100},
		Width:    64,
		Height:   64,
	}

	shape := common.Shape{
		Ellipse: common.Ellipse{Rx: 10, Ry: 10},
		Lines:   nil,
		N:       0,
	}
	ps.player.SpaceComponent.AddShape(shape)

	texture, err := common.LoadedSprite("textures/main-char.png")
	if err != nil {
		log.Printf("failed to load texture: %v", err)
	}

	ps.player.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{1, 1},
	}

	ps.player.CollisionComponent = common.CollisionComponent{}

	for _, system := range ps.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&ps.player.BasicEntity, &ps.player.RenderComponent, &ps.player.SpaceComponent)
		case *common.CollisionSystem:
			sys.Add(&ps.player.BasicEntity, &ps.player.CollisionComponent, &ps.player.SpaceComponent)
		case *PhysicsSystem:
			sys.Add(&ps.player.BasicEntity, &ps.player.SpaceComponent)
		}
	}
}

// Update the system per frame.
func (ps *PlayerSystem) Update(dt float32) {
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
		if g <= 10 {
			g += 0.06 * g * g
		}

		ps.player.SpaceComponent.Position.Y = Gravity(ps.player.SpaceComponent.Position.Y, g)
	}

	if ps.player.SpaceComponent.Position.Y >= NullLvl {
		g = 0.6
	}

	engo.Mailbox.Dispatch(common.CameraMessage{
		Axis:        common.YAxis,
		Value:       ps.player.SpaceComponent.Position.Y,
		Incremental: false,
	})

	engo.Mailbox.Dispatch(common.CameraMessage{
		Axis:        common.XAxis,
		Value:       ps.player.SpaceComponent.Position.X,
		Incremental: false,
	})

	engo.Mailbox.Listen("CollisionMessage", func(msg engo.Message) {
		log.Println("Collision")
	})
}

func (ps *PlayerSystem) Remove(entity ecs.BasicEntity) {

}
