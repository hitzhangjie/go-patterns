Context:

    The proxy design pattern provides a surrogate or placeholder object that
    controls access to another object, called the real object. The proxy acts 
    as an interface to the real object and controls access to it. Some common
    situations where the proxy pattern is applicable:
    
    - Remote proxies provide a local representative for a real object in a
      different address space. This is commonly used for accessing remote 
      objects across networks.
    - Virtual proxies create expensive real objects on demand. The proxy 
      handles requests by creating the real object only when needed (lazy 
      initialization).
    - Protection proxies control access to the real object. This is used for
      access rights and permissions.
    - Smart proxies add additional logic to the real object request handling 
      like caching frequently used results.

Solutions:

    - The real object defines the common interface that is implemented by both
      the actual real object and the proxy so they can be used interchangeably.
    - The proxy keeps a reference to the real object and implements the same
      interface. The proxy handles the request before/after forwarding to the 
      real object.
    - The client interacts with the real object via the common interface. The
      usage is identical for real or proxy objects.
    
Variants:

    - Remote proxies manage interaction with remote real objects.
    - Virtual proxies defer creation of expensive real objects.
    - Protection proxies control access permissions to real objects.
    - Smart proxies add caching and more logic for real objects.

Relation:

    The proxy pattern provides a placeholder object that controls access to the
    real object. This allows additional functionality like remote access,
    on-demand creation, permissions and caching. The client interacts with the
    proxy identically to the real object since they share a common interface.
