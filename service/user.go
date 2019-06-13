package service

import "fmt"

type User struct {
	Context string
}

func (t User) S() {
	fmt.Println(t.Context)
	fmt.Println("Size S")
}
func (t User) M() {
	fmt.Println(t.Context)
	fmt.Println("Size M")
}
func (t User) L() {
	fmt.Println(t.Context)
	fmt.Println("Size L")
}
func (t User) XL() {
	fmt.Println(t.Context)
	fmt.Println("Size XL")
}
func (t User) XXL() {
	fmt.Println(t.Context)
	fmt.Println("Size XXL")
}
