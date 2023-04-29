package main

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sevenc-nanashi/sonolus-emulator/pkg/processor"
	log "github.com/sirupsen/logrus"
)

type Game struct{
  processor *processor.Processor
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
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
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Sonolus Emulator")

  log.Info("Loading processor")
	processor, err := processor.Load("https://servers.sonolus.com/performance-test", "sequential")

	if err != nil {
		panic(err)
	}

  processor.Prepare()

	if err := ebiten.RunGame(&Game{
    processor: &processor,
  }); err != nil {
		panic(err)
	}
}
