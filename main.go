package main

import (
	"fmt"
	"github.com/slarsar/ulam-spiral/ants"
	"github.com/slarsar/ulam-spiral/prime"
	"go.uber.org/zap"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
)

var log *zap.SugaredLogger

func initLog() *zap.SugaredLogger {
	var err error
	cfg := zap.NewDevelopmentConfig()
	cfg.Level.SetLevel(zap.ErrorLevel)
	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return l.Sugar()
}

func main() {
	var err error
	if len(os.Args) != 2 {
		fmt.Println("usage: ./ulam-spiral [number]")
		return
	}

	top2, _ := strconv.ParseInt(os.Args[1], 10, 64)
	top := int(top2)

	filename := fmt.Sprintf("ulam-spiral-%s.png", os.Args[1])
	var file *os.File
	if file, err = os.Create(filename); err != nil {
		log.Panic(err)
	}

	log = initLog()

	fmt.Println("generating ...")

	prime.InitPmap(4 * top * top)

	img := image.NewRGBA(image.Rect(0, 0, 2*top, 2*top))

	ant := ants.NewAnt(log, top)

	for ant.Next() == nil {
		x, y, now, b := ant.Look()
		log.Infof("ant now in %d %d:%d %+v", x, y, now, b)

		x1, y1 := x+top, y+top
		c := color.Black
		if b {
			c = color.White
		}
		img.Set(x1, y1, c)
	}

	if err := png.Encode(file, img); err != nil {
		log.Panic(err)
		return
	}
	fmt.Println(filename, " have generated !")
}
