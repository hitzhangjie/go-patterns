package bridge_test

import (
	"bridge"
	"testing"
)

func TestBridge(t *testing.T) {
	// red square
	var redSquare = &bridge.Square{}
	redSquare.Color = &bridge.Red{}
	redSquare.Edge = 100
	println(redSquare.String())

	// blue circle
	var blueCircle = &bridge.Circle{}
	blueCircle.Color = &bridge.Blue{}
	blueCircle.Radius = 50
	println(blueCircle.String())
}
