package unicode

import "fmt"

const Norsk = "\x22\x6e\x6f\x72\x64\x20\x6f\x67\x20\x73\xc3\xb8\x72\x22"
//"nord og sør" i heksadesimale tall

const Islandsk = "\x22\x6e\x6f\x72\xc3\xb0\x75\x72\x20\x6f\x67\x20\x73\x75\xc3\xb0\x75\x72\x22"
//“norður og suður” i heksadesimale tall

const Japansk = "\x22\xe5\x8c\x97\xe3\x81\xa8\xe5\x8d\x97\x22"
//“北と南” i heksadesimale tall

// Kode for Oppgave 4a
func Translate(expression string, language string) string {
	if language == "jp" {
		fmt.Printf("%s\n", expression)
		fmt.Println("    =")
		fmt.Printf("%s\n", Japansk)
	}
	if language == "is" {
		fmt.Printf("%s\n", expression)
		fmt.Println("       	=")
		fmt.Printf("%s\n", Islandsk)
	}
	return ""
}
