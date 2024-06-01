package main

import (
	"bytes"
	"golang.org/x/image/font/gofont/gomedium"
	"image/color"
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/memojito/tgg-game/systems"
)

type myScene struct{}

// Type uniquely defines your game type
func (*myScene) Type() string {
	return "myGame"
}

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*myScene) Preload() {
	err := engo.Files.Load("textures/main-char.png",
		"textures/enemy-char.png",
		"tilemap/bg.tmx",
		"tilemap/mid.tmx",
		"tilemap/side.tmx",
		"tilemap/wall.tmx",
	)
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

	// register buttons
	engo.Input.RegisterButton("MoveRight", engo.KeyD)
	engo.Input.RegisterButton("MoveLeft", engo.KeyA)
	engo.Input.RegisterButton("Jump", engo.KeySpace)

	// background
	common.SetBackground(color.White)

	// add common systems
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&common.MouseSystem{})
	w.AddSystem(&common.CollisionSystem{Solids: common.CollisionGroup(1)}) // collision group 1 will be solid

	// Add custom systems
	w.AddSystem(&systems.PhysicsSystem{})
	w.AddSystem(&systems.LocationBuildingSystem{})
	w.AddSystem(&systems.Background{})
	w.AddSystem(&systems.Boundaries{})
	w.AddSystem(&systems.HUDSystem{})
	w.AddSystem(&systems.HUDTextSystem{})
	w.AddSystem(&systems.PlayerSystem{})
	//w.AddSystem(&systems.EnemySystem{})
}

func main() {
	opts := engo.RunOptions{
		Title:        "tgg",
		Width:        1280,
		Height:       640,
		NotResizable: true,
		FPSLimit:     60,

		//StandardInputs: true, // allows using arrow keys to move the camera around.
	}
	engo.Run(opts, &myScene{})

}
