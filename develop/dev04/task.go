package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	res:= Annogramsearch([]string{"тяпка", "пятак", "пятка", "листок", "слиток", "столик"})
	for key, val := range res{
		fmt.Println(key, val)
	}
}
func Annogramsearch(data []string) map[string] []string {
	result := make(map[string][]string)
out:	
	for _, v := range data{
		lowerString := strings.ToLower(v)
		sortedString := SortString(lowerString)
		for key, value := range result{
			sortedKey := SortString(strings.ToLower(key))
			if sortedKey ==sortedString{
				for _, value := range value {
					if lowerString == strings.ToLower(value) {
						continue out 
					}
				}
				result[key] = append(result[key], v)
				continue out
			}
			}
			result[v] = []string{v}
		}
		for key, value := range result {
			if len(value) < 2 {
				delete(result, key)
			}
		}
	
	return result
}
func SortString(str string) string {
	sortedRunes := []rune(str)
	sort.Slice(sortedRunes, func(i, j int) bool {
		return sortedRunes[i] < sortedRunes[j]
	})
	return string(sortedRunes)
}