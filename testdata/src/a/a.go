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
