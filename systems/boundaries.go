package systems

import (
	"fmt"
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type Bounder interface {
	Configure(engo.Resource) []*Boundaries
}

type BoundaryEntity struct {
	resource engo.Resource
}

type Boundaries struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	common.CollisionComponent
}

// NewBoundaryEntity creates a new BoundaryEntity with the resource from the given URL
// It returns an error if the resource cannot be loaded.
func NewBoundaryEntity(url string) (*BoundaryEntity, error) {
	r, err := engo.Files.Resource(url)
	if err != nil {
		return &BoundaryEntity{}, fmt.Errorf("failed to load resource from %s: %w", url, err)
	}
	return &BoundaryEntity{resource: r}, nil
}

// Configure the position and collision of the Boundary
func (e BoundaryEntity) Configure(position engo.Point, isCollidable bool, collisionGroup ...byte) []*Boundaries {
	data := e.resource.(common.TMXResource).Level

	// loop through TileLayers from the .tmx and add each tile to a slice
	boundaries := make([]*Boundaries, 0)
	for _, tileLayer := range data.TileLayers {
		for _, tileElement := range tileLayer.Tiles {
			if tileElement.Image != nil {
				tile := &Boundaries{BasicEntity: ecs.NewBasic()}
				tile.RenderComponent = common.RenderComponent{
					Drawable: tileElement.Image,
					Scale:    engo.Point{X: 1, Y: 1},
				}
				tile.SpaceComponent = common.SpaceComponent{
					Position: engo.Point{
						X: tileElement.X + position.X,
						Y: tileElement.Y + position.Y,
					},
				}

				tile.RenderComponent.SetZIndex(-1)
				if isCollidable && len(collisionGroup) > 0 {
					tile.CollisionComponent = common.CollisionComponent{Group: common.CollisionGroup(collisionGroup[0])}
				} else if isCollidable {
					// using default collision group 1
					tile.CollisionComponent = common.CollisionComponent{Group: common.CollisionGroup(1)}
				}
				boundaries = append(boundaries, tile)
			}
		}
	}
	return boundaries
}

// New is the initialization of boundaries.
func (bs *Boundaries) New(w *ecs.World) {
	log.Println("BoundarySystem was added to the world")

	mid, err := NewBoundaryEntity("tilemap/mid.tmx")
	if err != nil {
		log.Fatalf("Error creating BoundaryEntity: %v", err)
	}

	midBoundaries := mid.Configure(engo.Point{X: TileSize * 2, Y: TileSize * 9}, true)

	// load the side element
	side, err := NewBoundaryEntity("tilemap/side.tmx")
	if err != nil {
		log.Fatalf("Error creating BoundaryEntity: %v", err)
	}
	sideBoundaries := side.Configure(engo.Point{}, true)

	// load the wall (top and bottom) element
	wall, err := NewBoundaryEntity("tilemap/wall.tmx")
	if err != nil {
		log.Fatalf("Error creating BoundaryEntity: %v", err)
	}
	topBoundaries := wall.Configure(engo.Point{X: TileSize * 2, Y: 0}, true)
	bottomBoundaries := wall.Configure(engo.Point{X: TileSize * 2, Y: TileSize * 18}, true)

	boundaries := append(midBoundaries, sideBoundaries...)
	boundaries = append(boundaries, topBoundaries...)
	boundaries = append(boundaries, bottomBoundaries...)

	// add the tiles to the RenderSystem and CollisionSystem
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, v := range boundaries {
				sys.Add(&v.BasicEntity, &v.RenderComponent, &v.SpaceComponent)
			}
		case *common.CollisionSystem:
			for _, v := range boundaries {
				sys.Add(&v.BasicEntity, &v.CollisionComponent, &v.SpaceComponent)
			}
		}
	}
}

func (bs *Boundaries) Update(dt float32) {

}

func (bs *Boundaries) Remove(entity ecs.BasicEntity) {

}
