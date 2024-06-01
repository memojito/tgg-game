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
		Position: engo.Point{X: PlayerSize, Y: PlayerSize},
		Width:    PlayerSize,
		Height:   PlayerSize,
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

	ps.player.Movable.h = hGlobal
	ps.player.Movable.hMax = hMaxGlobal
	ps.player.Movable.n = 0

	for _, system := range ps.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&ps.player.BasicEntity, &ps.player.RenderComponent, &ps.player.SpaceComponent)
		case *common.CollisionSystem:
			sys.Add(&ps.player.BasicEntity, &ps.player.CollisionComponent, &ps.player.SpaceComponent)
		case *PhysicsSystem:
			sys.Add(&ps.player.BasicEntity, &ps.player.SpaceComponent, &ps.player.Movable)
		}
	}
}

// Update the system per frame.
func (ps *PlayerSystem) Update(dt float32) {
	if engo.Input.Button("MoveRight").Down() {
		Move(&ps.player.SpaceComponent, 5, 0)
	}

	if engo.Input.Button("MoveLeft").Down() {
		Move(&ps.player.SpaceComponent, -5, 0)
	}

	if engo.Input.Button("Jump").JustPressed() {
		ps.player.Movable.h = hGlobal
		engo.Mailbox.Dispatch(JumpTextMessage{
			BasicEntity:    ps.player.BasicEntity,
			SpaceComponent: ps.player.SpaceComponent,
		})
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

	if ps.player.Movable.n == nGlobal {
		ps.player.Movable.n = 0
	}
	ps.player.Movable.n++
	log.Printf("n: %d", ps.player.Movable.n)
	log.Printf("h: %f", ps.player.Movable.h)

	engo.Mailbox.Listen("CollisionMessage", func(m engo.Message) {
		_, ok := m.(common.CollisionMessage)
		if !ok {
			return
		}
		ps.player.Movable.h = hGlobal
		//log.Printf("Collision in CollisionGroup: %v", msg.Entity.Main)
	})
}

func (ps *PlayerSystem) Remove(entity ecs.BasicEntity) {

}
