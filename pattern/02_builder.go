package pattern
//Шаблон "Строитель" - это порождающий шаблон проектирования, используемый для создания сложных объектов,
// Используется когда создаваемый объект очень большой и состоит из нескольких стадий. Это помогает уменьшить размер конструктора.
//	Когда необходимо создать другую версию того же объект
// Когда не может существовать частично инициализированного объекта.

// Плюсы:
//позволяет изменять внутреннее представление продукта;
//изолирует код, реализующий конструирование и представление;
//дает более тонкий контроль над процессом конструирования.

// Минусы:
// алгоритм создания сложного объекта не должен зависеть от того, из каких частей состоит объект и как они стыкуются между собой;
// процесс конструирования должен обеспечивать различные представления конструируемого объекта.
// усложнен из за большого количества структур

type MarketPalceUserbulder interface {
	SetUserType(selectedUserType string) MarketPalceUserbulder
	SetMarketPalceType(selectedMarketPalceType string) MarketPalceUserbulder
	Build() *MarketPlaceAndUser
}
type MarketPlaceAndUser struct {
	userType        string
	marketplaceType string
}

type marketPalceUserbulder struct {
	mpu *MarketPlaceAndUser
}

// Build implements MarketPalceUserbulder.
func (m *marketPalceUserbulder) Build() *MarketPlaceAndUser {
	return m.mpu
}


// SetMarketPalceType implements MarketPalceUserbulder.
func (m *marketPalceUserbulder) SetMarketPalceType(selectedMarketPalceType string) MarketPalceUserbulder {
	m.mpu.marketplaceType = selectedMarketPalceType
	return m
}

// SetUserType implements MarketPalceUserbulder.
func (m *marketPalceUserbulder) SetUserType(selectedUserType string) MarketPalceUserbulder {
	m.mpu.userType = selectedUserType
	return m
}

func NewMarketplaceUser() MarketPalceUserbulder {
	return &marketPalceUserbulder{
		mpu: &MarketPlaceAndUser{},
	}
}

type Director struct {
	builder MarketPalceUserbulder
}

func (d *Director) Constructor(selectedMarketPalceType,selectedUserType string) *MarketPlaceAndUser{
	d.builder.SetMarketPalceType(selectedMarketPalceType)
	d.builder.SetUserType(selectedUserType)
	return d.builder.Build()
}
// func main(){
// 	build :=NewMarketplaceUser()
// 	director := &Director{builder: build}
// 	newmpu := director.Constructor("WB", "seller")
// 	fmt.Printf("MarketplaceType: %s\n UserType: %s\n",newmpu.marketplaceType, newmpu.userType)

// }