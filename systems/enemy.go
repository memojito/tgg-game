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
func (ps *EnemySystem) Update(dt float32) {}

func (ps *EnemySystem) Remove(entity ecs.BasicEntity) {

}
