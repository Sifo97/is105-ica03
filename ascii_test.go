package main

//import "testing"
import "./ascii"
import "testing"
//var ascii_tests_string = []struct {
//	n1		 string
//	expected string
//}{
//	{"\x22\x48\x65\x6C\x6C\x6F\x20\x3A\x2D\x29\x22"},
//	{"\x22\x48\x65\x6C\x6C\x6F\x20\x3A\x2D\x29\x22"},
//}
//
//func TestGreetingASCII(t *testing.T) {
//	for _, v := range ascii_tests_string {
//		if val := ascii.GreetingASCII(); val != v.expected {
//			t.Errorf("String (%c) returned %c, expected %c", val, v.expected)
//		}
//	}
//}


func TestFatal(t *testing.T) {
	str := "x22\x48\x65\x6C\x6C\x6F\x20\x3A\x2D\x29\x22"
	for i:= 0; i < len(str); i++ {
		for j := 0; i < len(ascii.Ascii); j++ {
			if str[i] != ascii.Ascii[j] {
				t.Error("expected", str[i], "got", ascii.Ascii[j])
			}
		}
	}
}