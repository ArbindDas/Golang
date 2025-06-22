package main

import "fmt"

type Calculator struct{}

// method with value receiver
func (this Calculator) AddValue(num1 int, num2 int) {
	fmt.Println("Addition is : ", num1+num2)
}

// method with pointer receiver
func (this *Calculator) Sum(num int, num2 int) {

	fmt.Println("sum is -> ", num+num2)
}
func main() {

	addAns := Calculator{}
	addAns.AddValue(90, 8)

	addAns.Sum(90, 90)

	fmt.Println("jai shree ram")
}
