package cmd

import (
	"fmt"
	"os"

	"gitlab.com/junland/sled/server"
	"gitlab.com/junland/sled/utils"
	flag "github.com/spf13/pflag"
)

var BinVersion string
var GoVersion string

// Default parameters when program starts without flags or enviroment variables.
const (
	defLvl  = "info"
	defPort = "8080"
	defPID  = "/var/run/sled.pid"
	defTLS  = false
	defCert = ""
	defKey  = ""
)

// Configuration falgs and enviroment variables.
var (
	confLogLvl = flag.String("log-level", utils.GetEnvString("SLED_SERVER_PORT", defLvl), "Specify log level for output")
	confPort   = flag.String("port", utils.GetEnvString("SLED_SERVER_PORT", defPort), "Starting server port")
	confPID    = flag.String("pid-file", utils.GetEnvString("SLED_SERVER_PID", defPID), "Specify server PID file path")
	confTLS    = flag.Bool("tls", utils.GetEnvBool("SLED_TLS", false), "Specify weather to run server in secure mode")
	confCert   = flag.String("tls-cert", utils.GetEnvString("SLED_TLS_CERT", defCert), "Specify TLS certificate file path")
	confKey    = flag.String("tls-key", utils.GetEnvString("SLED_TLS_KEY", defKey), "Specify TLS key file path")
	srvVersion = flag.Bool("version", false, "Display version information")
)

// Runs the command line interface.
func Run() {
	flag.Usage = func() {
		fmt.Printf("Usage: sled [options] <command> [<args>]\n")
		fmt.Printf("\n")
		fmt.Printf("A simple web app template.\n")
		fmt.Printf("\n")
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
		fmt.Printf("  -h, --help               display this help and exit\n")
		fmt.Printf("\n")
	}

	config := server.Config{
		LogLvl: *confLogLvl,
		Port:   *confPort,
		PID:    *confPID,
		TLS:    *confTLS,
		Cert:   *confCert,
		Key:    *confKey,
	}

	flag.Parse()

	if *srvVersion == true {
		printVersion()
		os.Exit(0)
	}

	server.Start(config)
}

// Prints version information about the binary.
func printVersion() {
	fmt.Printf("Made with love.\n")
	fmt.Printf("Version: %s\n", BinVersion)
	fmt.Printf("Go Version %s\n", GoVersion)
	fmt.Printf("License: MIT\n")
}
