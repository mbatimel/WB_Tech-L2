package pattern
// Паттерн Factory Method относится к порождающим паттернам уровня класса и сфокусирован
// только на отношениях между классами
// Плюсы: удобно расширяемый, понадобиться только реализизовать классы продуктов и немного изменить фабричный метод
// Минусы: то что объекты придерживаются одного интерфейса и не связаны между собой 

import "fmt"

const (
	ServType     string = "serv"
	NotebookType string = "notebook"
)

type StoreObject interface {
	GetType() string
	PrintDetails()
}

type Notebook struct {
	Display  string
	Keyboard string
	Switch string
}

type Serv struct {
	CPU    string
	Memory int
}

func NewServ() Serv {
	return Serv{
		CPU:    "Intel",
		Memory: 512,
	}
}

func NewNotebook() Notebook {
	return Notebook{
		Display:  "HP",
		Keyboard: "FullSize",
		Switch: "Chinese switches",
	}
}

func (n Notebook) GetType() string {
	return "notebook"
}

func (n Notebook) PrintDetails() {
	fmt.Printf("Display %s, Keyboard %s, Switch %s\n", n.Display, n.Keyboard, n.Switch)
}

func (s Serv) GetType() string {
	return "serv"
}

func (s Serv) PrintDetails() {
	fmt.Printf("CPU %s, Mem %d\n", s.CPU, s.Memory)
}

func New(typeName string) StoreObject {
	switch typeName {
	default:
		fmt.Printf("Несуществующий тип %s\n", typeName)
		return nil
	case ServType:
		return NewServ()
	case NotebookType:
		return NewNotebook()
	}
}
// func main() {
// 	res:=New("notebook")
// 	res.PrintDetails()

// }