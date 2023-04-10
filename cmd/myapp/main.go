package main

import (
	"bytes"
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/memojito/tgg-game/systems"
	"golang.org/x/image/font/gofont/gomedium"
	"image/color"
	"log"
)

type myScene struct {
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

	// load font
	err = engo.Files.LoadReaderData("go.ttf", bytes.NewReader(gomedium.TTF))
	if err != nil {
		log.Fatalf("failed to preload font: %v", err)
		return
	}
}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (*myScene) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)
	engo.Input.RegisterButton("AddLocation", engo.KeyF)
	common.SetBackground(color.White)

	// Add common systems
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&common.MouseSystem{})
	kbs := common.NewKeyboardScroller(
		400, engo.DefaultHorizontalAxis,
		engo.DefaultVerticalAxis)
	w.AddSystem(kbs)

	// Add custom systems
	w.AddSystem(&systems.LocationBuildingSystem{})
	w.AddSystem(&systems.HUDSystem{})
	w.AddSystem(&systems.HUDTextSystem{})
	w.AddSystem(&systems.TileSystem{})
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
