Context: 
    when you want to build some products (products could be any objects) with 
    different styles, while you don't want to bind the products/styles with 
    the specific class. We can use an abstract factory to serve as the producer.

Solution: 
    view the file.puml, the IFactory is responsible for create the different 
    products like 'IProductA' and 'IProductB' with different styles.

Variants: 
    1. factory pattern: when you want to build different products with only one 
        style, and you don't want to bind the products with the specific class.
    2. 