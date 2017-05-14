package main
import (
	"fmt"
	"./fileutils"
	//"./treasure"
)

func main() {
	var fileToOpen = "treasure/treasure.txt"
	c := fileutils.FileToByteslice(fileToOpen)
	fmt.Printf("%c", c)
	for i := 0; i < len(c); i++ {
		fmt.Printf("%c", c[i])
	}
	}