package systems

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

// Background is a piece of a TileMap which forms the background
type Background struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// New adds the Background.
func (b *Background) New(w *ecs.World) {
	log.Println("Background was added to the Scene")

	// load bg background
	bg, err := engo.Files.Resource("tilemap/bg.tmx")
	if err != nil {
		panic(err)
	}
	bgData := bg.(common.TMXResource).Level

	// loop through TileLayers from the .tmx and add each tile to a slice
	tiles := make([]*Background, 0)
	for _, tileLayer := range bgData.TileLayers {
		for _, tileElement := range tileLayer.Tiles {
			if tileElement.Image != nil {
				tile := &Background{BasicEntity: ecs.NewBasic()}
				tile.RenderComponent = common.RenderComponent{
					Drawable: tileElement.Image,
					Scale:    engo.Point{1, 1},
				}
				tile.SpaceComponent = common.SpaceComponent{
					Position: engo.Point{
						X: tileElement.X - TileSize,
						Y: tileElement.Y - TileSize,
					},
				}
				tiles = append(tiles, tile)
				tile.RenderComponent.SetZIndex(-2)
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
