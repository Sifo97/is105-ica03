package iso

import "fmt"

func IterateOverExtendedASCIIStringLiteral(sl string) {	
	for i := 0; i < len(sl); i++ {
	c := []byte {sl[i]}
	fmt.Printf("%X %s %b\n", c, c, sl[i])
	}
}

// Kode for Oppgave 2b
func GreetingExtendedASCII() {
	//greeting := "\x53\x67"
}