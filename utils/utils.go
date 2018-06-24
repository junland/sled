package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// Pidfile is a struct that describes a PID file.
type Pidfile struct {
	Name string
}

// NewPID creates a new PID file.
func NewPID(name string) *Pidfile {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Error("pidfile: failed to open pid ", err)
		os.Exit(2)
	}
	defer file.Close()

	pid := fmt.Sprintf("%d", os.Getpid())
	file.Write([]byte(pid))

	return &Pidfile{name}
}

// RemovePID removes the PID file.
func (pf *Pidfile) RemovePID() {
	err := os.Remove(pf.Name)
	if err != nil {
		log.Error("pidfile: failed to remove ", err)
	}
}

// ReadPID defines a PID file with a specified filename.
func ReadPID(fileName string) (int, error) {
	var pid int
	p, err := ioutil.ReadFile(fileName)
	if err != nil {
		return pid, err
	}

	pid, err = strconv.Atoi(string(p))
	if err != nil {
		return pid, err
	}
	return pid, nil
}

// GetEnvString defines a environment variable with a specified name, fallback value.
// The return is a string value.
func GetEnvString(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GetEnvBool defines a environment variable with a specified name, fallback value.
// The return is either a true or false.
func GetEnvBool(key string, fallback bool) bool {
	switch os.Getenv(key) {
	case "true":
		return true
	case "false":
		return false
	default:
		return fallback
	}
}
