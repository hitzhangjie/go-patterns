Context: 
    You have an abstraction with multiple implementations. The abstraction and 
    implementations are tightly coupled together through inheritance 
    hierarchies. This leads to issues with rigidity, difficulty extending the 
    abstraction or implementations independently, and being unable to swap 
    implementations at runtime.

    One typical example is to add Color into existing Shape hierarchy, it could
    be depicted as following:

                --------------
                |    Shape   |
                --------------
                  ^         ^
                 /           \
        ------------         ------------
        |  Circle  |         |  Square  | <-- add color to Shape hierarchy
        ------------         ------------
           ^      ^            ^        ^
           |      |            |        | 
    ----------- ------------ ----------- ------------
    |CircleRed| |CircleBlue| |SquareRed| |SquareBlue| <-- subclasses explode!
    ----------- ------------ ----------- ------------

    There's the Shape abstraction, which is extended or implemented by Circle
    and Square. Well, now we want to render them in different colors, like red
    and blue. One solution is adjusting the Shape hierarchy to add more 
    subclasses, Circle{Red,Blue}, Square{Red,Blue}.

    If the problem scales, like adding kinds of materials, the number of 
    subclasses will explode.

Solution:

    The Bridge pattern decouples the abstraction from the implementations by 
    extracting them into two separate class hierarchies. This allows the 
    abstraction and implementations to vary and evolve independently of each 
    other. 

    The abstraction contains a reference to the implementation rather than 
    inheriting from it. The abstraction and implementation define interfaces 
    that dictate how they interact, and concrete classes implement these 
    interfaces.

    ps: For this example, that is extracting the color into separate hierarchy, 
    and Shape abstraction uses a reference to Color abstraction to use the
    different Color implementions. And the Shape and the Color hierarchies 
    could develop separately...

    This enables:
    - Abstraction and implementations can be extended easily without affecting 
      each other
    - New abstractions and implementations can be added independently 
    - Changes to one hierarchy don't impact the other
    - Runtime switching between different implementations

    In summary, Bridge aims to decouple a tightly coupled abstraction-implementation 
    hierarchy into two separate hierarchies that can evolve independently. 
    This provides flexibility, better extensibility, and reusability.

Variants:


Relation:
