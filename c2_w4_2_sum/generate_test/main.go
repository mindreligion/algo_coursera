package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// v1
func main1() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	out := ""
	s := time.Now()
	for i := 0; i < 10000; i++ {
		n := int(r.Int31n(2e9+1) - 1e9)
		out += fmt.Sprintf("%d\n", n)
	}
	t := time.Now().Sub(s)
	fmt.Println(out)
	fmt.Println()
	fmt.Println(t / time.Millisecond)
}

// v2
func main2() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pattern := ""
	var nn []interface{}
	s := time.Now()
	for i := 0; i < 100000; i++ {
		n := int(r.Int31n(2e8+1) - 1e8)
		nn = append(nn, n)
		pattern += "%d\n"
	}
	out := fmt.Sprintf(pattern, nn...)
	t := time.Now().Sub(s)
	fmt.Println(out)
	fmt.Println()
	fmt.Println(t / time.Millisecond)
}

// v3
func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var nn []string
	//s := time.Now()
	for i := 0; i < 1000; i++ {
		n := int(r.Int31n(2e9+1) - 1e9)
		nn = append(nn, fmt.Sprint(n))
	}
	out := strings.Join(nn, "\n")
	//t := time.Now().Sub(s)
	if err := ioutil.WriteFile("2_sum_test.txt", []byte(out), 0666); err != nil {
		panic(err)
	}
	//fmt.Println()
	//fmt.Println(t/time.Millisecond)
}
