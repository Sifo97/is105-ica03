package iso

import "fmt"

func IterateOverExtendedASCIIStringLiteral(sl string) {	
	for i := 0; i < len(sl); i++ {
	c := []byte {sl[i]}
	fmt.Printf("%X %c %b\n", c, c, sl[i])
	}
	fmt.Println()
}

// Kode for Oppgave 2b
const Greeting = "\x22\x53\x61\x6C\x75\x74\x2c\x20\xC3\xA7\x61\x20\x76\x61\x20\xC2\xB0\x2D\x29\x20\xE2\x82\xAC\x35\x30\x22"
//"Salut, ça va °-) €50" i heksadesimale tall 

func GreetingExtendedASCII() string{
	return Greeting
}