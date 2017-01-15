package main

import (
	"image/color"

	"fmt"
	"os"
	"text/tabwriter"

	log "github.com/Sirupsen/logrus"
	"github.com/arastu/body/organs"
)

func main() {
	red := color.RGBA{255, 0, 0, 255}
	green := color.RGBA{0, 255, 0, 255}

	const format = "%v\t%v\t%v\t\n"

	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	defer tw.Flush()

	fmt.Fprintf(tw, format, "Type", "ID", "Color")
	fmt.Fprintf(tw, format, "-----", "-----", "-----")

	head := organs.MakeHead(green)
	fmt.Fprintf(tw, format, "Head", head.ID(), head.Color())

	tongue := organs.MakeTongue(red)
	fmt.Fprintf(tw, format, "Tongue", tongue.ID(), tongue.Color())

	if tongue.Color() == red && head.Color() == green {
		headID, err := head.Destroy()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("green head Is Destroyed!, headID is: %d", headID)
	}

}
