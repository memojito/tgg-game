package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type physicsEntity struct {
	*ecs.BasicEntity
	*common.SpaceComponent

	*Movable
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
		// move by the factor g * h where g is const and h belongs to an entity increases over time and
		// resets if collision occurs
		if e.Movable.n == nMaxGlobal && e.Movable.h <= e.Movable.hMax {
			e.h *= e.h
		}
		Move(e.SpaceComponent, 0, e.h*g)
	}
}

// Add an entity to the physics system
func (ps *PhysicsSystem) Add(basic *ecs.BasicEntity, space *common.SpaceComponent, movable *Movable) {
	ps.entities = append(ps.entities, physicsEntity{basic, space, movable})
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
