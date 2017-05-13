package treasure

import (
	"encoding/hex"
)
// Kode for Oppgave 3c
// Bruk strengen fra filen treasure.txt som in-data for denne funksjonen
func PrintTreasureUTF8(treasure_string string) string {
	tr, err := hex.DecodeString(treasure_string)
	if err != nil {
	panic(err)
	}
	return string(tr) // returverdien er her kun en stedsholder
}
