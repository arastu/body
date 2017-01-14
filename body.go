package body

import (
	"image/color"
	"math/rand"
	"time"
)

type Organ interface {
	Color() color.Color
	Destroy() (int, error)
}

type Head struct{}

func (h Head) Color() color.Color {
	return color.RGBA{randomUint8(0, 255), randomUint8(0, 255), randomUint8(0, 255), randomUint8(0, 255)}
}

func (h Head) Destroy() (int, error) {
	return randomInt(0, 1000), nil
}

type Tongue struct{}

func (t Tongue) Color() color.Color {
	return color.RGBA{randomUint8(0, 255), randomUint8(0, 255), randomUint8(0, 255), randomUint8(0, 255)}
}

func GetHead() *Head {
	return &Head{}
}

func GetTongue() *Tongue {
	return &Tongue{}
}

func randomUint8(min, max int) uint8 {
	return uint8(randomInt(min, max))
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
