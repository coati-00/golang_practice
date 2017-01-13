package main

import "fmt"
import "strings"
import "strconv"


func StrToInt(str string) (int, error) {
    nonFractionalPart := strings.Split(str, ".")
    return strconv.Atoi(nonFractionalPart[0])
}


func main() {
	someString := "1 2 3 4 5 6"
	
	words := strings.Fields(someString)
	
	var total int;
	var intword int;

	fmt.Println(words, len(words))
	
	for i := 0; i < len(words); i++ {
	    fmt.Println(words[i])
	    //var intword int = int(words[i])
	    intword, _ = StrToInt(words[i])
	    fmt.Println(intword)
	    total = intword + total
	}
	fmt.Println(total)
}
