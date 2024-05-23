package systems

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type physicsEntity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
}

type PhysicsSystem struct {
	entities []physicsEntity
	world    *ecs.World
	Velocity float32
}

// New initializes the system
func (ps *PhysicsSystem) New(world *ecs.World) {}

// Update the system per frame
func (ps *PhysicsSystem) Update(dt float32) {
	//Set World components to the Render/Space Components
	for _, e := range ps.entities {
		e.SpaceComponent.Position.Y += 5
	}

	//for _, system := range ps.world.Systems(){
	//	switch sys := system.(type) {
	//	case *common.SpaceComponent:
	//
	//	}
	//}
}

// Add an entity to the physics system
func (ps *PhysicsSystem) Add(basic *ecs.BasicEntity, space *common.SpaceComponent) {
	log.Printf("added %d to the system", basic.GetBasicEntity().ID())
	ps.entities = append(ps.entities, physicsEntity{basic, space})
}

func (ps *PhysicsSystem) Remove(basic ecs.BasicEntity) {
	del := -1
	for index, e := range ps.entities {
		if e.BasicEntity.ID() == basic.ID() {
			del = index
			break
		}
	}
	if del >= 0 {
		ps.entities = append(ps.entities[:del], ps.entities[del+1:]...)
	}
}
