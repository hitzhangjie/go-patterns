package state

// In this demo, the vending machine changes its state by externally calling
// vendingMachine.SetState(state). I believe that when vendingMachine.SelectItem
// succeeds, it should automatically transition to the next state. Now, the
// question is, how can we determine the next state? Should we hardcode it in
// the return value of func SelectItem() State, or should we read it from a
// configuration file?
