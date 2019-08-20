package str

import (
	"fmt"
	"testing"
)

func TestPrintEtr (t *testing.T) {
	fmt.Println ("Testing function PrintEtr ()...")
	PrintEtr (1,              "std", "some")
	PrintEtr ("2",            "wrn", "tome")
	PrintEtr (&[]string{"3"}, "err", "uome")
	PrintEtr (true,           "foo", "vome")
	fmt.Println ("End of test!")
}
