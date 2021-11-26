package random

import (
	"fmt"
	"reflect"

	"github.com/andrew-j-price/journey/logger"
)

func RandomLogMessages() {
	fmt.Println("need to move logging into its own package")
	logger.Info.Println("Informational only")
	logger.Warn.Println("Warning message")
	logger.Error.Println("Error! message")
	// NOTE: `Fatal` or `Panic` instead of `Println`
	// logger.Fatal.Fatal("Quit program")
	// logger.Fatal.Panic("Quit program with traceback")
}

func RandomTypesAndKind() {
	b := true
	s := ""
	n := 1
	f := 1.0
	a := []string{"foo", "bar", "baz"}

	fmt.Println("### TypeOf ###")
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(n))
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(a))

	fmt.Println("### Kind ###")
	fmt.Println(reflect.ValueOf(b).Kind())
	fmt.Println(reflect.ValueOf(s).Kind())
	fmt.Println(reflect.ValueOf(n).Kind())
	fmt.Println(reflect.ValueOf(f).Kind())
	fmt.Println(reflect.ValueOf(a).Kind())
	fmt.Println(reflect.ValueOf(a).Index(0).Kind()) // For slices and strings
}
