package templatemethod

import "fmt"

type Reporter interface {
	initializeReport()
	collectData()
	analyze()
	finalize()
	finalReport()
}

// Abstract base class
type baseReport struct {
}

func GenerateReport(rr Reporter) {
	rr.initializeReport()
	rr.collectData()
	rr.analyze()
	rr.finalize()
	rr.finalReport()
}

// Base class steps implemented
func (r *baseReport) initializeReport() {
	fmt.Println("Initializing report...")
}

func (r *baseReport) collectData() {
	fmt.Println("Collecting data...")
}

func (r *baseReport) analyze() {
	fmt.Println("Analyzing report data...")
}

func (r *baseReport) finalize() {
	panic("not implemented")
}

func (r *baseReport) finalReport() {
	panic("not implemented")
}

func NewFinancialReport() *FinancialReport {
	return &FinancialReport{}
}

// Concrete sub-class
type FinancialReport struct {
	baseReport // Embedded Report
}

// Override finalize method
func (f *FinancialReport) finalize() {
	fmt.Println("💲Finanical finalize")
}

// Override hook method
func (f *FinancialReport) finalReport() {
	fmt.Println("💲Generating financial report")
}

func NewEducationReport() *EducationReport {
	return &EducationReport{}
}

// Concrete sub-class
type EducationReport struct {
	baseReport // Embedded Report
}

// Override finalize method
func (r *EducationReport) finalize() {
	fmt.Println("📚 Education finalize")
}

// Override hook method
func (r *EducationReport) finalReport() {
	fmt.Println("📚 Generating education report")
}
