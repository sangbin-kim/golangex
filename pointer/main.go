package main

import "fmt"

type Student struct {
	name string
	age  int

	grade string
	class string
}

func (s *Student) PrintGPA() {
	fmt.Println(s.class, s.grade)
}

func (s *Student) InputGPA(class string, grade string) {
	s.class = class
	s.grade = grade

}
func main() {
	var s Student = Student{name: "Sangbin", age: 29, class: "수학", grade: "A"}

	s.InputGPA("history", "C")
	s.PrintGPA()
}
