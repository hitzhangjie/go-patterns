package templatemethod

import (
	"fmt"
	"testing"
)

func TestTemplateMethod(t *testing.T) {

	fmt.Println("----- financial report -----")
	r1 := NewFinancialReport()
	GenerateReport(r1)

	fmt.Println("----- education report -----")
	r2 := NewEducationReport()
	GenerateReport(r2)
}
