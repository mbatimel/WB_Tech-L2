package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/pborman/getopt"
)

type CustomCut struct {
	fields		string
	delimiter	string
	separated    bool
}

func NewCustomCut(fields string,delimiter string,separated bool) *CustomCut {
	return &CustomCut{
		fields: fields,
		delimiter: delimiter,
		separated: separated,
	}
}
func parseFields(fields string) ([]int, error){
	fields = strings.TrimSpace(fields)
	intStrList := strings.Split(fields, ",")
	res := make([]int, len(intStrList))
	for ind, intStr := range intStrList {
		num, err := strconv.Atoi(intStr)
		if err != nil {
			return nil, errors.New("bad field list")
		}
		res[ind] = num
	}
	return res, nil
}
func (c *CustomCut) CustomCutFunc(str string) string{
	var result string
	colums := strings.Split(str,c.delimiter)
	if len(colums) == 1 && !c.separated {
		result =str
	}else if len(colums) > 1 && c.separated{
		fieldsINT, err := parseFields(c.fields)
		if err != nil {
			panic(err)
		}
		for _, colid :=range fieldsINT{
				result += colums[colid-1] + " "
		}
	}
	return result
}
func (c *CustomCut) readTeaxtFromSTDIN(){
	reader:= bufio.NewReader(os.Stdin)
	for{
		str, err := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		if err == io.EOF {
			break
		}
		fmt.Print(c.CustomCutFunc(str))
		fmt.Println()
	}
}
func main() {
	fields := getopt.StringLong("fields",'f', "" ,"выбрать поля (колонки)")
	delimiter := getopt.StringLong("delimiter",'d',"\t","использовать другой разделитель")
	separated := getopt.BoolLong("separated",'s',"только строки с разделителем")
	getopt.Parse()
	
	sf := NewCustomCut(*fields, *delimiter, *separated)
	sf.readTeaxtFromSTDIN()

}

