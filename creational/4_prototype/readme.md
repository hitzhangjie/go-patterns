Context: 
    we have an object and you want to create a copy of it. how could you do it?
    maybe you want to know the class of the object, so you can create a new 
    object of the same class, then copy all fields' values to the new object.

    but:
    - the class encapsulates some unvisible fields, so you can't copy the 
    fields' values. 
    - or, you can't know the class of the object because you 
    reference it by an interface, so you can't create a new object of the same
    class.
    - or, the object maybe complex, it's not so convenient to copy all fields.

    even if we know the class of the object and could copy all fields' values,
    there's another problem: your code will have to depend on the class of the 
    object, which may scare you.

Solution: 
    prototype is a creational design pattern that lets you copy existing 
    objects without making your code dependent on their classes.

    the class of the object could implements clone(), and the type of all its 
    fields could implements clone(), too. so, when you call clone() on the
    object, it will delegates the cloning process to the actual objects.

    it decrease the complexity to create a copy of the object, also your code
    decoupled with the class of the object.

Variants:
    even if prototype pattern itself has no variants, but it has relation with 
    other patterns. maybe it can be seen as the variants or extensions of 
    prototypes.

    1. Prototype Manager

       This variation of the Prototype pattern involves the use of a central 
       manager or registry that stores and manages a collection of prototype 
       objects. 
       
       Clients can request a new object from the manager by specifying a key 
       or identifier, and the manager will clone and return the corresponding 
       prototype object. This can be useful when there are multiple types of 
       objects that need to be created and managed centrally.

    2. Deep Copy vs. Shallow Copy
        
       When cloning objects in the Prototype pattern, there are two 
       approaches: deep copy and shallow copy. 
       
       In a deep copy, all the properties and sub-objects of the original 
       object are also cloned, resulting in a completely independent copy. 
       In a shallow copy, only the top-level properties of the original object 
       are copied, while the sub-objects are shared between the original and 
       the clone. 
       
       The choice between deep copy and shallow copy depends on the specific 
       requirements of the application.

    3. Serialization and Deserialization
    
       Another variation of the Prototype pattern involves serializing an 
       object into a stream of bytes or a string representation, and then 
       deserializing it to create a new object. 
       
       This can be useful when objects need to be transmitted over a network 
       or stored in a persistent storage medium. The serialization and 
       deserialization process effectively creates a new object that is 
       identical to the original object.
    
Relations:

    many designs start by Factory pattern (less complicated and more 
    customized), and evolve towards Abstract Factory, Prototype or Builder 
    (more flexible but more complicated)
