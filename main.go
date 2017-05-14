package main

import ( 
	"fmt"
	"./ascii"
	"./iso"
	"./unicode"
	"os"
 )

func main() {

	laang := os.Args[1] //Valg av hvilket språk som skal oversettes til
	
	//Oppgave 1a
	ascii.IterateOverASCIIStringLiteral(ascii.Ascii)
	
	//Oppgave 1b
	ascii.GreetingASCII()
	
	//Oppgave 2a
	var ExtAscii []byte
	
	for i := 0x80; i <= 0xFF; i++ {
		ExtAscii = append (ExtAscii, byte(i))
	}
	
	var ExtAsciiString string = string(ExtAscii)
	
	iso.IterateOverExtendedASCIIStringLiteral(ExtAsciiString)
	
	//Oppgave 2b
	a := iso.GreetingExtendedASCII()
	fmt.Println(a)
	
	//Oppgave 4a
	unicode.Translate(unicode.Norsk, laang)
	
}