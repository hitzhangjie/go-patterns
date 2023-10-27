package singleton

import (
	"singleton/report"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSingleton(t *testing.T) {
	// package report only exports interface `Reporter`` and function `GetInstance()`.
	//
	// 1. you can't create an instance via an interface{} definition
	// 2. you can only call function `GetInstance()`
	r1 := report.GetInstance()
	r1.Report("hello world 1")

	r2 := report.GetInstance()
	r2.Report("hello world 2")

	require.Equal(t, r1, r2)
}
