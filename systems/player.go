package systems

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

const playerSize = 64

// Movable provides the Move method
type Movable struct {
	common.SpaceComponent
}

// Player of the game
type Player struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	common.CollisionComponent

	Movable
}

// PlayerSystem controls the Player
type PlayerSystem struct {
	world  *ecs.World
	player Player
}

// New is the initialization of the system.
func (ps *PlayerSystem) New(w *ecs.World) {
	ps.world = w
	log.Println("PlayerSystem was added to the Scene")

	ps.player.BasicEntity = ecs.NewBasic()
	ps.player.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 100, Y: 100},
		Width:    100,
		Height:   100,
	}

	texture, err := common.LoadedSprite("textures/main-char.png")
	if err != nil {
		log.Printf("failed to load texture: %v", err)
	}

	ps.player.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{X: 1, Y: 1},
	}

	ps.player.CollisionComponent = common.CollisionComponent{Main: common.CollisionGroup(1)}

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

func (m *common.SpaceComponent) Move(x float32, y float32) {
	m.Position.X += x
	m.Position.Y += y
}

func (m *Movable) Move(x float32, y float32) {
	m.SpaceComponent.Position.X += x
	m.SpaceComponent.Position.Y += y
}

// Update the system per frame.
func (ps *PlayerSystem) Update(dt float32) {
	if engo.Input.Button("MoveRight").Down() {
		ps.player.Move(-5, 0)
	}

	if engo.Input.Button("MoveLeft").Down() {
		ps.player.SpaceComponent.Position.X -= 5
	}

	if engo.Input.Button("Jump").JustPressed() {
		ps.player.SpaceComponent.Position.Y -= 30
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

	engo.Mailbox.Listen("CollisionMessage", func(m engo.Message) {
		msg, ok := m.(common.CollisionMessage)
		if !ok {
			return
		}
		log.Printf("Collision in CollisionGroup: %v", msg.Entity.Main)
	})
}

func (ps *PlayerSystem) Remove(entity ecs.BasicEntity) {

}
