// Command airfoil is a 2D wind-tunnel toy: it streams a Lattice-Boltzmann flow
// past a NACA airfoil and visualizes speed, vorticity, smoke streaklines and the
// lift/drag vectors, with an adjustable angle of attack.
package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(winW, winH)
	ebiten.SetWindowTitle("airfoil — 2D wind tunnel")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	setWindowIcon()
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
