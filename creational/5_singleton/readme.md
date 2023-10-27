Context:

    sometimes you want to ensure that a class has only one single instance.

    1. the most common reason is to control access to some shared resources, 
       for example, a database or a file.
    2. provides a global access to that instance. you may store a essential
       object into a global variable that use the global variable here and 
       there in your code. but the global variable may be overwritten by
       other code.

Solution:

    singleton pattern can solve these two problems at the same time.
    (ps: violating the Single Responsibility Principle) :)

    for different languages, there maybe different idioms to this, all should
    focus on these points:
    - encapsulation and visibility, don't expose other ways to create new 
      object instance, for example, prevent users calling constructor 
      `new(ClassType)` or access the class `v := &ClassType{}` directly.
    - create the instance only once, when checking if one instance has been
      created before, use sync.Once or static variable or doublechecking lock
      to effiently check this.

Variants:


