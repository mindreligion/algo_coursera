package main

import "fmt"

func main() {

	// todo: make test cases from this
	h := heap{}
	fmt.Println("initial ", h)
	h.Put(10)
	fmt.Println("+10 ", h)
	h.Put(7)
	fmt.Println("+7 ", h)
	h.Put(19)
	fmt.Println("+19 ", h)
	min, _ := h.Pop()
	fmt.Println("pop: ", min, ", ", h)
	h.Put(3)
	fmt.Println("+3 ", h)
	h.Put(7)
	fmt.Println("+7 ", h)
	h.Put(5)
	fmt.Println("+5 ", h)
	i := h.Put(1)
	fmt.Println("+1 ", i, " ", h)
	h.Put(4)
	fmt.Println("+4 ", h)
	h.Put(2)
	fmt.Println("+2 ", h)
	h.Put(11)
	fmt.Println("+11 ", h)
	h.Put(16)
	fmt.Println("+16 ", h)
	h.Put(6)
	fmt.Println("+6 ", h)
	min, _ = h.Pop()
	fmt.Println("pop: ", min, ", ", h)
	min, _ = h.Min()
	fmt.Println("min ", min, ", ", h)
	h.Put(7)
	fmt.Println("+7 ", h)
	h.Put(20)
	fmt.Println("+20 ", h)
	h.Remove(3)
	fmt.Println("-3th ", h)
	h[2] = 9
	h[6] = 21
	fmt.Println("2th 3->9; 6th 6->21 ", h)
	h.Remove(5)
	fmt.Println("-5th ", h)
}
