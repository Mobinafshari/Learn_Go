package main

type Payable interface {
	CalculatePay() float64
}

type SalariedEmployee struct {
	Name         string
	AnnualSalary float64
}

func (se SalariedEmployee) CalculatePay() float64 {
	return se.AnnualSalary / 12
}

type HourlyEmployee struct {
	Name       string
	HourlyRate float64
	HourWorked float64
}

func (he HourlyEmployee) CalculatePay() float64 {
	return he.HourWorked * he.HourlyRate
}

func main() {

}
