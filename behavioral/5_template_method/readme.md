Context: 

    When designing an algorithm as a series of steps, some steps may be common
    across subclasses while others may vary.

Solution: 

    Define an abstract base class that specifies a skeleton template method
    defining the algorithm steps. Some steps are implemented in the base class
    while others are abstract methods to be implemented in subclasses.

Variants:

    - Template method in base class calls abstract methods defined in that base
    class. Subclasses implement these methods.

    - Template method calls methods defined in subclasses directly. Subclass can
    override these methods to provide custom implementation.

    - Can define hooks that subclasses can override to customize template method
    behavior without modifying its structure.

Relations:

    - Allows defining the overall algorithm structure in the base class while
    allowing subclasses to customize steps.

    - Implements the Hollywood Principle - "Don't call us, we'll call you". 
    Base class template method calls subclass methods.

    - Allows inversion of control between base class and subclasses.

    - Strategy pattern relies on composition - Template method relies on
    inheritance. (ps: but in Go, prefer composition over interitance)

In summary, Template Method provides a skeleton algorithm structure in a base
class, deferring some steps to subclasses. This allows subclasses to customize
the algorithm without changing its overall structure.