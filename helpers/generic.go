package helpers

import (
	"math/rand"
	"net"
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

func GetLocalHostname() string {
	name, err := os.Hostname()
	if err != nil {
		return ""
	}
	return name
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func IsStringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func RandomBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()%2 == 0
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
