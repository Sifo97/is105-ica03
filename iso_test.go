package main

import ( 
	"./iso"
	"testing"
	"fmt"
)

func isExtendedASCII(s string) bool {
	for _, c := range s {
		if c < 127 {
			return false
		}
	}
	return true
}

func TestGreetingExtendedASCII(t *testing.T) {
	if isExtendedASCII(iso.GreetingExtendedASCII()) == true {
	fmt.Println("Success: Inneholder bare tegn fra en Extended ASCII")
	}else{
		t.Error("Feil: Inneholder normal chars ")
	}
}
	
	