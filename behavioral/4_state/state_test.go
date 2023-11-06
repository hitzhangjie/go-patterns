package state_test

import (
	"fmt"
	"state"
	"testing"
)

// TestState shows the basic idea: a object has different states, its behavior
// changes based on its state.
//
// Here we changes `VendingMachine` state externally by explicitly call
// `vendingMachine.SetState(state)`.
//
// Actually in State Pattern, different states could know each other, and they
// can transit from one state to another state implicitly. So we can:
//
//   - method 1: when the `noSelectionState.SelectItem()` called,
//     in the SelectItem body end, call `vendingMachine.SetState(hasSelectionState)`
//     to change state (if passing *vendingMachine to it)
//   - method 2: when the `(*vendingMachine).state.SelectItem()` succeeds, call
//     `(*vendingMachine).SetState(hasSelectionState)` at `(*VendingMachine).SelectItem()`
//     body end to change state.
//
// Well, I prefer method2, because this way different State objects don't have
// to know each other. And we could change configuration to dynamically change
// the state transitions, which is much more flexible.
func TestState(t *testing.T) {
	vendingMachine := &state.VendingMachine{}
	noSelectionState := &state.NoSelectionState{}
	hasSelectionState := &state.HasSelectionState{}
	soldState := &state.SoldState{}

	fmt.Println("------------ CurrentState: NoSelection ---------------------")
	vendingMachine.SetState(noSelectionState)
	vendingMachine.InsertCoin()   // Output: Coin inserted.
	vendingMachine.SelectItem()   // Output: Please select an item.
	vendingMachine.DispenseItem() // Output: No item selected.

	fmt.Println("------------ CurrentState: HasSelection --------------------")
	vendingMachine.SetState(hasSelectionState)
	vendingMachine.InsertCoin()   // Output: Coin inserted.
	vendingMachine.SelectItem()   // Output: Item already selected.
	vendingMachine.DispenseItem() // Output: Item dispensed.

	fmt.Println("------------ CurrentState: HasSoldState --------------------")
	vendingMachine.SetState(soldState)
	vendingMachine.InsertCoin()   // Output: Please wait, item being dispensed.
	vendingMachine.SelectItem()   // Output: Item already dispensed.
	vendingMachine.DispenseItem() // Output: Item already dispensed.
}
