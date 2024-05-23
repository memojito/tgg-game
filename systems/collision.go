package systems

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type collisionEntity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
}

type CollisionSystem struct {
	entities []collisionEntity
}

func (cs *CollisionSystem) New(world *ecs.World) {}

func (cs *CollisionSystem) Update(dt float32) {

}

// Add an entity to the physics system
func (cs *CollisionSystem) Add(basic *ecs.BasicEntity, space *common.SpaceComponent) {
	log.Printf("added %d to the system", basic.GetBasicEntity().ID())
	cs.entities = append(cs.entities, collisionEntity{basic, space})
}

func (cs *CollisionSystem) Remove(ecs.BasicEntity) {}
