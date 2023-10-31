Context:

    In some scenarios, multiple objects may be able to handle a request:
    - request may be processed by all request one by one 
    - request may be processed by all objects in some order which matters
    - request may be processed by several objects rather than all
    - the objects maybe changed at runtime, and cannot be known in advance

    Here, multiple objects actually means some 'logic', like an article, the
    system will do some checking before accept this article:
    - checking if the user valid
    - checking if the user authorized
    - checking if the article data is empty
    - checking if the article data contains dirty words
    - checking if the article data is copied from others (checking similarity)
    - checking if the article data contains low resolution pictures 
    - checking if the article data has typo
    - select the proper cover for the article
    - select the title for the article
    - generate abstract for the article
    - formatting the article data, and convert it into structured data in order
      to render it in web, h5 or mobile app.
    - 
   
Solution: 

    The Chain of Responsibility pattern is a behavioral design pattern that
    allows an object to pass a request along a chain of potential handlers 
    until the request is handled or reaches the end of the chain. This pattern
    promotes loose coupling between the sender and receiver of a request,
    providing flexibility in dynamically assigning responsibilities.

    The Chain of Responsibility pattern addresses this by creating a chain of 
    objects, each having a chance to handle the request. 
    
    The sender of the request is unaware of the objects in the chain and 
    simply sends the request to the chain or the first object in the chain.


    
    The Chain of Responsibility pattern consists of three main components: 
    - Handler interface
      The Handler interface defines a common method for handling requests and 
      maintaining a reference to the next handler in the chain. 
    - ConcreteHandler classes
      Each ConcreteHandler class implements the handling logic and decides 
      whether to handle the request or pass it to the next handler in the chain. 
    - Client
      It creates the chain of handlers and sends requests to the first handler.

Variants:

    - Dynamic Chain: In this variant, the chain of handlers can be modified
      dynamically at runtime. Handlers can be added or removed from the chain
      based on certain conditions or events. This variant provides flexibility 
      in managing the chain of responsibility.

Relation: 

    The Chain of Responsibility pattern is often used in conjunction with other
    design patterns. It is commonly used with the Composite pattern, where a
    handler can be a composite object that contains multiple handlers. It can
    also be combined with the Command pattern, where a request is encapsulated
    as an object and passed along the chain.

Overall, the Chain of Responsibility pattern provides a flexible and extensible
way to handle requests by creating a chain of potential handlers. It promotes
loose coupling and allows for dynamic assignment of responsibilities, making it
a valuable pattern in many software development scenarios.