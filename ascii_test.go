package main

//import "testing"
//import "strings"
//import "./ascii"
import "testing"
import "fmt"
func isASCII(s string) bool {
	for _, c := range s {
		if c > 127 {
			return false
		}
	}
	return true
}

func TestGreetingASCII(t *testing.T) {
	if isASCII("a") == false {
	fmt.Println("Success: Inneholder bare ASCII chars")
	}else{
	t.Error("Feil: Inneholder normal chars ")
	}
	}
	


//func TestGreetingsASCII(t *testing.T) {
//    for_, v := range ascii_tests_string {
 //   if !strings.Contains(Ascii, strings.ToLower(string(v))) {
//	return false
//	}
 //    }
 // return true
//}


//func TestFatal(t *testing.T) {
//	str := "x22\x48\x65\x6C\x6C\x6F\x20\x3A\x2D\x29\x22"
//	for i:= 0; i < len(str); i++ {
//		for j := 0; i < len(ascii.Ascii); j++ {
//			if str[i] != ascii.Ascii[j] {
//				t.Error("expected", str[i], "got", ascii.Ascii[j])
//			}
//		}
//	}
//}