Context:

    The strategy pattern is used when we have multiple algorithms/strategies
    that can be used interchangeably for a specific task. The client using the
    algorithms does not need to know about the implementation details of the
    algorithms.

Solution:

    Define an interface that will be used by all the different strategies.
    Create multiple strategy classes that implement this common interface. The
    context object will receive a strategy object and delegate the work to this
    object, rather than doing the work itself.

Variants:

    - Pass the strategy to the context as constructor parameter. The context will
    store reference to the passed in strategy.
    - Allow the client to change the associated strategy dynamically at runtime.
    The context can define a setter method to allow changing strategies.

Relations:

    - Strategy pattern promotes the Open/Closed principle. New strategies can be
    added without changing existing client code.

    - Strategies encapsulate the algorithm implementation. The implementation
    can change without affecting the Context using it.

    - Strategy is similar to State pattern, but State allows changing algorithm
    dynamically based on state, while Strategy client must explicitly choose the
    strategy. And once the strategy is choosen, it can't be changed.

In summary, the strategy pattern encapsulates interchangeable algorithms and
decouples the client from implementation details of strategies being used. It
allows switching strategies at runtime to vary the behavior.
