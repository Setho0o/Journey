package main

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"

	f "github.com/Setho0o/Journey/Finders"
	gen "github.com/Setho0o/Journey/Generators"
	a "github.com/Setho0o/Journey/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Lx       = Wx // useable pixels
	Ly       = Wy
	Wx       = 1920 // window sizw
	Wy       = 1080
	TileSize = 16
)

type Game struct {
	s   *ebiten.Image
	gen gen.Generator
	f   f.Finder
	m   Matrix
  a  map[a.Id]*image.RGBA
}

func (g *Game) Update() error {
	Keys()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.s = screen

	for _, e := range g.m {
		for _, e := range e {
			g.s.Set(e.x, e.y, color.White)
		}
	}
  
  rgba := image.NewRGBA(image.Rect(0,0,Lx,Ly)) 
  draw.Draw(rgba, g.a[a.LandM].Bounds(), g.a[a.LandM],image.Pt(0,0),draw.Src)  
  g.s.WritePixels(rgba.Pix) 
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
    a: a.AssetInit(),
	}
}

func Keys() {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}
}
