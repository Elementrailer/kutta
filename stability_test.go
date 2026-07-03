package main

import (
	"math"
	"testing"

	"airfoil/lbm"
	"airfoil/viz"
)

// simGame builds the minimal Game the instability path touches: solver, smoke
// and the interactive foil. No Ebiten images, so it runs headless.
func simGame() *Game {
	g := &Game{
		alphaDeg: 4,
		u0:       defaultU,
		nacaCode: profiles[0],
	}
	g.sim = lbm.New(gridW, gridH, tau, g.u0)
	g.smoke = viz.NewParticles(nParticles, gridW, gridH, 1)
	g.applyBody(true)
	return g
}

// TestStepSimResetsAfterInstability pins the backstop for issue #1: if the
// solver somehow goes non-finite, the next stepSim must rebuild a finite flow
// and clear every polluted readout instead of feeding NaN to the renderer.
func TestStepSimResetsAfterInstability(t *testing.T) {
	g := simGame()
	g.fyEMA = 123
	g.clCur = 9e99
	// Poison the solver through its public API: a NaN inlet speed writes NaN
	// populations at the boundaries, and the next collide spreads them into the
	// fields. (Corrupting Rho directly would be laundered — collide recomputes
	// the macroscopic fields from the populations every step.)
	g.sim.SetInletSpeed(math.NaN())
	g.stepSim(2)
	if !g.sim.Finite() {
		t.Fatal("solver still non-finite after the instability reset")
	}
	if g.simErr == "" {
		t.Fatal("simErr not set after the instability reset")
	}
	if g.fyEMA != 0 || g.clCur != 0 {
		t.Fatal("polluted readouts were not cleared by the reset")
	}
	// A user change acknowledges the note.
	g.setSpeed(0.05)
	if g.simErr != "" {
		t.Fatal("simErr not cleared by a user change")
	}
}

// TestInstabilityResetKeepsSceneBody guards the scene-mode gap: the reset must
// rebuild the flow around the LOADED SCENE's mask, not the interactive foil's.
func TestInstabilityResetKeepsSceneBody(t *testing.T) {
	g := simGame()
	g.scn = nekoScene()
	g.sim.SetSolid(g.sceneMask(0))

	g.sim.SetInletSpeed(math.NaN())
	g.stepSim(2)

	if !g.sim.Finite() {
		t.Fatal("solver still non-finite after the scene-mode reset")
	}
	want := g.sceneMask(0)
	for y := range gridH {
		for x := range gridW {
			if g.sim.Solid(x, y) != want[y*gridW+x] {
				t.Fatalf("solid mask diverged from the scene at (%d,%d): reset used the wrong body", x, y)
			}
		}
	}
}
