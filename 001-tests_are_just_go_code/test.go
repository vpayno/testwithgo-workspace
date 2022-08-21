package main

import (
	"fmt"
	"runtime"

	"github.com/vpayno/testwithgo-workspace/001-tests_are_just_go_code/math"
)

type test func() bool

func main() {
	var pass int

	tests := []test{
		testSum,
		testAdd,
	}

	for _, f := range tests {
		if f() {
			pass++
		}
	}

	fail := len(tests) - pass

	fmt.Println()
	if fail > 0 {
		fmt.Printf("Tests that passed: %d\n", pass)
		fmt.Printf("Tests that failed: %d\n", fail)
	}
	fmt.Printf("Score: %d%%\n", pass*100/len(tests)*100/100) // LOL
}

func testSum() (Ok bool) {
	pc, _, _, _ := runtime.Caller(0)
	name := runtime.FuncForPC(pc).Name() + "()"

	Ok = true
	want := 11
	got := math.Sum([]int{10, -2, 3})

	if want != got {
		Ok = false
		msg := fmt.Sprintf("%s FAIL: Wanted %d, got %d", name, want, got)
		fmt.Println(msg)
		return
	}

	fmt.Printf("%s PASS\n", name)
	return
}

func testAdd() (Ok bool) {
	pc, _, _, _ := runtime.Caller(0)
	name := runtime.FuncForPC(pc).Name() + "()"

	Ok = true
	want := 5
	got := math.Add(2, 3)

	if want != got {
		Ok = false
		msg := fmt.Sprintf("%s FAIL: Wanted %d, got %d", name, want, got)
		fmt.Println(msg)
		return
	}

	fmt.Printf("%s PASS\n", name)
	return
}
