package main

import (
	"fmt"
	"os"
	"time"

	_ "embed"
	"net/http"

	_ "net/http/pprof"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sevenc-nanashi/sonolus-emulator/pkg/processor"
	log "github.com/sirupsen/logrus"
)

type Game struct {
	processor      *processor.Processor
	startTime      float64
	deltaTime      float64
	frameDeltaTime float64

	drawCalls []processor.DrawCall

	touchCounter int
	touches      []processor.TouchInfo
}

//go:embed perspectiveShader.kage
var perspective []byte

var perspectiveShader *ebiten.Shader
var breadImage *ebiten.Image

const aspectRatio = 16.0 / 9.0

func (g *Game) Update() error {
	if g.startTime == 0 {
		g.startTime = float64(time.Now().UnixMilli()) / 1000
	}
	g.deltaTime = float64(time.Now().UnixMilli())/1000 - g.startTime

	g.processor.Spawn()
	g.processor.Update(g.deltaTime, g.frameDeltaTime)

	g.updateTouches()

	g.frameDeltaTime = g.deltaTime - g.frameDeltaTime
	g.drawCalls = g.processor.DrawCalls
	g.processor.DrawCalls = make([]processor.DrawCall, 0)
	return nil
}

func (g *Game) updateTouches() {
	newTouches := make([]processor.TouchInfo, 0)
	for i := range g.touches {
		if g.touches[i].Status != processor.TouchStatusEnd {
			newTouches = append(newTouches, g.touches[i])
		}
	}
	g.touches = newTouches
	for i := range g.touches {
		currentX, currentY := ebiten.CursorPosition()
		g.touches[i].Status = processor.TouchStatusMiddle
		lastX, lastY := g.touches[i].X, g.touches[i].Y
		g.touches[i].X = float64(currentX)
		g.touches[i].Y = float64(currentY)
		g.touches[i].Dx = g.touches[i].X - lastX
		g.touches[i].Dy = g.touches[i].Y - lastY
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.touchCounter++

		x, y := ebiten.CursorPosition()

		g.touches = append(g.touches, processor.TouchInfo{
			Id:     g.touchCounter,
			Status: processor.TouchStatusStart,
			X:      float64(x),
			Y:      float64(y),
			Sx:     float64(x),
			Sy:     float64(y),
			Dx:     0,
			Dy:     0,
		})
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.touches[len(g.touches)-1].Status = processor.TouchStatusEnd
	}
	g.processor.Touch(g.touches)
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"TPS: %0.2f\nFPS: %0.2f\nEntities: %d\nDraw calls: %d\nDebugLog: %0.2f",
		ebiten.CurrentTPS(),
		ebiten.CurrentFPS(),
		len(g.processor.Entities),
		len(g.drawCalls),
		g.processor.DebugLog,
	))
	return

	for _, call := range g.drawCalls {
		// op := &ebiten.DrawRectShaderOptions{}
		// x1 := (call.X1/aspectRatio + 1) / 2
		// y1 := (call.Y1/aspectRatio + 1) / 2
		// x2 := (call.X2/aspectRatio + 1) / 2
		// y2 := (call.Y2/aspectRatio + 1) / 2
		// x3 := (call.X3/aspectRatio + 1) / 2
		// y3 := (call.Y3/aspectRatio + 1) / 2
		// x4 := (call.X4/aspectRatio + 1) / 2
		// y4 := (call.Y4/aspectRatio + 1) / 2

		// // dx1 := x1 - x1
		// // dy1 := y1 - y1
		// dx2 := x2 - x1
		// dy2 := y2 - y1
		// dx3 := x3 - x1
		// dy3 := y3 - y1
		// dx4 := x4 - x1
		// dy4 := y4 - y1

		// al := ((dx4 * dy2) - (dx2 * dy4)) - ((dx3 * dy2) - (dx2 * dy3))
		// bl := ((dx4*dy2)-(dx2*dy4))*x3 - ((dx3*dy2)-(dx2*dy3))*x4
		// cl := ((dx4*dy2)-(dx2*dy4))*y3 - ((dx3*dy2)-(dx2*dy3))*y4

		// sg := (((dx4 * dy3) - (dx3 * dy4)) * (dy3*cl - y2*dy3*al)) - (((dx2*dy3)-(dx3*dy2))*(dy3*bl-x2*dy3*al))/
		// 	(((dx4*dy3)-(dx3*dy4))*(y2*dy3*bl-x2*dy3*cl)) - (((dx2 * dy3) - (dx3 * dy2)) * (y4*dy3*bl - x4*dy3*cl))
		// sh := -((al + bl*sg) / cl)
		// sa := (((dy3*cl - y2*dy3*al) - (y2*dy3*bl - x2*dy3*cl)) * sg) / (((dx2 * dx3) - (dx3 * dx2)) * cl)
		// sd := ((((dx4 * dy3) - (dx3 * dy4)) * dy2) / (((dx4 * dy2) - (dx2 * dy4)) * dy3)) * sa
		// sb := -(dx3 / dy3) * sa
		// se := -(dx2 / dy2) * sd
		// sc := -(sa * x1) - (sb * y1)
		// sf := -(sd * x1) - (se * y1)

		// op.Uniforms = map[string]any{
		// 	"A": sa,
		// 	"B": sb,
		// 	"C": sc,
		// 	"D": sd,
		// 	"E": se,
		// 	"F": sf,
		// 	"G": sg,
		// 	"H": sh,
		// }
		// op.Images[0] = breadImage
		// w, h := breadImage.Size()
		// screen.DrawRectShader(
		// 	w, h,

		// 	perspectiveShader,
		// 	op,
		// )
		op := &ebiten.DrawImageOptions{}
		x2 := (call.X2/aspectRatio + 1) / 2 * 1280
		y2 := (-call.Y2 + 1) / 2 * 720
		w, h := breadImage.Size()
		x4 := (call.X4/aspectRatio + 1) / 2 * 1280
		y4 := (-call.Y4 + 1) / 2 * 720
		sx, sy := (x4-x2)/float64(w), (y4-y2)/float64(h)
		op.GeoM.Scale(sx, sy)
		op.GeoM.Translate(x2, y2)
		screen.DrawImage(breadImage, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

func init() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.InfoLevel)
}
func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Sonolus Emulator")

	log.Info("Loading processor")
	processorInstance := processor.Init(processor.ProcessorConfig{
		AspectRatio: aspectRatio,
	})
	err := processorInstance.Load("https://servers.sonolus.com/performance-test", "sequential")

	perspectiveShader, err = ebiten.NewShader([]byte(perspective))
	if err != nil {
		panic(err)
	}
	breadImage, _, err = ebitenutil.NewImageFromFile("bread.png")

	if err != nil {
		panic(err)
	}

	processorInstance.Prepare()

	if err := ebiten.RunGame(&Game{
		processor:    &processorInstance,
		touchCounter: 0,
		touches:      []processor.TouchInfo{},
	}); err != nil {
		panic(err)
	}
}
