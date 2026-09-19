package main

func literal() {
	integer2 := 0b01
	integer8 := 0o6
	integer16 := 0x8

	string1 := "あいうえお"
	string2 := `あいうえお\nかきくけこ`
	rune1 := '\141'

	println(integer2)
	println(integer8)
	println(integer16)
	println(string1)
	println(string2)
	println(rune1)
}

func main() {
	literal()
}
