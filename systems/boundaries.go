package systems

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type Boundary struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.CollisionComponent
}

type BoundarySystem struct {
	world    *ecs.World
	boundary Boundary
}

func (bs *BoundarySystem) New(w *ecs.World) {
	bs.world = w
	log.Println("BoundarySystem was added to the world")

	bs.boundary.BasicEntity = ecs.NewBasic()
	bs.boundary.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{0, 0},
		Width:    0,
		Height:   engo.WindowHeight(),
	}

	bs.boundary.CollisionComponent = common.CollisionComponent{}

	for _, system := range bs.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&bs.boundary.BasicEntity, nil, &bs.boundary.SpaceComponent)
		case *common.CollisionSystem:
			sys.Add(&bs.boundary.BasicEntity, &bs.boundary.CollisionComponent, &bs.boundary.SpaceComponent)
		}
	}
}

func (bs *BoundarySystem) Update(dt float32) {

}

func (bs *BoundarySystem) Remove(entity ecs.BasicEntity) {

}
