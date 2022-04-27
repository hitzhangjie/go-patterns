context: when you want to build some products (products could be any objects) with different styles, while you don't want to bind the products/styles with the specific class. We can use a abstract factory to serve as the producer.

solution: view the file.puml, the IFactory is responsible for create the different products like 'IProductA' and 'IProductB' with different styles.

variants: 
- factory pattern: when you want to build different products with only 1 style, and you don't want to bind the products with the specific class.