package adapter_test

import (
	"adapter"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdapter(t *testing.T) {
	stock := new(adapter.Stock)
	xmldata, _ := stock.Data()

	// triggers an error because Analyzer only accepts JSON data
	az := new(adapter.PowerfulAnalyzer)
	require.Error(t, az.Visualize(xmldata))

	// create an adapter to do this
	adapter := adapter.NewJSONAnalyzerAdapter(az)
	require.Nil(t, adapter.Analyze(xmldata))
}
