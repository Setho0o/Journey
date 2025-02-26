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
	LandM1
	LandM2
	LandM3

	GrassA

	WaterM1
	WaterM2
	WaterM3

	WaterTL1
	WaterTL2
	WaterTL3
	WaterTR1
	WaterTR2
	WaterTR3
	WaterBL1
	WaterBL2
	WaterBL3
	WaterBR1
	WaterBR2
	WaterBR3

	WaterT1
	WaterT2
	WaterT3
	WaterB1
	WaterB2
	WaterB3
	WaterL1
	WaterL2
	WaterL3
	WaterR1
	WaterR2
	WaterR3

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

  LandHex = []string{
		"30a430",
    "26ab26",
	}
)

func AssetInit() map[Id]*ebiten.Image {
	return map[Id]*ebiten.Image{
		Blank:   ebiten.NewImage(u.TileSize, u.TileSize),
		
    LandM1:   DecodePng("assets/Land/LandM1.png"), 
		LandM2:  DecodePng("assets/Land/LandM2.png"),
		LandM3:  DecodePng("assets/Land/LandM3.png"),

		WaterM1: DecodePng("assets/Water/WaterM1.png"),
		WaterM2: DecodePng("assets/Water/WaterM2.png"),
		WaterM3: DecodePng("assets/Water/WaterM3.png"),

		WaterTL1: DecodePng("assets/Water/WaterTL1.png"),
		WaterTL2: DecodePng("assets/Water/WaterTL2.png"),
		WaterTL3: DecodePng("assets/Water/WaterTL3.png"),

		WaterTR1: DecodePng("assets/Water/WaterTR1.png"),
		WaterTR2: DecodePng("assets/Water/WaterTR2.png"),
		WaterTR3: DecodePng("assets/Water/WaterTR3.png"),

		WaterBR1: DecodePng("assets/Water/WaterBR1.png"),
		WaterBR2: DecodePng("assets/Water/WaterBR2.png"),
		WaterBR3: DecodePng("assets/Water/WaterBR3.png"),

		WaterBL1: DecodePng("assets/Water/WaterBL1.png"),
		WaterBL2: DecodePng("assets/Water/WaterBL2.png"),
		WaterBL3: DecodePng("assets/Water/WaterBL3.png"),

		WaterT1: DecodePng("assets/Water/WaterT1.png"),
		WaterT2: DecodePng("assets/Water/WaterT2.png"),
		WaterT3: DecodePng("assets/Water/WaterT3.png"),

		WaterB1: DecodePng("assets/Water/WaterB1.png"),
		WaterB2: DecodePng("assets/Water/WaterB2.png"),
		WaterB3: DecodePng("assets/Water/WaterB3.png"),

		WaterL1: DecodePng("assets/Water/WaterL1.png"),
		WaterL2: DecodePng("assets/Water/WaterL2.png"),
		WaterL3: DecodePng("assets/Water/WaterL3.png"),

		WaterR1: DecodePng("assets/Water/WaterR1.png"),
		WaterR2: DecodePng("assets/Water/WaterR2.png"),
		WaterR3: DecodePng("assets/Water/WaterR3.png"),

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
