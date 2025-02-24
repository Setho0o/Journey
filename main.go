package main

import (
	"image/color"
	"log"
	"os"

	f "github.com/Setho0o/Journey/Finders"
	gen "github.com/Setho0o/Journey/Generators"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
  Lx = Wx // useable pixels 
  Ly = Wy
  Wx = 1920 // window sizw
  Wy = 1080
  TileSize = 16 
)


type Game struct{
  s *ebiten.Image
  gen gen.Generator
  f f.Finder
  m Matrix
}


func (g *Game) Update() error {
  Keys()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  g.s = screen

  for _,e := range g.m {
    for _,e := range e {
      g.s.Set(e.x,e.y,color.White)
    }
  
  }
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
    m: MatrixInit(),
  }
}

func Keys() {
  if ebiten.IsKeyPressed(ebiten.KeyQ) {
    os.Exit(0)
  }
} 
