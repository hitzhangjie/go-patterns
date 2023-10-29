package proxy_test

import (
	"context"
	"fmt"
	"proxy"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	c := proxy.NewClient("testService")
	rsp, err := c.Call(context.TODO(), "Hello", []byte("helloworld"))
	fmt.Println(err)

	require.NotNil(t, err) // read: connection refused
	require.Nil(t, rsp)
}

func TestClientProxy(t *testing.T) {
	c := proxy.NewClientProxy("testService")
	rsp, err := c.Hello(context.TODO(), []byte("helloworld"))

	require.NotNil(t, err)
	require.Nil(t, rsp)
}
