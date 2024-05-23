package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type physicsEntity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
}

type PhysicsSystem struct {
	entities []physicsEntity
	//world    *ecs.World
	Velocity float32
}

// New initializes the system
func (ps *PhysicsSystem) New(world *ecs.World) {}

// Update the system per frame
func (ps *PhysicsSystem) Update(dt float32) {
	for _, e := range ps.entities {
		e.SpaceComponent.Position.Y += 2
	}
}

// Add an entity to the physics system
func (ps *PhysicsSystem) Add(basic *ecs.BasicEntity, space *common.SpaceComponent) {
	ps.entities = append(ps.entities, physicsEntity{basic, space})
}

// Remove an entity from the system
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
