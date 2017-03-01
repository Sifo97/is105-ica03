package main

import ( 
	"./ascii"
	"testing"
	"fmt"
)

func isASCII(s string) bool {
	for _, c := range s {
		if c > 127 {
			return false
		}
	}
	return true
}

func TestGreetingASCII(t *testing.T) {
	if isASCII(ascii.GreetingASCII()) == true {
	fmt.Println("Success: Inneholder bare ASCII chars")
	}else{
	t.Error("Feil: Inneholder normal chars ")
	}
	}
	