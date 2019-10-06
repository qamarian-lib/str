package str

import (
	"fmt"
	"testing"
)

func TestPrintEtr (t *testing.T) {
	fmt.Print ("Testing function PrintEtr ()...\n\n\n")

	PrintEtr ("0 - ground",   "nte", "note")
	PrintEtr (1,              "std", "stnd")
	PrintEtr ("2",            "wrn", "warn")
	PrintEtr (&[]string{"3"}, "err", "err!")
	PrintEtr (true,           "foo", "invl")

	fmt.Println ("\n\nEnd of test!")
}
