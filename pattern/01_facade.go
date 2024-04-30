package pattern

import (
	"fmt"
	"time"
)

/* Шаблон предназначен для того, чтобы скрыть сложности базовой системы и предоставить клиенту простой интерфейс
Минусы: сложный в разработке, есть вероятность превращения в божественный объект(супер-класс)
Плюсы: простой в использовании интерфейс*/

type ComputerSystemFacade interface{
	InitSystem()
	SystemUp()
	SystemDown()
	CheckStatus()
}
type Laptop struct{
	*Power
	*Keyboard
	*Screen
}

func (l *Laptop)SystemUp() {
	l.Keyboard.Start()
	l.Power.Start()
	l.Screen.Start()
	fmt.Println("System up")
}
func (l *Laptop)SystemDown() {
	l.Keyboard.Stop()
	l.Power.Stop()
	l.Screen.Stop()
	fmt.Println("System Down")
}
func (l *Laptop)CheckStatus(){
	if !l.Keyboard.status && !l.Screen.status && !l.Power.status{
		fmt.Println("Systeam status : off")
	} else{
		fmt.Println("Systeam status : on")
	}
}
func InitSystem() *Laptop{
	return &Laptop{
		Power: &Power{status: false},
		Keyboard: &Keyboard{status: false},
		Screen: &Screen{status: false},
	}
}
type StatusSystem interface {
	Start()
	Stop()
}

type Keyboard struct{
	status bool
}
type Screen struct{
	status bool
}
type Power struct{
	status bool
}


func (k *Keyboard) Start() {
	fmt.Println("The keyboard is turn on")
	time.Sleep(1 * time.Second)
	k.status = true
}

func (k *Keyboard) Stop() {
	fmt.Println("The keyboard is turn off")
	time.Sleep(1 * time.Second)
	k.status = false
}
func (s *Screen) Start() {
	fmt.Println("The screen is turn on")
	time.Sleep(1 * time.Second)
	s.status = true
}

func (s *Screen) Stop() {
	fmt.Println("The screen is turn off")
	time.Sleep(1 * time.Second)
	s.status = false
}
func (p *Power) Start() {
	fmt.Println("The power is turn on")
	time.Sleep(1 * time.Second)
	p.status = true
}

func (p *Power) Stop() {
	fmt.Println("The power is turn off")
	time.Sleep(1 * time.Second)
	p.status = false
}
