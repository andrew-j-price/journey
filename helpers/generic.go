package helpers

import (
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/andrew-j-price/journey/logger"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func IsStringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func RandomNumberInRange(minNum int, maxNum int) int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(maxNum-minNum) + minNum
	// fmt.Printf("Random number: %v\n", num)
	return num
}

// Takes a string and returns an int.
// Params (string) to convert, (bool) if conversion fails, should program exit
// Returns (int) of string
func StringToInt(theString string, quit bool) int {
	theInt, err := strconv.Atoi(theString)
	if err != nil {
		logger.Error.Println(err)
		logger.Debug.Printf("During processing of: %v which returned: %v with type: %v", theString, theInt, reflect.TypeOf(theInt))
		if quit {
			logger.Fatal.Panic("Conversion failed. Exiting.")
		}
	}
	// NOTE: returns 0 if conversion failed
	return theInt
}

func TimeNow() string {
	return time.Now().Format(time.RFC3339)
}
