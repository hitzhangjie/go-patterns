// Package state provides a demo to show how State Pattern works.
//
// It defines `VendingMachine` struct that represents the context
// object. It maintains a reference to the current state and delegates the
// behavior to that state.
//
// The `State` interface defines the methods `InsertCoin()`, `SelectItem()`, and
// `DispenseItem()`, which represent the different behaviors of the vending
// machine in each state.
//
// We have three concrete state implementations:
// - `NoSelectionState`,
// - `HasSelectionState`,
// - `SoldState`.
//
// Each state implements the `State` interface and provides its own behavior for
// the methods.
//
// In the `TestState` function, we create an instance of the `VendingMachine`
// and set the initial state to `NoSelectionState`. We then call the methods on
// the vending machine to simulate different actions, such as inserting a coin,
// selecting an item, and dispensing the item. The output demonstrates how the
// behavior of the vending machine changes based on its current state.
//
// This example showcases how the State pattern allows the vending machine to
// have different behaviors depending on its internal state. It provides a
// flexible and maintainable solution by encapsulating the behavior of each
// state in separate classes, making it easier to add new states or modify
// existing ones without modifying the existing code.
package state

import "fmt"

// State represents the interface for different states of the vending machine
type State interface {
	InsertCoin()
	SelectItem()
	DispenseItem()
}

// VendingMachine represents the context object that maintains the current state
type VendingMachine struct {
	currentState State
}

func (v *VendingMachine) SetState(state State) {
	v.currentState = state
}

func (v *VendingMachine) InsertCoin() {
	v.currentState.InsertCoin()
}

func (v *VendingMachine) SelectItem() {
	v.currentState.SelectItem()
}

func (v *VendingMachine) DispenseItem() {
	v.currentState.DispenseItem()
}

// NoSelectionState represents the state when no item is selected
type NoSelectionState struct{}

func (n *NoSelectionState) InsertCoin() {
	fmt.Println("✅ Coin inserted.")
}

func (n *NoSelectionState) SelectItem() {
	fmt.Println("✅ Please select an item.")
}

func (n *NoSelectionState) DispenseItem() {
	fmt.Println("❌ No item selected.")
}

// HasSelectionState represents the state when an item is selected
type HasSelectionState struct{}

func (h *HasSelectionState) InsertCoin() {
	fmt.Println("❌ Coin already inserted.")
}

func (h *HasSelectionState) SelectItem() {
	fmt.Println("❌ Item already selected.")
}

func (h *HasSelectionState) DispenseItem() {
	fmt.Println("✅ Item dispensed.")
}

// SoldState represents the state when an item is sold
type SoldState struct{}

func (s *SoldState) InsertCoin() {
	fmt.Println("❌ Please wait, item being dispensed.")
}

func (s *SoldState) SelectItem() {
	fmt.Println("❌ Item already selected and dispensed.")
}

func (s *SoldState) DispenseItem() {
	fmt.Println("❌ Item already dispensed.")
}
