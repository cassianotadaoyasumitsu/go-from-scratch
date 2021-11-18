package interfaces

import "go-from-scratch/interfaces"

// Syntax
type Launcher interface {
	launch()
	calculateRoute(lat, lon float64) bool
}

type SpaceLauncher struct {
	title string
	number float64
}

func (e, *SpaceLauncher) Launch() {
	// TODO
	return true
}

func (e, *SpaceLauncher) calculateRoute(lat, lon float64) bool {
	// calculate route
	return true
}

func Attack(l Launcher) {
	l.launch()
}

func interfaces() {
	//somewhere in main func
	sl := SpaceLauncher{title = "Space Bird"}
	ml := MarineLauncer{//...fields}
	if Attack(sl) {
		Attack(ml)
	}
}
