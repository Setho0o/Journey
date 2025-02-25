package assets

import (
	"image/png"
	"log"
	"os"

	u "github.com/Setho0o/Journey/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

/*
We want to be consistant with our assets so im laying out some rules,
there must be three types of each asset for randomness, and
only three colors for each asset. Lastly just write
the colors your using incase we want to redo some of them.
*/

type Id int

const (
	Blank Id = iota
	LandM
	LandTL
	LandTR
	LandBL
	LandBR

	GrassA

	WaterM1
	WaterM2
	WaterM3

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

var (
	WaterHex = []string{
		"00bfff",
		"87ceeb",
		"ffffff",
	}
)

func AssetInit() map[Id]*ebiten.Image {
	return map[Id]*ebiten.Image{
		Blank:   ebiten.NewImage(u.TileSize, u.TileSize),
		LandM:   DecodePng("assets/land.png"), // gonna need to organize and name the files better this was just for a test
		GrassA:  DecodePng("assets/grassA.png"),
		WaterM1: DecodePng("assets/Water/WaterM1.png"),
		WaterM2: DecodePng("assets/Water/WaterM2.png"),
		WaterM3: DecodePng("assets/Water/WaterM3.png"),
		BushA:   DecodePng("assets/BushA.png"),
		BushB:   DecodePng("assets/BushB.png"),
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
