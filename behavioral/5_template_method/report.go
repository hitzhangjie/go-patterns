package templatemethod

import "fmt"

type Reporter interface {
	//mistake1: composition isn't inheritance, you can't overwritten the
	//          base class method
	//GenerateReport()
	//
	//mistake2: defines this method as interface method, it's weired by
	//          calling `r.GenerateReport(rr)`, it works but weired.
	//GenerateReport(rr Reporter)
	//
	// so, just define `GenerateReport(rr Reporter)` as the template method.

	initializeReport()
	collectData()
	analyze()
	finalize()
	finalReport()
}

// Abstract base class
type Report struct {
}

// Template method that calls steps
func GenerateReport(rr Reporter) {
	rr.initializeReport()
	rr.collectData()
	rr.analyze()
	rr.finalize()
	rr.finalReport()
}

// Base class steps implemented
func (r *Report) initializeReport() {
	fmt.Println("Initializing report...")
}

func (r *Report) collectData() {
	fmt.Println("Collecting data...")
}

func (r *Report) analyze() {
	fmt.Println("Analyzing report data...")
}

func (r *Report) finalize() {
	fmt.Println("Finalizing report...")
}

// Hook method for customization
func (r *Report) finalReport() {
	// Default implementation
	fmt.Println("Generating final report")
}

// Concrete sub-class
type FinancialReport struct {
	Reporter // Embedded Report
}

// Override finalize method
func (f *FinancialReport) finalize() {
	fmt.Println("Finanical finalize")
}

// Override hook method
func (f *FinancialReport) finalReport() {
	fmt.Println("Generating financial report")
}

// Concrete sub-class
type EducationReport struct {
	Reporter // Embedded Report
}

// Override finalize method
func (r *EducationReport) finalize() {
	fmt.Println("Education finalize")
}

// Override hook method
func (r *EducationReport) finalReport() {
	fmt.Println("Generating education report")
}
