package proxy_test

import (
	"proxy"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProxy(t *testing.T) {
	// normal use: we use the proxy object like this
	//
	// log, _ := proxy.NewSysLogger("127.0.0.1:8000", "xxx")
	// log.Debug("xxxxxxxxxxxxx")
	// log.Debug("yyyyyyyyyyyyy")
	// log.Debug("zzzzzzzzzzzzz")

	// but, we didn't provide a working syslog service,
	// so initialize the proxy object and log with it will fail.
	//
	// well, here we just use it to show how Proxy Pattern works.
	// so, you can just ignore the asserts here.
	log, err := proxy.NewSysLogger("127.0.0.1:8000", "xxx")
	require.Error(t, err)
	require.Panics(t, func() {
		log.Debug("xxxxxxxxxxxxx")
		log.Debug("yyyyyyyyyyyyy")
		log.Debug("zzzzzzzzzzzzz")
	})
}
