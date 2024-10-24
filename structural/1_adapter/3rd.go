package adapter

import (
	"encoding/json"
	"fmt"
)

type PowerfulAnalyzer struct {
}

func (_ *PowerfulAnalyzer) Analyze(dat []byte) error {
	recs := []Record{}
	if err := json.Unmarshal(dat, &recs); err != nil {
		return fmt.Errorf("invalid JSON format: %v", err)
	}

	fmt.Println("analyzing ... done")
	return nil
}
