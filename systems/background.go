package systems

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

// TileSize represents the size of a tile in pixels.
const TileSize = 32

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
	common.CollisionComponent
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

	// load bg background
	mid, err := engo.Files.Resource("tilemap/mid.tmx")
	if err != nil {
		panic(err)
	}
	midData := mid.(common.TMXResource).Level

	// loop through TileLayers from the .tmx and add each tile to a slice
	midTiles := make([]*Background, 0)
	for _, tileLayer := range midData.TileLayers {
		for _, tileElement := range tileLayer.Tiles {
			if tileElement.Image != nil {
				tile := &Background{BasicEntity: ecs.NewBasic()}
				tile.RenderComponent = common.RenderComponent{
					Drawable: tileElement.Image,
					Scale:    engo.Point{1, 1},
				}
				tile.SpaceComponent = common.SpaceComponent{
					Position: engo.Point{
						X: tileElement.X + TileSize,
						Y: tileElement.Y + 8*TileSize,
					},
				}
				tile.RenderComponent.SetZIndex(-1)
				tile.CollisionComponent = common.CollisionComponent{Group: common.CollisionGroup(1)}
				midTiles = append(midTiles, tile)
			}
		}
	}
	// add the tiles to the RenderSystem and CollisionSystem
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, v := range midTiles {
				sys.Add(&v.BasicEntity, &v.RenderComponent, &v.SpaceComponent)
			}
		case *common.CollisionSystem:
			for _, v := range midTiles {
				sys.Add(&v.BasicEntity, &v.CollisionComponent, &v.SpaceComponent)
			}
		}
	}
}

func (*Background) Update(dt float32) {

}

func (*Background) Remove(entity ecs.BasicEntity) {}
