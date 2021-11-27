package random

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"

	"syreclabs.com/go/faker"
)

var debugFlow bool

type employee struct {
	personId  int
	firstName string
	lastName  string
}

func FakeDataMain(enableDebug bool) {
	debugFlow = enableDebug
	// randomName()
	// randomNumberInRange(1, 9)
	/*
		result := staticEmployee()
		fmt.Printf("result is: %v with firstName %v\n", result, result.firstName)
		employeeStruct(false)
		employeeStruct(true)
	*/
	employeeSliceOfStructs()
}

func randomName() string {
	name := faker.Name().FirstName()
	fmt.Printf("Random name: %v\n", name)
	return name
}

func randomNumberInRange(minNum int, maxNum int) int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(maxNum-minNum) + minNum
	fmt.Printf("Random number: %v\n", num)
	return num
}

func employeeSliceOfStructs() {
	var allEmployees []*employee
	numRange := randomNumberInRange(2, 10)
	for i := 0; i < numRange; i++ {
		// fmt.Println("index:", i)
		randomEmployee := employeeStruct(false)
		allEmployees = append(allEmployees, randomEmployee)
	}
	fmt.Printf("allEmployees: %v\n", allEmployees)
	fmt.Printf("Index 0 Employee: %v with first name: %v\n", allEmployees[0], allEmployees[0].firstName)
}

func staticEmployee() *employee {
	newEmployee := employee{123, "John", "Smith"}

	fmt.Println("# staticEmployee")
	fmt.Printf("The type is %v\n", reflect.TypeOf(newEmployee))
	fmt.Printf("The kind is %v\n", reflect.TypeOf(newEmployee).Kind())
	fmt.Printf("The value is %v\n", reflect.ValueOf(newEmployee))
	fmt.Printf("The firstName is %v\n", newEmployee.firstName)
	return &newEmployee
}

func employeeStruct(static bool) *employee {
	var newEmployee employee

	if static {
		newEmployee = employee{456, "Jane", "Doe"}
	} else {
		newEmployee = employee{
			faker.Number().NumberInt(3),
			faker.Name().FirstName(),
			faker.Name().LastName(),
		}
	}

	if debugFlow {
		fmt.Println("# employeeStruct")
		fmt.Printf("The type is %v\n", reflect.TypeOf(newEmployee))
		fmt.Printf("The kind is %v\n", reflect.TypeOf(newEmployee).Kind())
		fmt.Printf("The value is %v\n", reflect.ValueOf(newEmployee))
		fmt.Printf("The firstName is %v\n", newEmployee.firstName)
	}
	return &newEmployee
}
