package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type Pidfile struct {
	Name string
}

// Creates a new PID file.
func NewPID(name string) *Pidfile {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Error("pidfile: failed to open pid %s (%s)", name, err)
		os.Exit(2)
	}
	defer file.Close()

	pid := fmt.Sprintf("%d", os.Getpid())
	file.Write([]byte(pid))

	return &Pidfile{name}
}

// Removes the PID file.
func (pf *Pidfile) RemovePID() {
	err := os.Remove(pf.Name)
	if err != nil {
		log.Error("pidfile: failed to remove %s (%s)", pf.Name, err)
	}
}

// Reads the PID file, returns the contents of the PID file.
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

// Looksup current enviroment variable with a default variable if given ENV is not found.
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
