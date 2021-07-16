package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet //匿名嵌套类型
}

//重载
func (d *Dog) Speak() {
	fmt.Println("Wang!")
}

// func (d *Dog) SpeakTo(host string) {
// 	d.Speak()
// 	fmt.Println(" ", host)
// }

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.Speak()
	dog.SpeakTo("Chao") //未继承重载
}
