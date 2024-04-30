package pattern

// Шаблон посетителя позволяет добавлять поведение к структуре, не изменяя ее.
// Это поведенческий шаблон объектного уровня.
// Минусы:
// Добавление новых классов может быть затруднено, поскольку требует обновления иерархии посетителей и ее подклассов.
// Плюсы:
// Класс Visitor может сохранять некоторое состояние при перемещении по контейнеру.
// Связанные операции могут быть сгруппированы в классе Visitor.

type Visitor interface {
	visitDataBaseServer(*DataBaseServer)
	visitBackServer(*BackServer)
	visitFrontServer(*FrontServer)
}

type ServerInfo interface {
	GetType() string
	Accept(Visitor)
}

type DataBaseServer struct {
	workStatus string
}

func NewDataBaseServer(workStatus string) *DataBaseServer {
	return &DataBaseServer{
		workStatus: workStatus,
	}
}

func (s *DataBaseServer) GetType() string {
	return "It is DataBase Server"
}

func (s *DataBaseServer) Accept(v Visitor) {
	v.visitDataBaseServer(s)
}

type BackServer struct {
	workStatus string
}

func NewBackServer(workStatus string) *BackServer {
	return &BackServer{
		workStatus: workStatus,
	}
}

func (b *BackServer) GetType() string {
	return "It is Back server"
}

func (b *BackServer) Accept(v Visitor) {
	v.visitBackServer(b)
}

type FrontServer struct {
	workStatus string
}

func NewFrontServer(workStatus string) *FrontServer {
	return &FrontServer{
		workStatus: workStatus,
	}
}

func (b *FrontServer) GetType() string {
	return "It is Front server"
}

func (b *FrontServer) Accept(v Visitor) {
	v.visitFrontServer(b)
}

type ServerStatus struct {
	newStatus string
}

func NewServerStatus(newStatus string) *ServerStatus {
	return &ServerStatus{
		newStatus: newStatus,
	}
}

func (s *ServerStatus) visitDataBaseServer(d *DataBaseServer) {
	d.workStatus = s.newStatus
}

func (s *ServerStatus) visitBackServer(b *BackServer) {
	b.workStatus = s.newStatus
}

func (s *ServerStatus) visitFrontServer(f *FrontServer) {
	f.workStatus = s.newStatus
}

// func main() {
// 	serverDB := NewDataBaseServer("off")
// 	serverBack := NewBackServer("off")
// 	serverFront := NewFrontServer("off")

// 	newStatusServers := NewServerStatus("on")

// 	serverBack.Accept(newStatusServers)
// 	serverDB.Accept(newStatusServers)
// 	serverFront.Accept(newStatusServers)

// 	fmt.Println(serverDB.workStatus, serverBack.workStatus, serverFront.workStatus)
// }
