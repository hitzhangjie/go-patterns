Context: 
    the object you want to build is too complex, we need too many parameters
    to control how to build the object. 

    if we define a constructor to accept all parameters, the parameter list 
    will be too long, that's difficult to be understood and maintained.

    maybe we could pack all parameters into a new type named like 'config' or 
    'options'. well, another problems occur: 1) do we have to set a value for all 
    'config' or 'options' fileds? it's likely not. 2) if we set several fields
    of it, are you sure you know the exact behavior of the contstructor? No.

    we need to build the complex object more conveniently and more clearly.

Solution: 
    use builder pattern to build the final object step by step.

    don't create the complex object by its own class's constructor, but rather
    move the creation logic to a separate builder class. this builder will
    build the complex object step by step.

    this solution will eliminate the long parameter list and the complex
    creation logic.

Variants:
    use function option pattern to control the constructor behavior. the 
    constructor define default values for all underlying parameters that 
    affects how to build the complex object, we can use function option 
    pattern to change the default parameter value to change the default 
    building logic.
