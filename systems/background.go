package systems

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

// Tile forms the background. It's mostly static.
type Tile struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// Background is a piece of a TileMap which forms the background
type Background struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// New adds the Background.
func (t *Background) New(w *ecs.World) {
	log.Println("Background was added to the Scene")

	// load top background
	top, err := engo.Files.Resource("tilemap/top.tmx")
	if err != nil {
		panic(err)
	}
	topData := top.(common.TMXResource).Level
	topData.RenderOrder = "top-up"

	// loop through TileLayers from the .tmx and add each tile to a slice
	tiles := make([]*Background, 0)
	for _, tileLayer := range topData.TileLayers {
		for _, tileElement := range tileLayer.Tiles {
			if tileElement.Image != nil {
				tile := &Background{BasicEntity: ecs.NewBasic()}
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

func (*Background) Update(dt float32) {

}

func (*Background) Remove(entity ecs.BasicEntity) {}
