package debug

import (
	"testing"
)

func TestPrint(t *testing.T) {
	Print("Hello")
	Print("Hello", 123)
	Printf("Hello %d", 123)
}
