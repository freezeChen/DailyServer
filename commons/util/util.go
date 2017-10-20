package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ToString returns the result converted to type string.
func ToString(value interface{}) string {
	if v, ok := value.([]string); ok {
		if len(v) > 0 {
			return v[0]
		}
		return ""
	} else if v, ok := value.(string); ok {
		return v
	} else if v, ok := value.([]interface{}); ok {
		if len(v) > 0 {
			return fmt.Sprintf("%v", v[0])
		}
		return ""
	}
	return fmt.Sprintf("%v", value)
}

// ToInt returns the result converted to type int.
func ToInt(value interface{}) int {
	if result := ToString(value); result == "" {
		return 0
	} else if i, err := strconv.Atoi(result); err == nil {
		return i
	} else {
		return 0
	}
}

// ToInt64 returns the result converted to type int64.
func ToInt64(value interface{}) int64 {
	if result := ToString(value); result == "" {
		return 0
	} else if i, err := strconv.ParseInt(result, 10, 64); err == nil {
		return i
	} else {
		return 0
	}
}

// ToUint64 returns the result converted to type uint64.
func ToUint64(value interface{}) uint64 {
	if result := ToString(value); result == "" {
		return 0
	} else if i, err := strconv.ParseUint(result, 10, 64); err == nil {
		return i
	} else {
		return 0
	}
}

// ToFloat64 returns the result converted to type float64.
func ToFloat64(value interface{}) float64 {
	if result := ToString(value); result == "" {
		return 0
	} else if i, err := strconv.ParseFloat(result, 64); err == nil {
		return i
	} else {
		return 0
	}
}

// ToBool returns the boolean value represented by the string.
func ToBool(value string) bool {
	if value == "" {
		return false
	}
	val, _ := strconv.ParseBool(value)
	return val
}

// GetRemoteIp returns the server ip.
func GetRemoteIp(r *http.Request) (ip string) {
	ip = r.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = r.RemoteAddr
	}
	ip = strings.Split(ip, ":")[0]
	if len(ip) < 7 || ip == "127.0.0.1" {
		ip = "localhost"
	}
	return
}

// RandInt64 returns a random number between two numbers.
func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

// Random returns random string.
// returns pure number if the value of kind is 0.
// returns small letters if the value of kind is 1.
// returns capital letters if the value of kind is 2.
// returns number or letters if the value of kind is 3.
func Random(size, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all {
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

// QuickValidate returns whether the express matches the string value.
func QuickValidate(express, value string) bool {
	if len(value) == 0 || len(express) == 0 {
		return false
	}
	myRegex, _ := regexp.Compile(express)
	return myRegex.MatchString(value)
}

// IsEmpty returns true if the string is empty
func IsEmpty(value string) bool {
	if len(value) == 0 || value == "&nbsp;" {
		return true
	}
	return false
}

// IsNumber retunrs true if the string is a pure number.
func IsNumber(value string) bool {
	return QuickValidate("^[1-9]*[0-9]*$", value)
}

// IsMobile retunrs true if the string is a mobile format.
func IsMobile(value string) bool {
	return QuickValidate("^(1[3-9])\\d{9}$", value)
}

// IsEmail retunrs true if the string is a email format.
func IsEmail(value string) bool {
	return QuickValidate(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`, value)
}

// IsCurrency retunrs true if the string is a currency format.
func IsCurrency(value string) bool {
	return QuickValidate("^[1-9]*[0-9].*$", value)
}

// If retunrs trueVal if the condition is correct.
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// GetMd5String retunrs a 32-bit md5 string.
func GetMd5String(value string) string {
	h := md5.New()
	h.Write([]byte(value))
	return hex.EncodeToString(h.Sum(nil))
}

// PadLeft make up number of bits.(left)
func PadLeft(value string, totalWidth int, paddingChar byte) string {
	length := len(value)
	if length >= totalWidth {
		return value
	}
	var retStr string
	num := totalWidth - length
	for i := 1; i < num; i++ {
		retStr = retStr + string(paddingChar)
	}
	retStr = retStr + value
	return retStr
}

// GetCurrentDirectory returns the current path.
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return dir
}
