package main

import (
	"log"
	"os"

	f "github.com/Setho0o/Journey/Finders"
	gen "github.com/Setho0o/Journey/Generators"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
  Lx = Wx / 4 // useable pixels 
  Ly = Wy / 4
  Wx = 1920 // window sizw
  Wy = 1080
)


type Game struct{
  s *ebiten.Image
  gen gen.Generator
  f f.Finder
}


func (g *Game) Update() error {
  Keys()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  g.s = screen
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Lx, Ly
}

func main() {
	ebiten.SetWindowSize(Wx, Wy)
	ebiten.SetWindowTitle("~Journey~")
  g := GameInit()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func GameInit() *Game {
  return &Game{
  }
}

func Keys() {
  if ebiten.IsKeyPressed(ebiten.KeyQ) {
    os.Exit(0)
  }
} 
