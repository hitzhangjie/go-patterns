Context: 
    when you want to build some products (products could be any objects), 
    while you don't want to bind the products with the specific class. 
    We can use a factory to serve as the producer.

Solution: 
    view the file.puml, the factory is responsible for create the different 
    products like 'IProductA' and 'IProductB'.

variants: 
    1. abstract factory pattern: when you want to build different styles 
       of different products, and you don't want to bind the products 
       with the specific class.
    2.