Context:

    The Memento pattern is used when we need to capture and restore the internal
    state of an object without violating encapsulation. Sometimes, we need to
    save the state of an object at a particular point in time and be able to
    restore it later. However, exposing the internal state directly can lead to
    encapsulation issues. The Memento pattern provides a solution to this
    problem.
    
Solution:

    The Memento pattern involves three main components: 
    - the Originator, 
    - the Memento, 
    - and the Caretaker. 

    The Originator is the object whose state we want to save and restore. The
    Memento is an object that represents the saved state of the Originator. The
    Caretaker is responsible for managing the Memento objects and interacting
    with the Originator.
    
    The Originator object creates a Memento object that captures its internal
    state. The Memento object is typically immutable and provides methods for
    the Originator to retrieve the state it represents. The Originator can also
    provide methods for setting its state based on a Memento object.

    The Caretaker object requests Mementos from the Originator to save its state
    at a particular point in time. It stores these Mementos and can provide them
    back to the Originator for state restoration when needed.

    This pattern allows us to save and restore the state of an object without
    exposing its internal details. It promotes encapsulation by keeping the
    state information separate from the object itself.

Variants:

    The Memento pattern can be extended with variants such as multiple Mementos,
    wide Memento, or using a Caretaker with additional functionality. These
    variants provide flexibility in managing and restoring object states based
    on specific requirements.
    
Relation:

    The Memento pattern can be related to other patterns in the following ways:

    - Command Pattern: The Memento pattern can be used in conjunction with the
    Command pattern to implement undo and redo functionality. The Command
    pattern encapsulates operations as objects, and the Memento pattern can be
    used to save and restore the state of the objects involved in the
    operations.
    
    - Prototype Pattern: The Memento pattern can be used in conjunction with the
    Prototype pattern to create deep copies of objects. The Memento can capture
    the state of an object, and the Prototype pattern can be used to create a
    new object with the same state.
    
    - State Pattern: The Memento pattern can be used to implement the internal
    state management of objects in the State pattern. The Memento captures the
    state of the object, and the State pattern defines the behavior based on
    that state.
    
In summary, the Memento pattern provides a way to capture and restore the
internal state of an object without violating encapsulation. It involves the
Originator, Memento, and Caretaker components. The Caretaker manages the Memento
objects and interacts with the Originator to save and restore states. The
pattern can be extended with variants and can be related to other patterns such
as Command, Prototype, and State.