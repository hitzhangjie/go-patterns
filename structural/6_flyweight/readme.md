Context:

    There're too many similar objects or copies created, so the RAM consumes
    very high! These objects (or copies) contains the serveral same attributes,
    while other fields are different.

    The objects should be created, but how to make the RAM consumption low,
    and how to avoid creating redundant uncessary objects.

Solutions:

    The flyweight pattern is used in situations where a large number of objects
    need to be created and stored in memory. For example, a graphic editor needs
    to handle thousands of shape objects like lines, circles, rectangles etc.
    
    Creating full-fledged objects for each shape would be expensive in terms of
    memory usage. Especially since many of these objects share common states
    like color, stroke style, fill pattern etc.
    
    Without flyweights, the editor would consume large amounts of memory
    creating redundant copies of these shared states in each object. This is
    inefficient and can cause the application to slow down due to excessive
    memory usage.
    
    The flyweight pattern provides a solution by separating shared intrinsic
    states from extrinsic object-specific states. The intrinsic states are
    stored in flyweight objects that can be shared across all shape objects. For
    example, a Circle flyweight would store the data needed to render a circle.
    
    The shape-specific extrinsic states like object coordinates are stored
    externally and associated with flyweight references. This allows significant
    memory savings and improves performance of the editor by reusing flyweight
    objects instead of duplicating states.
    
    This background motivates the need for flyweights to efficiently handle
    large object populations in memory-constrained contexts. Let me know if you
    need any clarification or have additional suggestions!
    
Variants:

Relation: