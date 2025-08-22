package main

import "fmt"

type Employee interface {
	salary() float64
}

type fulltime struct {
	workingDays int
}

func (f fulltime) salary() float64 {
	return float64(f.workingDays) * 15000
}

type contractor struct {
	workingDays int
}

func (c contractor) salary() float64 {
	return float64(c.workingDays) * 3000
}

type freelancer struct {
	workinghours int
}

func (ff freelancer) salary() float64 {
	return (float64(ff.workinghours) / 20) * 2000
}
func main() {
	ftEmp := fulltime{
		20,
	}
	cEmp := contractor{
		16,
	}
	fEmp := freelancer{
		89,
	}
	arr := []Employee{ftEmp, cEmp, fEmp}
	for _, val := range arr {
		fmt.Printf("The salary of the Employee is %v\n", val.salary())
		if _, ok := val.(contractor); ok {
			fmt.Println("This employee is a contractor")
		}
		if _, ok := val.(freelancer); ok {
			fmt.Println("This employee is a freelancer")
		}
		if _, ok := val.(fulltime); ok {
			fmt.Println("This employee is a freelancer")
		}
	}

}
