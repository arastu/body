package organs

import (
	"fmt"
	"image/color"
	"math/rand"
	"sync"
	"time"
)

// Organ represent body organs
type Organ interface {
	ID() int

	// Color get color of organ
	Color() color.Color

	//Destroy organ
	Destroy() (int, error)
}

var organs = make(map[string]Organ)
var mu sync.Mutex

type organ struct {
	id        int
	color     color.Color
	destroyed bool
}

// Head yes Head
type Head struct {
	*organ
	sync.Mutex
}

func (h *Head) ID() int {
	return h.id
}

// Color get color of head
func (h *Head) Color() color.Color {
	return h.color
}

//Destroy head
func (h *Head) Destroy() (int, error) {
	h.Lock()
	defer h.Unlock()

	if h.destroyed {
		return 0, fmt.Errorf("Head already destroyed!")
	}

	h.destroyed = true

	delete(organs, "Head")

	return randomInt(1, 1000), nil
}

// MakeHead ...
func MakeHead(organColor color.Color) Organ {
	if head, ok := organs["Head"]; ok {
		return head
	}

	mu.Lock()
	defer mu.Unlock()

	head := Head{
		organ: &organ{
			id:        randomInt(1000, 10000),
			color:     organColor,
			destroyed: false,
		},
	}

	organs["Head"] = &head
	return &head
}

// Tongue yes Tongue
type Tongue struct {
	*organ
	sync.Mutex
}

func (t *Tongue) ID() int {
	return t.id
}

// Color get color of head
func (t *Tongue) Color() color.Color {
	return t.color
}

//Destroy head
func (t *Tongue) Destroy() (int, error) {
	t.Lock()
	defer t.Unlock()

	if t.destroyed {
		return 0, fmt.Errorf("Tongue already destroyed!")
	}

	t.destroyed = true

	delete(organs, "Tongue")

	return randomInt(1, 1000), nil
}

// MakeTongue ...
func MakeTongue(organColor color.Color) Organ {
	if tongue, ok := organs["Tongue"]; ok {
		return tongue
	}

	mu.Lock()
	defer mu.Unlock()

	tongue := Tongue{
		organ: &organ{
			id:        randomInt(1000, 10000),
			color:     organColor,
			destroyed: false,
		},
	}

	organs["Tongue"] = &tongue
	return &tongue
}

func randomUint8(min, max int) uint8 {
	return uint8(randomInt(min, max))
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
