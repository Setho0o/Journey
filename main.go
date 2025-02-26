package main

import (
	"image"
	"log"
	"os"

	f "github.com/Setho0o/Journey/Finders"
	gen "github.com/Setho0o/Journey/Generators"
	a "github.com/Setho0o/Journey/assets"
	u "github.com/Setho0o/Journey/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	s    *ebiten.Image
	rgba *image.RGBA // Not sure what to name this but we draw the pixels onto this before we write it to the screen
	gen  gen.Generator
	f    f.Finder
	m    Matrix
	a    map[a.Id]*ebiten.Image
}

func (g *Game) Update() error {
	Keys()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.s = screen

	g.DrawMatrix()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return u.Lx, u.Ly
}

func main() {
	ebiten.SetWindowSize(u.Wx, u.Wy)
	ebiten.SetWindowTitle("~Journey~")
	g := GameInit()

  g.FillLand()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func GameInit() *Game {
	return &Game{
		m:    MatrixInit(),
		a:    a.AssetInit(),
		rgba: image.NewRGBA(image.Rect(0, 0, u.Lx, u.Ly)),
	}
}

func Keys() {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}
}
