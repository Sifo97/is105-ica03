package main

import ( 
	"fmt"
	"./ascii"
	"./iso"
 )

func main() {
	ascii.IterateOverASCIIStringLiteral(ascii.Ascii)
	
	s := ascii.GreetingASCII()
	fmt.Println(s)
	
	var ExtAscii []byte
	
	for i := 0x80; i <= 0xFF; i++ {
		ExtAscii = append (ExtAscii, byte(i))
	}
	
	var ExtAsciiString string = string(ExtAscii)
	
	iso.IterateOverExtendedASCIIStringLiteral(ExtAsciiString)
	
	a := iso.GreetingExtendedASCII()
	fmt.Println(a)
	
}