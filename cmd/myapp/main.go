package main

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/memojito/tgg-game/systems"
	"image"
	"image/color"
	"log"
)

type myScene struct {
}

type HUD struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

type Tile struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// Type uniquely defines your game type
func (*myScene) Type() string {
	return "myGame"
}

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*myScene) Preload() {
	err := engo.Files.Load("textures/main-char.png", "tilemap/bg.tmx")
	if err != nil {
		log.Fatalf("failed to preload texture: %v", err)
		return
	}
}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (*myScene) Setup(u engo.Updater) {
	world, _ := u.(*ecs.World)
	engo.Input.RegisterButton("AddLocation", engo.KeyF)
	common.SetBackground(color.White)

	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})
	kbs := common.NewKeyboardScroller(
		400, engo.DefaultHorizontalAxis,
		engo.DefaultVerticalAxis)
	world.AddSystem(kbs)

	world.AddSystem(&systems.LocationBuildingSystem{})

	hud := HUD{BasicEntity: ecs.NewBasic()}
	hud.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{0, engo.WindowHeight() - 300},
		Width:    200,
		Height:   200,
	}

	hudImage := image.NewUniform(color.RGBA{205, 205, 205, 255})
	hudNRGBA := common.ImageToNRGBA(hudImage, 200, 200)
	hudImageObj := common.NewImageObject(hudNRGBA)
	hudTexture := common.NewTextureSingle(hudImageObj)

	hud.RenderComponent = common.RenderComponent{
		Repeat:   common.Repeat,
		Drawable: hudTexture,
		Scale:    engo.Point{1, 1},
	}
	hud.RenderComponent.SetShader(common.HUDShader)
	hud.RenderComponent.SetZIndex(1)

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&hud.BasicEntity, &hud.RenderComponent, &hud.SpaceComponent)
		}
	}

	// load background
	resource, err := engo.Files.Resource("tilemap/bg.tmx")
	if err != nil {
		panic(err)
	}
	tmxResource := resource.(common.TMXResource)
	levelData := tmxResource.Level

	// loop through TileLayers from the .tmx and add each tile to a slice
	tiles := make([]*Tile, 0)
	for _, tileLayer := range levelData.TileLayers {
		for _, tileElement := range tileLayer.Tiles {
			if tileElement.Image != nil {
				tile := &Tile{BasicEntity: ecs.NewBasic()}
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
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, v := range tiles {
				sys.Add(&v.BasicEntity, &v.RenderComponent, &v.SpaceComponent)
			}
		}
	}

	//common.CameraBounds = levelData.Bounds()
}

func main() {
	opts := engo.RunOptions{
		Title:          "tgg",
		Width:          1280,
		Height:         640,
		StandardInputs: true, // allows using arrow keys to move the camera around.
	}
	engo.Run(opts, &myScene{})

}
