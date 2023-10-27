package builder_test

import (
	"fmt"
	"testing"

	"builder"
)

func Test_Creational_Builder(t *testing.T) {
	bd := builder.NewBuilder()
	bd.BuildFloor([]builder.Pos{{0, 0, 0}, {0, 10, 0}, {10, 10, 0}, {10, 0, 0}}, builder.MaterialStyle(0))
	bd.BuildWall(builder.Pos{0, 0, 0}, builder.ColorStyle(0)).
		BuildDoor(builder.Pos{0, 1, 0}, 2, 1, builder.MaterialStyle(0), builder.ColorStyle(0)).
		BuildWindow(builder.Pos{0, 2, 0}, 1, 1, builder.GlassStyle(0))
	//builder another 3 walls with doors and windows
	//bd.BuildWall(...).
	//  BuildDoor(...).
	//  Buildwindow(...)
	bd.BuildRooves(builder.RoofStyle(0))
	house := bd.Build()
	fmt.Printf("The house is built: %+v", house)
}

func Test_Creational_BuilderWithOrderer(t *testing.T) {
	// actually we should always limit the building order: floor -> wall -> rooves
	// we must build the house to avoid the house collapse.
	//
	// so we can import an orderer which limits the building order.
	ord := builder.NewOrderer(builder.NewBuilder())
	//ord.WantFloor(...)
	//ord.WantWall(...)
	//ord.WantWindow(...)
	//ord.WantDoor(...)
	house := ord.Build()
	fmt.Printf("The house is built: %+v", house)
}
