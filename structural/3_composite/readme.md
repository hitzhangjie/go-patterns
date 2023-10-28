Context: 

    We need to represent hierarchical tree structures of objects where 
    individual objects and compositions of objects should be treated uniformly. 
    For example, representing files and folders in a file system.

Solution: 

    The composite pattern describes that a group of objects are to be treated 
    in the same way as a single instance of an object. The key objects are:

    - Component: Declares the interface for objects in the composition and 
      implements default behavior for the interface common to all classes.
    - Leaf: Represents leaf objects in the composition. A leaf has no children.
    - Composite: Defines behavior for components having children and store 
      child components. Implements child related operations in the component 
      interface.
    - Client: Manipulates the objects in the composition through the component 
      interface.

Variants:

    - Safe Composite: Methods in the component interface can allow handling of 
      leaf nodes differently from composites when required. For example file 
      operations may not be applicable to folders.
    - Sharing: Composite pattern allows sharing of components, which can lead 
      to aliasing issues. Variants may restrict sharing of components.

Relations:

    - Decorator: Both patterns abstract an object from the client. But 
      decorator adds additional responsibilities dynamically, composite 
      structures objects recursively.
    - Iterator: Can be used to traverse composites.
    - Visitor: Can be used to execute an operation over a composite.

Appendix:

    The meaning of term 'composition' in Design Pattern area differs from
    'composition over inheritance' in Object-Oriented programming principle.

    - Composition in the Composite design pattern refers specifically to the 
      pattern's use of composite objects that are composed of other objects in 
      a tree structure. The composite objects delegate work to their child 
      objects to implement the common interface.
    - "Composition over inheritance" is a general principle for object-oriented 
      design. It means favoring object composition (having objects contain 
      other objects and invoking their methods) over class inheritance 
      (extending classes by subclassing) where possible.
