package pattern

// import "fmt"

// Паттерн Chain Of Responsibility позволяет избежать привязки объекта-отправителя запроса к
// объекту-получателю запроса, при этом давая шанс обработать этот запрос нескольким
// объектам. Получатели связываются в цепочку, и запрос передается по цепочке, пока не
// будет обработан каким-то объектом.
// Минусы: нет
// Плюсы: Простота добавления нового обьекта для обработки

type HandlerCarServer interface{
	SendRequest(request int) string
}
type CashierHandler struct {
	next HandlerCarServer
}

func (c *CashierHandler) SendRequest(request int) (result string){
	if request == 1{
		result = "i'm cashier"
	} else if c.next != nil{
		result = c.next.SendRequest(request)
	}
	return
}

type MasterHandler struct {
	next HandlerCarServer
}

func (m *MasterHandler) SendRequest(request int) (result string){
	if request == 2{
		result = "i'm master"
	} else if m.next != nil{
		result = m.next.SendRequest(request)
	}
	return
}

type ClientHandler struct {
	next HandlerCarServer
}

func (c *ClientHandler) SendRequest(request int) (result string){
	if request == 3{
		result = "i'm client"
	} else if c.next != nil{
		result = c.next.SendRequest(request)
	}
	return
}

// func main(){
// 	handler:= &CashierHandler{
// 		next: &MasterHandler{
// 			next: &ClientHandler{},
// 		},
// 	}
// 	result:=handler.SendRequest(2)
// 	fmt.Println(result)
// }