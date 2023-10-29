// Even if we use term 'proxy' here and there in this file, but actually,
// this demo isn't a Proxy Pattern in Design Pattern area, but indeed, it's
// a Local Proxy Pattern in Transport Pattern in Architecture Patterns.
//
// Well, in design pattern area, the wrapper object and the wrapped real object
// has the common interface.

package proxy

import (
	"context"
	"fmt"
	"math/rand"
	"net"
)

type Client interface {
	Call(ctx context.Context, method string, req []byte) ([]byte, error)
}

type NamingService interface {
	Register(ctx context.Context, service string, addr net.Addr) error
	List(ctx context.Context, service string) ([]Node, error)
}

type simpleNamingService struct {
}

func (s *simpleNamingService) Register(ctx context.Context, service string, addr net.Addr) error {
	return nil
}

func (s *simpleNamingService) List(ctx context.Context, service string) ([]Node, error) {
	nodes := []Node{
		{
			Addr:  "127.0.0.1:8000",
			Proto: "udp",
			Codec: "simple",
		},
	}
	return nodes, nil
}

type Node struct {
	Addr  string
	Proto string //  TCP or UDP
	Codec string
}

type LoadBalancer interface {
	Select(nodes []Node) (Node, error)
}

type randomLoadBalancer struct {
}

func (b *randomLoadBalancer) Select(nodes []Node) (Node, error) {
	i := rand.Int() % len(nodes)
	return nodes[i], nil
}

type Codec interface {
	Encode(body []byte) (pkg []byte, err error)
	Decode(pkg []byte) (body []byte, err error)
}

type simpleCodec struct {
}

func (c *simpleCodec) Encode(body []byte) (pkg []byte, err error) {
	return body, nil
}

func (c *simpleCodec) Decode(dat []byte) (body []byte, err error) {
	return dat, nil
}

// client proxies the request to the connection to the remote service
type client struct {
	service string

	naming   NamingService
	balancer LoadBalancer
	codec    Codec
}

func NewClient(service string) Client {
	c := &client{
		service:  service,
		naming:   &simpleNamingService{},
		balancer: &randomLoadBalancer{},
		codec:    &simpleCodec{},
	}
	return c
}

func (c *client) Call(ctx context.Context, method string, req []byte) ([]byte, error) {
	nodes, err := c.naming.List(ctx, c.service)
	if err != nil {
		return nil, fmt.Errorf("query service nodes fail: %v", err)
	}
	fmt.Printf("query service nodes: %v\n", nodes)

	node, err := c.balancer.Select(nodes)
	if err != nil {
		return nil, fmt.Errorf("select node fail: %v", err)
	}
	fmt.Printf("select node: %v\n", node)

	pkg, err := c.codec.Encode(req)
	if err != nil {
		return nil, fmt.Errorf("encode req fail: %v", req)
	}
	fmt.Printf("encode req ok, pkg:%v\n", pkg)

	conn, err := net.Dial(node.Proto, node.Addr)
	if err != nil {
		return nil, fmt.Errorf("dial node: %v fail: %v", node, err)
	}
	fmt.Printf("dial node ok, conn: %v\n", conn)

	t, ok := ctx.Deadline()
	if ok {
		conn.SetDeadline(t)
	}
	fmt.Printf("set req deadline ok, deadline: %v\n", t)

	n, err := conn.Write(pkg)
	if err != nil || n != len(pkg) {
		return nil, fmt.Errorf("write pkg fail: %v, write bytes: %d != %d", err, n, len(pkg))
	}
	fmt.Printf("write pkg ok, %d bytes written\n", n)

	rspPkg := make([]byte, 65536)
	n, err = conn.Read(rspPkg)
	if err != nil {
		return nil, fmt.Errorf("read response err: %v", err)
	}
	fmt.Printf("read response ok, rsp pkg: %v\n", rspPkg)

	rspBody, err := c.codec.Decode(rspPkg)
	if err != nil {
		return nil, fmt.Errorf("decode response pkg fail: %v", err)
	}
	fmt.Printf("decode response ok, rsp body: %v\n", rspBody)

	return rspBody, nil
}

func NewClientProxy(service string) *ClientProxy {
	return &ClientProxy{
		NewClient(service),
	}
}

type ClientProxy struct {
	Client
}

func (p *ClientProxy) Hello(ctx context.Context, req []byte) ([]byte, error) {
	return p.Client.Call(ctx, "Hello", req)
}
