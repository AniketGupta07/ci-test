package mypack

import (
	"strconv"
	"strings"
)

func ips(ip string) bool {
	var i = 0
	var ipSlice = strings.Split(ip, ".")
	if len(ipSlice) == 4 {
		for i < 4 {
			var ip_piece, _ = strconv.Atoi(ipSlice[i])
			if !(ip_piece >= 0 && ip_piece <= 255) {
				// fmt.Println("Err")
				return false
			}
			i++
		}
	} else {
		// fmt.Println("Err")
		return false
	}
	return true
}

func clean(str string) string {

	s := make([]string, 0)
	l := len(str)
	for i := 0; i < l; i++ {
		if (str[i] <= 'z' && str[i] >= 'a') || ((str[i] >= 'Z' && str[i] <= 'A') || (str[i] >= '9' && str[i] <= '0')) {
			s = append(s, string(str[i]))
			l--
		}
	}
	s2 := str[:l+1]
	return s2

}

// func main() {
// 	var reader string
// 	fmt.Scanf("%s", &reader)
// 	c := ips(reader)
// 	if c {
// 		fmt.Println("VALID")
// 	}
// 	if !c {
// 		fmt.Println("INVALID IP")
// 	}
// 	var st string

// 	fmt.Scanf("%s", &st)
// 	str := clean(st)
// 	fmt.Println(str)
// }
