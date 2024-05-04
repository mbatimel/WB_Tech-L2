package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/pborman/getopt"
)

type SortFile struct {
	filename string
	flagK    int
	flagN    bool
	flagR    bool
	flagU    bool
	flagM    bool
	flagB    bool
	flagC    bool
	flagH    bool
	data 	[]string
}

func NewSortFile(filename string, flagK int, flagN, flagR, flagU, flagM, flagB, flagC, flagH bool) *SortFile {
	return &SortFile{
		filename: filename,
		flagK:    flagK,
		flagN:    flagN,
		flagR:    flagR,
		flagU:    flagU,
		flagM:    flagM,
		flagB:    flagB,
		flagC:    flagC,
		flagH:    flagH,
		data: make([]string,0),
	}
}
func (s *SortFile) TextFromFile() (error)  {
	if s.filename == ""{
		return errors.New("file not found")
	}
	file, err := os.Open(s.filename)
	if err != nil {
		return err
	}
	defer file.Close() 
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan(){
		
		s.data = append(s.data, scanner.Text())
		
    }
	return nil

}
func (s *SortFile)SortText() ([]string, error) {
	if s.flagU{
		s.data = MakeUnique(s.data)
	}

	sort.Slice(s.data, func(i, j int) bool {
		lineA := s.data[i]
		lineB := s.data[j]

		if s.flagB {
			lineA = strings.TrimRight(lineA, " ")
			lineB = strings.TrimRight(lineB, " ")
		}

		if s.flagK > 0 && s.flagK <= len(strings.Fields(lineA)) && s.flagK <= len(strings.Fields(lineB)) {
			// Обрабатываем случай с указанием колонки
			fieldA := strings.Fields(lineA)[s.flagK-1]
			fieldB := strings.Fields(lineB)[s.flagK-1]

			if s.flagH {
				// Извлекаем числовое значение и суффикс из поля
				numA, suffA := extractNumericSuffix(fieldA)
				numB, suffB := extractNumericSuffix(fieldB)

				// Сортируем по числовому значению и сравниваем суффиксы
				if numA != numB {
					return numA < numB
				}
				return suffA < suffB
			}

			if s.flagN {
				// Сортировка по числовому значению
				numA, errA := strconv.Atoi(fieldA)
				numB, errB := strconv.Atoi(fieldB)

				if errA == nil && errB == nil {
					return numA < numB
				}
			}

			// Сортировка по строковому значению колонки
			return fieldA < fieldB
		}

		if s.flagM {
			
			dateA, errA := time.Parse("Jan", lineA)
			dateB, errB := time.Parse("Jan", lineB)

			if errA == nil && errB == nil {
				return dateA.Before(dateB)
			}
		}

		// Сортировка по всей строке
		return lineA < lineB
	})

	if s.flagR{
		for i, j := 0, len(s.data)-1; i < j; i, j = i+1, j-1 {
			s.data[i], s.data[j] = s.data[j], s.data[i]
		}
	}
	if s.flagC && isSorted(s.data){
		fmt.Println("Sorting data complete")
	}

	return s.data,nil
}

func (s *SortFile) WriteSortedStringToFile(sortedString []string) error {
	f, err := os.Create("sorted_"+s.filename)
    if err != nil { return err }

    defer f.Close()
    for _, word := range sortedString {
        _, err := f.WriteString(word + "\n")
        if err != nil { return err}
    }

    
	return nil
}

func main() {
	filename := getopt.String('f', "", "Путь до файла")
	flagK := getopt.Int('k', 0, "номер колонки для сортировки (по умолчанию 0, разделитель - пробел)")
	flagN := getopt.Bool('n', "сортировать по числовому значению")
	flagR := getopt.Bool('r', "сортировать в обратном порядке")
	flagU := getopt.Bool('u', "не выводить повторяющиеся строки")
	flagM := getopt.Bool('M', "сортировать по названию месяца")
	flagB := getopt.Bool('b', "игнорировать хвостовые пробелы")
	flagC := getopt.Bool('c', "проверять отсортированы ли данные")
	flagH := getopt.Bool('h', "сортировать по числовому значению с учетом суффиксов")
	getopt.Parse()

	sf := NewSortFile(*filename, *flagK, *flagN, *flagR, *flagU, *flagM, *flagB, *flagC, *flagH)
	err:= sf.TextFromFile()
	if err != nil {
		panic(err)
	}
	sortedString, err := sf.SortText()
	if err != nil {
		panic(err)
	}
	err = sf.WriteSortedStringToFile(sortedString)
	if err != nil{
		panic(err)
	}

}
func MakeUnique(strs []string) []string {
	unique := make(map[string]struct{})
	newStrs := []string{}
	for _, str := range strs {
		if _, ok := unique[str]; !ok {
			unique[str] = struct{}{}
			newStrs = append(newStrs, str)
		}
	}
	return newStrs
}
func extractNumericSuffix(s string) (int, string) {
	numericPart := ""
	suffixPart := ""

	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			numericPart = string(s[i]) + numericPart
		} else {
			suffixPart = string(s[i]) + suffixPart
		}
	}

	num, _ := strconv.Atoi(numericPart)

	return num, suffixPart
}
func isSorted(lines []string) bool {
	for i := 1; i < len(lines); i++ {
		if lines[i-1] > lines[i] {
			return false
		}
	}
	return true
}