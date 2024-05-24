package systems

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

const enemySize = 64

// Enemy of the game
type Enemy struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	common.CollisionComponent
}

// EnemySystem controls the Enemy
type EnemySystem struct {
	world *ecs.World
	enemy Enemy
}

// New is the initialization of the system.
func (ps *EnemySystem) New(w *ecs.World) {
	ps.world = w
	log.Println("EnemySystem was added to the Scene")

	ps.enemy.BasicEntity = ecs.NewBasic()
	ps.enemy.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 100, Y: 100},
		Width:    100,
		Height:   100,
	}

	//s := common.Shape{}
	//s.Lines = make([]engo.Line, 4)
	//s.Lines = append(s.Lines, engo.Line{
	//	P1: engo.Point{0, -1},
	//	P2: engo.Point{0, PlayerSize + 1},
	//})
	//s.Lines = append(s.Lines, engo.Line{
	//	P1: engo.Point{-1, PlayerSize + 1},
	//	P2: engo.Point{PlayerSize + 1, PlayerSize + 1},
	//})
	//s.Lines = append(s.Lines, engo.Line{
	//	P1: engo.Point{PlayerSize, PlayerSize + 1},
	//	P2: engo.Point{PlayerSize, -1},
	//})
	//s.Lines = append(s.Lines, engo.Line{
	//	P1: engo.Point{PlayerSize + 1, 0},
	//	P2: engo.Point{-1, 0},
	//})
	//ps.enemy.SpaceComponent.AddShape(s)

	texture, err := common.LoadedSprite("textures/enemy-char.png")
	if err != nil {
		log.Printf("failed to load texture: %v", err)
	}

	ps.enemy.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{X: 1, Y: 1},
	}

	ps.enemy.CollisionComponent = common.CollisionComponent{Group: common.CollisionGroup(1)}

	for _, system := range ps.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&ps.enemy.BasicEntity, &ps.enemy.RenderComponent, &ps.enemy.SpaceComponent)
		case *common.CollisionSystem:
			sys.Add(&ps.enemy.BasicEntity, &ps.enemy.CollisionComponent, &ps.enemy.SpaceComponent)
			//case *PhysicsSystem:
			//	sys.Add(&ps.enemy.BasicEntity, &ps.enemy.SpaceComponent)
		}
	}
}

// Update the system per frame.
func (ps *EnemySystem) Update(dt float32) {

	//engo.Mailbox.Listen("CollisionMessage", func(m engo.Message) {
	//	msg, ok := m.(common.CollisionMessage)
	//	if !ok {
	//		return
	//	}
	//	log.Printf("Collision in CollisionGroup: %v", msg.Entity.Main)
	//})
}

func (ps *EnemySystem) Remove(entity ecs.BasicEntity) {

}
