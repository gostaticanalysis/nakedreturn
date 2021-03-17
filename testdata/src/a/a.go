package a

func f1() (n int) {
	return // want "should not use naked return"
}

func f2() (n int) {
	return 10 // OK
}

func f3() {
	return // OK
}

func f4() (n int, _ int) {
	return // want "should not use naked return"
}

func f5() (n int, _ struct{}) {
	return // want "should not use naked return"
}

func f6() (n int, _ string) {
	return // want "should not use naked return"
}

func f7() (n int, _ interface{}) {
	return // want "should not use naked return"
}

func f8() (n int, _ error) {
	return // want "should not use naked return"
}

func f9() (n int, _ *int) {
	return // want "should not use naked return"
}

func f10() (n int, _ <-chan int) {
	return // want "should not use naked return"
}

func f11() (n int, _ [1]int) {
	return // want "should not use naked return"
}
