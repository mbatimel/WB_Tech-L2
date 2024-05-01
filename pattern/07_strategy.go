package pattern
//Паттерн Strategy относится к поведенческим паттернам уровня объекта.
// Паттерн Strategy определяет набор алгоритмов схожих по роду деятельности,
// инкапсулирует их в отдельный класс и делает их подменяемыми. Паттерн Strategy
// позволяет подменять алгоритмы без участия клиентов, которые используют эти алгоритмы.
// Из минусов => поскольку все поведенческие классы должны быть известны, его нелегко расширить.
// Плюс в том что мы можем легко взаимозаменять алгоримт внутри семейства

// стратегия которая выполняет логику типа оплаты заказов
type Payment interface{
	Pay() error
}

type cardPayment struct{
	cardNumber, cvv string
}
func NewCardPayment(cardNumber string, cvv string) Payment{
	return &cardPayment{
		cardNumber: cardNumber,
		cvv: cvv,
	}
}

func(c *cardPayment) Pay() error {

	// implement
	return nil
}

type moneyPayment struct{
	currency string
}
func NewMoneyPatment(currency string) Payment{
	return &moneyPayment{
		currency: currency,
	}
}
func (m *moneyPayment) Pay() error{
	// implement
	return nil
}
//основаная логика обработки заказа
func processOrder(product string, payment Payment){
	// implement
	//вызов логики типа оплаты заказов
	err:=payment.Pay()
	if err != nil {
		return 
	}

}

// func main() {
// 	product:="geely"
// 	payWay:= 2
// 	var payment Payment
// 	switch payWay{
// 	case 1:
// 		payment = NewCardPayment("123","123")
// 	case 2:
// 		payment = NewMoneyPatment("Rub")

// 	}
// 	processOrder(product, payment)
// }