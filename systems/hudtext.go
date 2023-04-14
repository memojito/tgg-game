package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"image/color"
	"log"
	"strings"
)

// Text is an entity containing text printed to the screen
type Text struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}

type HUDText struct {
	Text
	y    float32
	text string
}

// HUDTextEntity is an entity for the text system. This keeps track of the position
// size and text associated with that position.
type HUDTextEntity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	text []string
}

// HUDTextSystem prints the text to our HUDSystem based on the current state of the game
type HUDTextSystem struct {
	hudText  []HUDText
	corner   Text
	entities []HUDTextEntity
}

// HUDTextMessage updates the HUDSystem text based on messages sent from other systems
type HUDTextMessage struct {
	ecs.BasicEntity
	common.SpaceComponent
	text []string
}

const HUDTextMessageType string = "HUDTextMessage"

var PrimaryColor color.Color = color.RGBA{R: 158, G: 74, B: 224, A: 225}

// Type implements the engo.Message Interface
func (HUDTextMessage) Type() string {
	return HUDTextMessageType
}

// CenterString centers a string
func CenterString(str string, width int) string {
	spaces := int(float64(width-len(str)) / 2)
	return strings.Repeat(" ", spaces) + str + strings.Repeat(" ", width-(spaces+len(str)))
}

// New is called when the system is added to the world.
// Adds text to our HUDSystem that will update based on the state of the game.
func (h *HUDTextSystem) New(w *ecs.World) {
	log.Println("HUDTextSystem was added to the Scene")

	fnt := &common.Font{
		URL:  "go.ttf",
		FG:   PrimaryColor,
		Size: 24,
	}

	err := fnt.CreatePreloaded()
	if err != nil {
		log.Printf("failed to load preloaded font: %v", err)
		return
	}

	h.hudText = []HUDText{
		{Text: Text{BasicEntity: ecs.NewBasic()}, y: engo.WindowHeight() - 240, text: "speed:"},
		{Text: Text{BasicEntity: ecs.NewBasic()}, y: engo.WindowHeight() - 210, text: "jumps:"},
		{Text: Text{BasicEntity: ecs.NewBasic()}, y: engo.WindowHeight() - 180, text: "kills:"},
		{Text: Text{BasicEntity: ecs.NewBasic()}, y: engo.WindowHeight() - 150, text: "points:"},
	}
	for i := range h.hudText {
		hudText := &h.hudText[i]
		hudText.RenderComponent.Drawable = common.Text{
			Font: fnt,
			Text: CenterString(hudText.text, 24),
		}
		hudText.SetShader(common.TextHUDShader)
		hudText.RenderComponent.SetZIndex(1001)
		hudText.SpaceComponent.Position = engo.Point{X: 0, Y: hudText.y}

		for _, system := range w.Systems() {
			switch sys := system.(type) {
			case *common.RenderSystem:
				sys.Add(&hudText.BasicEntity, &hudText.RenderComponent, &hudText.SpaceComponent)
			}
		}
	}

	h.corner = Text{BasicEntity: ecs.NewBasic()}
	h.corner.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: "move somewhere with W,A,S,D",
	}
	h.corner.SetShader(common.TextHUDShader)
	h.corner.RenderComponent.SetZIndex(1001)
	h.corner.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 0, Y: engo.WindowHeight() - 40},
	}
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&h.corner.BasicEntity, &h.corner.RenderComponent, &h.corner.SpaceComponent)
		}
	}

	engo.Mailbox.Listen(HUDTextMessageType, func(m engo.Message) {
		msg, ok := m.(HUDTextMessage)
		if !ok {
			return
		}
		for _, system := range w.Systems() {
			switch sys := system.(type) {
			case *HUDTextSystem:
				sys.Add(&msg.BasicEntity, &msg.SpaceComponent, msg.text)
			}
		}
	})
}

// Add adds an entity to the system
func (h *HUDTextSystem) Add(b *ecs.BasicEntity, s *common.SpaceComponent, text []string) {
	h.entities = append(h.entities, HUDTextEntity{b, s, text})
}

// Update is called each frame to update the system.
func (h *HUDTextSystem) Update(dt float32) {
	//for _, e := range h.entities {
	//	if e.MouseComponent.Clicked {
	//		txt := h.text1.RenderComponent.Drawable.(common.Text)
	//		txt.Text = e.Line1
	//		h.text1.RenderComponent.Drawable = txt
	//		txt = h.text2.RenderComponent.Drawable.(common.Text)
	//		txt.Text = e.Line2
	//		h.text2.RenderComponent.Drawable = txt
	//		txt = h.text3.RenderComponent.Drawable.(common.Text)
	//		txt.Text = e.Line3
	//		h.text3.RenderComponent.Drawable = txt
	//		txt = h.text4.RenderComponent.Drawable.(common.Text)
	//		txt.Text = e.Line4
	//		h.text4.RenderComponent.Drawable = txt
	//	}
	//}
	switch {
	case engo.Input.Button("MoveUp").Down():
		txt := h.corner.RenderComponent.Drawable.(common.Text)
		txt.Text = "moving up"
		h.corner.RenderComponent.Drawable = txt
	case engo.Input.Button("MoveDown").Down():
		txt := h.corner.RenderComponent.Drawable.(common.Text)
		txt.Text = "moving down"
		h.corner.RenderComponent.Drawable = txt
	case engo.Input.Button("MoveRight").Down():
		txt := h.corner.RenderComponent.Drawable.(common.Text)
		txt.Text = "moving right"
		h.corner.RenderComponent.Drawable = txt
	case engo.Input.Button("MoveLeft").Down():
		txt := h.corner.RenderComponent.Drawable.(common.Text)
		txt.Text = "moving left"
		h.corner.RenderComponent.Drawable = txt
	default:
		txt := h.corner.RenderComponent.Drawable.(common.Text)
		txt.Text = "move somewhere with W,A,S,D"
		h.corner.RenderComponent.Drawable = txt
	}
}

// Remove takes an enitty out of the system.
// It does nothing as HUDTextSystem has no entities.
func (h *HUDTextSystem) Remove(entity ecs.BasicEntity) {}
