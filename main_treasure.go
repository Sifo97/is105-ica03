package main
import (
	"fmt"
	"./fileutils"
	"./treasure"
)

func main() {
	var fileToOpen = "treasure/treasure.txt"
	
	c := fileutils.FileToByteslice(fileToOpen)
	s := string(c[:len(c)])
	r := treasure.PrintTreasureUTF8(s)
	fmt.Print(r)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c", r[i])
	}
	}