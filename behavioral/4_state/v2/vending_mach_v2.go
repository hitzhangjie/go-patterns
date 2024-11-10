package state

// In this demo, the vending machine changes its state by externally calling
// vendingMachine.SetState(state). I believe that when vendingMachine.SelectItem
// succeeds, it should automatically transition to the next state. Now, the
// question is, how can we determine the next state? Should we hardcode it in
// the return value of func SelectItem() State, or should we read it from a
// configuration file?
//
// State Pattern doesn't require that where the transition controlled:
// 1. we can control it via the Context.SetState(next State) method;
// 2. or we can control it in the concrete State methods.
//
// The 1st method, each concrete State don't know about the next state. It's
// the responsibility of the Context or the caller to determine which State
// to transition to.
//
// The 2nd method, each concrete State knows about its next State. This results
// in some strong coupling between States and their transitions.
//
// Well, I think that the 1st method is more common. It's easier for a caller.
// But if we want to make it more generic, we can use table driven design.
// Or, we can use a State Machine Pattern. If so, we can understand the whole
// more clearly.
//
// see also:
// - FSM framework, https://github.com/qmuntal/stateless
// - FSM framework, https://github.com/looplab/fsm
//
// So, you see, the State Pattern is not a must for every situation. And It's
// not a silver bullet. But it's a good starting point to understand how to use
// the pattern to solve problems.
//
// Business complexity is always hard to predict, complexity may means chances,
// so we should welcome it and find out how to reduce the complexity of solution.
