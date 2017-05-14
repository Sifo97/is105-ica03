package main

import ( 
	"fmt"
	"./ascii"
	"./iso"
	//"./fileutils"
	"./unicode"
	//"strconv"
	//"encoding/hex"
	//"strings"
	"os"
 )

func main() {

	laang := os.Args[1]
	
	ascii.IterateOverASCIIStringLiteral(ascii.Ascii)
	
	//d := 
	ascii.GreetingASCII()
	//fmt.Println(d)
	
	var ExtAscii []byte
	
	for i := 0x80; i <= 0xFF; i++ {
		ExtAscii = append (ExtAscii, byte(i))
	}
	
	var ExtAsciiString string = string(ExtAscii)
	
	iso.IterateOverExtendedASCIIStringLiteral(ExtAsciiString)
	a := iso.GreetingExtendedASCII()
	fmt.Println(a)
	
	//byteslice := fileutils.FileToByteslice("treasure/treasure.txt")
	
	//for i := 0, i < len(byteslice); i++ {
	//	fmt.Println(byteslice[i])
	//}
	//fmt.Printf("%s", byteslice)
	//s := fmt.Sprintf("%s", byteslice)
	//s = strings.TrimPrefix(s, "\\x")
	//fmt.Println(s)
	//fmt.Printf("%c", 120)
	
	//dest := make([]byte, len(byteslice))
	//hex.Decode(dest,(byteslice))
	//fmt.Println(dest)
	
	
	unicode.Translate(unicode.Norsk, laang)
	
}