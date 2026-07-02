# airfoil — 2D wind tunnel

A small, pretty 2D wind-tunnel toy for aeromodelers and hobbyists. It streams a
flow past an airfoil and shows the speed field, the vorticity field, smoke
streaklines, and the lift/drag vectors, with an adjustable angle of attack.

It is **qualitative, not validated CFD**: lattice units stand in for physical
ones, so it captures the *shape* of the flow — stagnation point, wake, suction
over the top, separation as the angle of attack grows — rather than calibrated
numbers. Good for intuition and for looking nice; not for sizing a real wing.

## Run

```sh
go run .
```

| Key        | Action                          |
|------------|---------------------------------|
| ↑ / ↓      | angle of attack ±1°             |
| Tab        | cycle NACA profile              |
| V          | toggle speed / vorticity field  |
| `[` / `]`  | inlet speed ∓ / ±               |
| Space      | pause                           |
| N          | single step (while paused)      |
| R          | reset the flow                  |

## How it works

- **Solver** (`internal/lbm`): a 2D Lattice-Boltzmann method (D2Q9, BGK). An
  open channel with a free stream from the left; the body is a no-slip wall via
  half-way bounce-back. Forces come from a pressure integral over the body
  faces (pressure/form drag and lift; skin friction is neglected).
- **Geometry** (`internal/foil`): NACA 4-digit airfoils from their closed-form
  formula — a whole "database" with nothing stored — placed, rotated by angle of
  attack, and rasterized into the solver's solid mask (point-in-polygon).
- **Visualization** (`internal/viz`): color maps for the scalar fields and a
  smoke-tracer particle system advected by the velocity field.
- **App** (`main.go`, `game.go`): the Ebitengine render loop tying it together.

The `lbm`, `foil` and `viz` packages have no rendering dependency and are unit
tested headlessly. `cmd/snapshot` renders the fields to PNG without a GPU, which
is how the physics is sanity-checked (lift rising with angle of attack, the drag
bucket, correct force signs).

## Status

Initial spike. Next: a draw-your-own-profile editor (reusing the vector editor
kit and bloom from the `linefire` project) and richer on-screen controls.
