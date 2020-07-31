package uberGoGuide

import (
	"fmt"
	"testing"
)

func TestPointToInterface(t *testing.T) {
	var f1 S1 = S1{}
	var f2 S2 = S2{}
	f1.f()
	f2.f()
	fmt.Println("f1 can't change the value:", f1.Value)
	fmt.Println("f2 can change the value:", f2.Value)
}
