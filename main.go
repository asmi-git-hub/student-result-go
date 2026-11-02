package main

import (
	"fmt"

	"student-result-go/result"
	"student-result-go/student"
)

func main() {

	var name string
	var rollNo int

	fmt.Print("Enter Student Name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter Roll No: ")
	fmt.Scanln(&rollNo)

	marks := make([]int, 5)

	subjects := []string{
		"Tamil",
		"English",
		"Maths",
		"Science",
		"Computer",
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("Enter %s Mark: ", subjects[i])
		fmt.Scanln(&marks[i])
	}

	s := student.Student{
		Name:   name,
		RollNo: rollNo,
		Marks:  marks,
	}

	total := result.CalculateTotal(s.Marks)
	average := result.CalculateAverage(total, len(s.Marks))
	grade := result.CalculateGrade(average)
	status := result.GetResult(s.Marks)

	fmt.Println("\n----- STUDENT RESULT -----")
	fmt.Println("Student Name :", s.Name)
	fmt.Println("Roll No      :", s.RollNo)
	fmt.Println("Total        :", total)
	fmt.Printf("Average      : %.2f\n", average)
	fmt.Println("Grade        :", grade)
	fmt.Println("Result       :", status)
}