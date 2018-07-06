package server

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// Pidfile is a struct that describes a PID file.
type Pidfile struct {
	Name string
}

// NewPID creates a new PID file.
func NewPID(name string) *Pidfile {
	log.Debug("Creating and opening PID file...")

	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Error("pidfile: failed to open pid ", err)
		os.Exit(2)
	}

	defer file.Close()

	log.Debug("Writing PID to PID file...")

	pid := fmt.Sprintf("%d", os.Getpid())
	file.Write([]byte(pid))

	log.Debug("PID creation has been completed...")

	return &Pidfile{name}
}

// RemovePID removes the PID file.
func (pf *Pidfile) RemovePID() {
	log.Debug("Removing PID file...")

	err := os.Remove(pf.Name)
	if err != nil {
		log.Error("pidfile: failed to remove ", err)
	}
	log.Debug("PID file removed...")
}
