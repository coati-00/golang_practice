package main

import (
    "fmt"
    "strings"
    "strconv"
)


func StrToInt(str string) (int, error) {
    nonFractionalPart := strings.Split(str, ".")
    return strconv.Atoi(nonFractionalPart[0])
}


func main() {
	someString := "1 2 3 4 5 6"
	words := strings.Fields(someString)
	
	var total int;
	var intword int;
	
	for i := 0; i < len(words); i++ {
	    intword, _ = StrToInt(words[i])
	    total = intword + total
	}
	fmt.Println(total)
}
