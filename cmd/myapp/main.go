package main

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/memojito/tgg-game/systems"
	"image/color"
	"log"
)

type myScene struct {
}

// Type uniquely defines your game type
func (*myScene) Type() string {
	return "myGame"
}

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*myScene) Preload() {
	err := engo.Files.Load("textures/main-char.png")
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
}

func main() {
	opts := engo.RunOptions{
		Title:          "tgg",
		Width:          1200,
		Height:         800,
		StandardInputs: true, // allows using arrow keys to move the camera around.
	}
	engo.Run(opts, &myScene{})

}
