package assets

import (
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Id int 

const (
  LandM Id = iota
  LandTL 
  LandTR
  LandBL
  LandBR
  
  WaterM
  WaterTL
  WaterTR
  WaterBL
  WaterBR

  TreeA
  TreeB
  TreeC
  BushA
  BushB
  BushC
)

func AssetInit() map[Id]*ebiten.Image{
  return map[Id]*ebiten.Image {
    LandM : DecodePng("assets/land.png"), // gonna need to organize and name the files better this was just for a test 
  }
}


func DecodePng(name string) *ebiten.Image {
  file, err := os.Open(name)
  if err != nil {
    log.Fatal(err, "failed to open file")
  }
  defer file.Close()

  img, err := png.Decode(file)
  if err != nil {
    log.Fatal(err, "failed to decode png")
  }

  return ebiten.NewImageFromImage(img)
}



