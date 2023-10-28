Context:

    Supposing you're developing a monitor and analyze system for a stock 
    market. You can export stock market data to XML format, and you have 
    written some core classes to analyze it.

    But days later, you want to import a 3rd-party library to enhance your 
    system. Well, then you find it only supports JSON data.

    You want to the source code of 3rd-party library to let it support XML, 
    but:
    - the 3rd-party may be not opensourced.
    - the 3rd-party library maybe complex.
    - you may make bugs when change the code.

Solution:

    Use adapter pattern to create an adapter between the XML data to the 
    3rd-party library which only supports JSON data. 
    
    The adapter can provide the APIs which is same as the 3rd-party library, 
    and it internally convert the XML data to JSON format, then pass it to
    the 3rd-party library object.

    The 3rd-party library object could be wrapped into the adapter itself,
    you can pass the 3rd-party library object as the adapter constructor 
    paramemter, or implicitly create one 3rd-party object.

    ps: I suggest using the first method, to make the invisible visible!

    Adapter is actually a special object that converts the interface of one 
    object to the proper interface so that another object can understand it.

    ps: You an implement Adapter Pattern in composition way, while in C++ you 
    can use multiple inheritance. Different programming languages has its own
    idioms.

Aliases:

    wrapper, it wraps an object which contains interface that we want to  fit.

Variants:


Relations:
