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
const Greeting = "\x22\x53\x61\x6c\x75\x74\x2c\x20\xE7\x61\x20\x76\x61\x20\xB0\x96\x29\x20\x80\x35\x30\x22"
//"Salut, ça va °-) €50" i heksadesimale tall 

func GreetingExtendedASCII() string{

	for i := 0; i < len(Greeting); i++ {
		fmt.Printf("%c", Greeting[i])
	}
	return Greeting
}