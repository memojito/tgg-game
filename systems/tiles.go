package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"log"
)

// TileSystem is a piece of a TileMap which forms the background
type TileSystem struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// New is called when the system is added to the world.
// Adds the Background.
func (t *TileSystem) New(w *ecs.World) {
	log.Println("TileSystem was added to the Scene")

	// load background
	resource, err := engo.Files.Resource("tilemap/bg.tmx")
	if err != nil {
		panic(err)
	}
	tmxResource := resource.(common.TMXResource)
	levelData := tmxResource.Level

	// loop through TileLayers from the .tmx and add each tile to a slice
	tiles := make([]*TileSystem, 0)
	for _, tileLayer := range levelData.TileLayers {
		for _, tileElement := range tileLayer.Tiles {
			if tileElement.Image != nil {
				tile := &TileSystem{BasicEntity: ecs.NewBasic()}
				tile.RenderComponent = common.RenderComponent{
					Drawable: tileElement.Image,
					Scale:    engo.Point{1, 1},
				}
				tile.SpaceComponent = common.SpaceComponent{
					Position: tileElement.Point,
					Width:    0,
					Height:   0,
				}
				tiles = append(tiles, tile)
				tile.RenderComponent.SetZIndex(-1)
			}
		}
	}
	// add the tiles to the RenderSystem
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, v := range tiles {
				sys.Add(&v.BasicEntity, &v.RenderComponent, &v.SpaceComponent)
			}
		}
	}
}

func (*TileSystem) Update(dt float32) {

}

func (*TileSystem) Remove(entity ecs.BasicEntity) {}
