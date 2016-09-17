package debug

import (
	"testing"
	"errors"
)

func TestPrint(t *testing.T) {
	Print("Hello")
	Print("Hello %d %s", 123, "Alice")
	Print(errors.New("some error"))
	
	Logger = nil
	Print("No output")
}
