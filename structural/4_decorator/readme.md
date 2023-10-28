Context: 

    You want to augment or alter the functionality of objects at runtime.

    To extend the base class is a solution, you could augment or alter the 
    functionality of base class. But inheritance has some serious ceveats,
    that you must know:

    - Inheritance is static. You can't alter the behavior of an existing
      object at runtime. You can only replace the whole subclass object with
      another one that's created from a different subclass.
    - Subclasses can have just one parent class. In most languages, inheritance
      doesn't let a class inherit behaviors of multiple classes at the same
      time. (C++ supports multiple inheritance, while most languages don't)
    
    So, firstly we need implement the goal to argument or alter the 
    functionality, then we must take the inheritance ceveats into 
    consideration.

    ps: In previous patterns, we described the Adapter Pattern, which wraps an 
    reference to the implemention object, we can do something before calling
    the referenced object method. Well, we still could use this `wrapper` way
    to solve this problem.

Solution: 

    If you want to augment or alter the functionality of objects at runtime,
    you can wrap them inside decorator objects. The wrapped object and 
    decorator should have a common interface so they are interchangeable.

    Define a decorator base class that implements the common interface. 
    
    And each concrete decorator extend this and add extra functionality in 
    its methods. The decorator keeps a reference to the wrapped object and 
    delegates to it after adding its own behavior.

    ps: Decorator Pattern is also known as Wrapper Pattern, I think you should
    realize that Adapter Pattern is also known as Wrapper Pattern. Well, 
    wrapper is just the method to implement, you could differentiate the two
    patterns by its function or goal.

Variants:

    - Multiple decorators can wrap an object to combine functionality. They 
      are applied recursively.
    - Decorators can be abstract classes instead of concrete classes, allowing 
      custom implementations.
    - The base decorator class can be an interface instead of a class.

Relation: 

    The decorator pattern is an alternative to subclassing for extending 
    functionality. It provides more flexibility since new decorators can be 
    created independently without modifying original classes. It avoids 
    exploding subclass combinations when multiple extensions are possible.


Appendix:

    The key aspect is that the decorator conforms to the interface of the 
    original object so it's transparently interchangeable. This allows many 
    permutations of added behavior at runtime without complexity. The pattern 
    extends objects instead of static classes.
