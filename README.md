# README

[TOC]

## Introduction

go-patterns actually is short for design patterns in go, which shows list of demos behind the concepts.

We must think about the following questions before we dive into the demos, it is really important.

- What's a pattern?
- What's a design pattern?
- History of patterns?
- Why should I learn patterns?
- Criticism of patterns?
- Classification of patterns?
- etc.
 
## What's a pattern?

Pattern is a solution, which works well in practices, to a commonly reoccurring problems. Patterns could be created and shared in every area, like building body, improving representation skills, architecture skills, or software design.

I'm a developer, when I talk about patterns, I mean the software design pattern. Learning patterns can help us build software on the collective experience of skilled software engineers. Then when we work on a particular problem, we could recall a similar problem which had already been solved and reuse the essense of its solution to solve this new problem. That's `举一反三` in Chinese.

With the help of patterns, novices will work better as if they were (or almost as if they were) experts on modest-sized projects, without having to gain many years of experience.

## What makes a pattern?

A pattern for software architecture describes a particular recurring design problem that arises in specific design contexts, and presents a well-proven generic scheme for its solution. The solution scheme is specified by describing its consitituent components, their responsibilities and relationships, and the ways in which they collaborate.  

See: Pattern-Oriented Software Architecture, Volume 1, Page 8~11.

- Context: a situation giving rise to a problem.
- Problem: the recurring problem arising in that context.
- Solution: a proven resolution of the problem.

## Pattern Categories

In software design area, patterns could be split into different layers. 

A closer look at existing patterns reveals that they cover various ranges of scale and abstraction. 
- Some patterns help in structuring a software system into subsystems. 
- Other patterns support the refinement of subsystems and components, or of the relationships between them. 
- Further patterns help in implementing particular design aspects in a specific programming language.

Patterns also range from domain-independent ones, such as those for decoupling interacting components, to patterns addressing domain-specific aspects such as transaction policies in business applications, or call routing in telecommunication.

To refine our classification, we group patterns into three categories:

- Architecture Patterns  
  Viable software architectures are built according to some overall structuring principle. We describe these principles with architectural patterns.  
  >A architectural pattern expresses a fundamental structural organization schema for software systems. It provides a set of predefined subsystems, specifies their responsibilities, and includes rules and guidelines for organizing the relationships between them.  
  
  Architectural patterns are templates for concrete software architectures. They specify the system-wide structual properties of an application, and have an impact on the architecture of its subsystems. The selection of an architecture pattern is therefore a fundamental design decision when developing a software system.
  
- Design Patterns 
  The subsystems of a software architecture, as well as the relationships between them, usually consist of several smaller architectural units. We describe these using design patterns.
  >A design pattern provides a scheme for refining the subsystems or components of a software system, or the relationships between them. It describes a commonly-recurring structure of communicating components that solves a general design problem within a particular context.
  
  Design patterns are medium-scale patterns. They are smaller in scale than architectural patterns, but tend to be independent of a particular programming language or programming paradigm. The application of a design pattern has no effect on the fundamental structure of a software system, but may have a strong influence on the architecture of a subsystem.  
  Many design patterns provides structures for decomposing more complex services or components. Others address the effective cooperation between them, such as the following pattern: Observer or Publisher-Subscriber.
   
- Idioms
  Idioms deal with the implemention of particular design issues.
  >An idiom is a low-level pattern specific to a programming language. An idiom describes how to implement particular aspects of components or the relationships between them using the features of the given language.
  
  Idiom represent the lowest-level patterns. They address aspects of both design and implemention.  
  Most idioms are language-specific, they capture existing programming experience. Often the same idiom looks different for different languages, and sometimes an idiom that is useful for one programming language doesn't make sense in another.

## Relationships between Patterns

A pattern solves a particular problem, but its application may raise new problems. Some of these can be solved by other patterns.

Most patterns for software architecture raise problems that can be solved by smaller patterns. Patterns do not usually exist in isolation. Each pattern depends on the smaller patterns it contains and on the larger patterns in which it is contained. 

And a pattern may also be a variant of another.

Patterns can also combine in more complex structures at the same level of abstraction. Each pattern resolves a particular subset of the forces to balance the forces when solving the problem.

## Pattern Description

Patterns must be presented in an appropriate form if we are to understand and discuss them. A good description helps us grasp the essense of a pattern immediately:

- what's the problem the pattern addresses?
- what's the proposed solution?  
  A good description also provides us with all the details necessary to implement a pattern, and to consider the consequences of its application.
- describe the solution uniformly!  
  This helps us to compare one pattern with another, especially when we are looking for alternative solutions to a problem.

The basic Context-Prolem-Solution structure provides a good starting point for a description format, but it is not enough.

A pattern must be named - preferably with an intuitive name - if we are to share it and discuss it.

OK, a good pattern template is showed below, please use it as a starting point:

```
          Name   The name and a short Summary of the pattern.

 Also Known As   Other names for the pattern, if any are known.
       Example   A real-world example demonstrating the existence of the problem and the need for 
                 the pattern.
                 Throughout the description we refer to the example to illustrate solution and
                 implementation aspects, where this is necessary for useful.
       Context   The situations in which the pattern may apply.
       Problem   the problem the pattern addresses, including a discussion of its associated forces.
      Solution   The fundamental solution principle underlying the pattern.
     Structure   A detailed specification of the structural aspects of the pattern, including 
                 CRC-cards and OMT  class diagram.
      Dynamics   Typical scenarios describing the runtime behavior of the pattern. We further 
                 illustrate the scenarios with Object Message Sequences Charts.
Implementation   Guidelines for implementing the pattern.
        ......   .......
      Variants   A brief description of variants or specializations of a pattern.
    Known Uses   Examples of the use of this pattern, taken from existing systems.
  Consequences   The benefits the pattern provides, and any potiential liabilities.
      See Also   References to patterns that solve similar problems, and to patterns that help us
                 refine the pattern we are describing.
```

## Will you learn patterns?

Now I think your answer will be definitely 100% Yes.

## About This Repository

go-patterns will aim to provide demos for design patterns and idioms in golang. Hope we could follow the design details of design patterns to quickly solve the recurring problems.

