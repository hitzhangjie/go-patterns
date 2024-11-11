Context:

    Supposing that you're writing a text-editor application, your current task
    is to create a toolbar with many buttons for open, save, undo, cancel
    and other buttons.

    While all of these buttons look similar, they will do different things
    when you press them. So you want to implement these buttons by creating
    some subclasses for each of, open, save, undo, cancel action.

    Well, two bad things will occur:
    - you'll see soon you have to maintain too enormous number of subclasses,
      when you change the logic of base class of button, you may worried about
      if it affect the behavior of all subclasses. That means your GUI code
      will become dependent on the volatile code of the business logc.
    - you may want to add shortcuts for different buttons for friendly
      accessibility. when you activating the shortcuts, you may write some code
      in its subclass. Now we want to add shortcuts, when shortcuts triggered,
      you have to copy & paste the duplicate code there.
    - And besides, sometimes, you may want to lazily process the user request,
      so you may want to queue the actions of the player.

    OK, how should we solve these kind of problems.

Solution:

    Command Pattern is a behavioral pattern that turns a request into a
    standalone object that contains all information about the request. This
    transformation lets you pass requests as a method arguments, delay or
    queue a request’s execution, and support undoable operations.

    And Command Pattern is a way to separate the thing calling for a task from
    the thing doing the task.

    For example, think of ordering food at a restaurant. The waiter takes your
    order (the command) and gives it to the cook (the receiver). The waiter
    doesn't need to know how to cook - they just pass the order to someone who
    does.

    The Command Pattern works similarly in code. You create a command object
    that has all the details needed to perform an action. This command is given
    to an invoker object, which calls the command's execute method. This invokes
    the receiver object to carry out the command's action.

    The invoker doesn't need to know how the receiver performs the action. The
    receiver doesn't need to know who called for the action. The command object
    serves as the link between them.

    This separation makes the code more modular. It decouples the execution from
    initiation. You can queue commands, execute them remotely, or implement
    undo/redo because the command is a self-contained task.

    In summary, the command pattern allows loosely coupled code by packaging
    tasks into stand-alone objects that fully define the operation to be
    performed. The sender triggers the command, the receiver executes it, and
    the command binds them.

Variants:

Reletion:
