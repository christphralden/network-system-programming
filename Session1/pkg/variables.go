package pkg

import (
	"fmt"
	"unicode/utf8"
)

func RunVariables(){
	// datatypes
	var num int // set default = 0
	fmt.Println(num)

	var num64 int64 //allocate the bit
	fmt.Println(num64)

	var dec float32 = 123456.9
	fmt.Println(dec)

	var dec64 float64 = 123456.9
	fmt.Println(dec64)

	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)

	fmt.Println(len(myString)) // byte of string, not number of character
	//GO uses UTF-8, not vanilla ascii bakal beda

	fmt.Println(utf8.RuneCountInString(myString)) // accurate string ln

	var myRune rune = 'a' //unicode point, krn GO pake UTF-8,seperti char
	fmt.Println(myRune)

	var myBool bool = false
	fmt.Println(myBool)

	// initialization
	// ada default value
	// var myVar = "ini string"
	// myVar := "ini string"
	// var myVar, myVar2 int = 2, 2
	// myVar, myVar2 := 2,2
	// myVar := someFucntion() -> lebih baik dikasi type biar jelas
	// const myConst string ="const Value" -> when creating also init

	// fmt.Println(myVar)

}