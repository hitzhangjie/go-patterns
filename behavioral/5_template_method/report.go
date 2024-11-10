package templatemethod

import "fmt"

// In go, there's no inheritance, so we use composition instead of inheritance.
//
// A very common mistake is define GenerateReport() as a method in Reporter interface{},
// and implement it in baseReport, like this:
//
// ```
// type Reporter interface {
//     GenerateReport()
//     ...
// }
//
// type baseReport struct {}
//
// func (br *baseReport) GenerateReport() {
//     br.initializeReport()
//     br.collectData()
//     br.analyze()
//     br.finalize()          // <= this is a mistake!
//     br.finalReport()       // <= this is a mistake!
// }
//
// type FinancialReport struct {
//     baseReport
// }
//
// func (fr *FinancialReport) finalize() { ... }
// func (fr *FinancialReport) finaleport() { ... }
//
// Then you may want to this:
//
// ```
// var r = &FinancialReport{}
// r.GenerateReport() // <= this will call baseReport's GenerateReport(),
//                    // <= but this template method will call baseReport's finalize() and finaleport(), not FinancialReport's finalize(), and finaleport().
//                    // <= why? because composition is not inheritance! baseReport isn't superclass of FinancialReport!
//                    // <=      in golang, there's no abstract class, only interface and concreate class..
//
// An easy way to work around this, is define the Template Method as package level function.
// ```

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
