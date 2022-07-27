package helpers

import (
	"encoding/json"
	"flag"
	"math/rand"
	"net"
	"os"
	"os/user"
	"reflect"
	"strconv"
	"time"

	"github.com/andrew-j-price/journey/logger"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetLocalHostname() string {
	name, err := os.Hostname()
	if err != nil {
		return "N/A"
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

func GetLocalUsername() string {
	user, err := user.Current()
	if err != nil {
		return "N/A"
	}
	return user.Username
}

// https://stackoverflow.com/a/54747682/17207113
func IsFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func IsStringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// consume like: fmt.Println(helpers.PrettyPrintStruct(response))
func PrettyPrintStruct(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func RandomBool() bool {
	return rand.Int()%2 == 0
}

func RandomNumberInRange(minNum int, maxNum int) int {
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
