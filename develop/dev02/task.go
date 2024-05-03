package main

import (
	"fmt"
	"unicode"
)

var ErrorString = fmt.Errorf("invalid string")


func UnpackingString(s string) (string,  error) {
	sr := []rune(s)
	sleshString:=false
	doublesleshString :=false

	var result string
	for itter, item := range sr {
		if unicode.IsDigit(item) && itter == 0 {
			return "", ErrorString
		}
		if unicode.IsLetter(item){
			result += string(item)
		}
		if unicode.IsDigit(item) && unicode.IsLetter(sr[itter-1]){
			
			for i := 0; i < int(item - '0')-1; i++ {
				result += string(sr[itter-1])
			}
			continue
		}
		if string(item) == `\` && !sleshString && !doublesleshString {
			sleshString = true
			continue
		}
		if unicode.IsLetter(item) && sleshString{
			result += string(item)
			sleshString = false
			continue
		}
		if string(item) == `\` && !doublesleshString && sleshString{
			result += string(item)
			doublesleshString = true
		}
		if unicode.IsDigit(item) && doublesleshString{
			for i := 0; i < int(item - '0')-1; i++ {
				result += string(sr[itter-1])
			}
			doublesleshString = false
			continue
		}
		if unicode.IsDigit(item) && sleshString{
			result += string(item)
			sleshString = false
			continue
		}
		if unicode.IsDigit(item) && unicode.IsDigit(sr[itter-1]){
			
			for i := 0; i < int(item - '0')-1; i++ {
				result += string(sr[itter-1])
			}
			continue
		}	
	}

	return result, nil
}
func main() {
	result, err := UnpackingString(`qwe\\5`)
	if err != nil {
		panic(err)
	}else{
		fmt.Printf("Распакованая строка: %s\n", result)
	}
}