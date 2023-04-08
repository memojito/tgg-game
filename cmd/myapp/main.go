package main

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
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
	world.AddSystem(&common.RenderSystem{})

	location := Location{BasicEntity: ecs.NewBasic()}
	location.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{10, 10},
		Width:    303,
		Height:   641,
	}

	texture, err := common.LoadedSprite("textures/main-char.png")
	if err != nil {
		log.Printf("failed to load texture: %v", err)
	}

	location.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{1, 1},
	}

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&location.BasicEntity, &location.RenderComponent, &location.SpaceComponent)
		}
	}

	common.SetBackground(color.White)

}

type Location struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func main() {
	opts := engo.RunOptions{
		Title:  "tgg",
		Width:  400,
		Height: 400,
	}
	engo.Run(opts, &myScene{})

}
