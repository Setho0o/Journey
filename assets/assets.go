package assets

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
)

/*
  planning on just getting the pixel data for the assests
  and maping them to some id for preformance
*/
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

/*
  could be worth it to just map the uint8 pixel instead of the rgba struct 
  but for now we will keep it in case we need it 
*/

func AssetInit() map[Id]*image.RGBA {
  return map[Id]*image.RGBA {
    LandM : DecodePng("assets/land.png"), // gonna need to organize and name the files better this was just for a test 
  }
}


func DecodePng(name string) *image.RGBA {
  file, err := os.Open(name)
  if err != nil {
    log.Fatal(err, "failed to open file")
  }
  defer file.Close()

  img, err := png.Decode(file)
  if err != nil {
    log.Fatal(err, "failed to decode png")
  }
  // image.Image doesnt give us the pixels so we have to draw it to another struct 
  rect := img.Bounds()
  rgba := image.NewRGBA(rect)
  draw.Draw(rgba,rect,img,rect.Min,draw.Src)
  return rgba
}


