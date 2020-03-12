package ants

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/slarsar/ulam-spiral/prime"
	"go.uber.org/zap"
)

type pixel struct {
	x, y int
}

type ant struct {
	log  *zap.SugaredLogger
	loc  pixel
	dest int
	spot int
}

// des is a positive number
func NewAnt(log *zap.SugaredLogger, des int) ant {
	return ant{loc: pixel{0, 0}, spot: 1, dest: des, log: log}

}

func abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

func (a *ant) Next() (err error) {
	if abs(a.loc.x) > a.dest || abs(a.loc.y) > a.dest {
		fmt.Println("ants' have finished his job at ",a.spot)
		return errors.New("stop")
	}

	defer func() {
		a.spot++
	}()
	if a.spot == 1 {
		a.loc.y++
		return
	}

	x := a.loc.x
	y := a.loc.y

	if a.loc.x >= 0 {

		if x == y {
			//a.loc.x--
			a.loc.y++
			return
		}
		if x == -y {
			a.loc.y++
			return
		}

		if x > abs(y) {
			a.loc.y++
			return
		}
		if y > 0 {
			a.loc.x--
			return
		}
		a.loc.x++
		return
	} else { // x < 0
		if x == y {
			a.loc.x++
			return
		}

		if x == -y {
			a.loc.y--
			return
		}
		if abs(x) > abs(y) {
			a.loc.y--
			return
		}
		if y < 0 {

			a.loc.x++
		} else {
			a.loc.x--
		}
		return
	}

}

func (a *ant) Look() (int, int, int, bool) {
	return a.loc.x, a.loc.y, a.spot,prime.CheckPrime(a.spot)
}
