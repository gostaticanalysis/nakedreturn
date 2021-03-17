package a

func f() (n int) {
	return // want "should not use naked return"
}

func g() (n int) {
	return 10 // OK
}
