module airfoil

go 1.26.4

require (
	github.com/crgimenes/filo v0.0.10
	github.com/crgimenes/glaze v0.0.0-00010101000000-000000000000
	github.com/crgimenes/minigui v0.0.0-00010101000000-000000000000
	github.com/hajimehoshi/ebiten/v2 v2.10.0-alpha.11
	golang.org/x/image v0.38.0
)

require github.com/ebitengine/purego v0.11.0-alpha.1 // indirect

require (
	github.com/crgimenes/native v0.0.0-20260624104951-fb41a0e4f845
	github.com/ebitengine/gomobile v0.0.0-20260211053922-3d992dae95d1 // indirect
	github.com/ebitengine/hideconsole v1.0.0 // indirect
	github.com/go-text/typesetting v0.3.5-0.20260328164731-48df487c1500 // indirect
	github.com/jezek/xgb v1.3.0 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/sys v0.45.0 // indirect
	golang.org/x/text v0.35.0 // indirect
)

replace github.com/crgimenes/native => ../native

replace github.com/crgimenes/minigui => ../minigui

replace github.com/crgimenes/glaze => ../glaze
