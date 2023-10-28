package adapter

import (
	"encoding/json"
	"encoding/xml"
)

type Analyzer interface {
	Analyze([]byte) error
}

func NewJSONAnalyzerAdapter(analyzer *PowerfulAnalyzer) Analyzer {
	return &jsonAnalyzerAdapter{analyzer: analyzer}
}

type jsonAnalyzerAdapter struct {
	analyzer *PowerfulAnalyzer
}

func (ja *jsonAnalyzerAdapter) Analyze(dat []byte) error {
	// convert the XML data to JSON so as to fit the ja.analyzer
	recs := []Record{}
	if err := xml.Unmarshal(dat, &recs); err != nil {
		return err
	}

	dat, err := json.Marshal(recs)
	if err != nil {
		return err
	}

	// call wrapped ja.analyzer
	return ja.analyzer.Visualize(dat)
}
