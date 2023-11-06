package templatemethod

import (
	"fmt"
	"testing"
)

func TestTemplateMethod(t *testing.T) {

	fmt.Println("----- financial report -----")
	r1 := &FinancialReport{&Report{}}
	GenerateReport(r1)

	fmt.Println("----- education report -----")
	r2 := &EducationReport{&Report{}}
	GenerateReport(r2)
}
