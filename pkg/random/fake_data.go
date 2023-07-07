package random

import (
	"fmt"
	"reflect"

	"github.com/andrew-j-price/journey/pkg/helpers"
	"syreclabs.com/go/faker"
)

var debugFlow bool

// NOTE: The json package can only serialize exported fields of your struct.
// Therefore start all fields with an uppercase letter so they can be included in the output
type employee struct {
	ID        int
	FirstName string
	LastName  string
	Active    bool
}

func FakeDataMain(enableDebug bool) {
	debugFlow = enableDebug
	// randomName()
	// helpers.RandomNumberInRange(1, 9)
	/*
		result := staticEmployee()
		fmt.Printf("result is: %v with firstName %v\n", result, result.firstName)
		employeeStruct(false)
		employeeStruct(true)
	*/
	EmployeeSliceOfStructs()
}

func randomName() string {
	name := faker.Name().FirstName()
	fmt.Printf("Random name: %v\n", name)
	return name
}

func EmployeeSliceOfStructs() []*employee {
	var allEmployees []*employee
	numRange := helpers.RandomNumberInRange(2, 10)
	for i := 0; i < numRange; i++ {
		// fmt.Println("index:", i)
		randomEmployee := employeeStruct(false)
		allEmployees = append(allEmployees, randomEmployee)
	}
	// NOTE: old reference for print analysis
	if false {
		fmt.Printf("allEmployees: %v\n", allEmployees)
		fmt.Printf("Index 0 Employee: %v with first name: %v\n", allEmployees[0], allEmployees[0].FirstName)
	}
	return allEmployees
}

func staticEmployee() *employee {
	newEmployee := employee{123, "John", "Smith", true}

	fmt.Println("# staticEmployee")
	fmt.Printf("The type is %v\n", reflect.TypeOf(newEmployee))
	fmt.Printf("The kind is %v\n", reflect.TypeOf(newEmployee).Kind())
	fmt.Printf("The value is %v\n", reflect.ValueOf(newEmployee))
	fmt.Printf("The firstName is %v\n", newEmployee.FirstName)
	return &newEmployee
}

func employeeStruct(static bool) *employee {
	var newEmployee employee

	if static {
		newEmployee = employee{456, "Jane", "Doe", true}
	} else {
		newEmployee = employee{
			faker.Number().NumberInt(3),
			faker.Name().FirstName(),
			faker.Name().LastName(),
			helpers.RandomBool(),
		}
	}

	if debugFlow {
		fmt.Println("# employeeStruct")
		fmt.Printf("The type is %v\n", reflect.TypeOf(newEmployee))
		fmt.Printf("The kind is %v\n", reflect.TypeOf(newEmployee).Kind())
		fmt.Printf("The value is %v\n", reflect.ValueOf(newEmployee))
		fmt.Printf("The firstName is %v\n", newEmployee.FirstName)
	}
	return &newEmployee
}
