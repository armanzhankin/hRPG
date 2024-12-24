package main

import (
	"fmt"
	"github.com/armanzhankin/hrpg/internal/characters"
	"math"
	"math/rand"
)

func main() {
	swordMan := characters.NewSwordMan(math.Abs(rand.Float64()*15) + 15)
	mage := characters.NewMage(math.Abs(rand.Float64()*15) + 15)

	for swordMan.IsAlive() && mage.IsAlive() {
		swordMan.Attack(mage)
		if !mage.IsAlive() {
			fmt.Printf("%s won!\n", swordMan.Name)
			break
		}

		mage.Attack(swordMan)
		if !swordMan.IsAlive() {
			fmt.Printf("%s won!\n", mage.Name)
			break
		}
	}
}
