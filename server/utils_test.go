package server

import (
	"os"
	"testing"
)

func TestCreateAndRemovePID(t *testing.T) {
	os.RemoveAll("./test-util.pid")

	pid := NewPID("./test-util.pid")

	pid.RemovePID()

}
