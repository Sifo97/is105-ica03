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
	fmt.Println(r)
}