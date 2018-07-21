package server

import (
	"os"
	"testing"
)

func TestCreateAndRemovePID(t *testing.T) {
	os.RemoveAll("./test-util.pid")

	pid := CreatePID("./test-util.pid")

	pid.RemovePID()

}
