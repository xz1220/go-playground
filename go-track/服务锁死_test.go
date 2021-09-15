package goTrack

import (
	"fmt"
	"testing"
)

func TestLock(t *testing.T) {
	Lock()
	fmt.Println("Test")
}
