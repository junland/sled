package cmd

import (
	"fmt"

	flag "github.com/spf13/pflag"
	"github.com/junland/sled/server"
)

// BinVersion describes built binary version.
var BinVersion string

// GoVersion Describes Go version that was used to build the binary.
var GoVersion string

// Default parameters when program starts without flags or environment variables.
const (
	defLvl    = "info"
	defAccess = true
	defPort   = "8080"
	defPID    = "/var/run/sled.pid"
	defTLS    = false
	defCert   = ""
	defKey    = ""
)

var (
	confLogLvl, confPort, confPID, confCert, confKey string
	enableTLS, enableAccess, version, help           bool
)

// init defines configuration flags and environment variables.
func init() {
	flags := flag.CommandLine
	flags.StringVar(&confLogLvl, "log-level", GetEnvString("SLED_LOG_LEVEL", defLvl), "Specify log level for output.")
	flags.BoolVar(&enableAccess, "access-log", GetEnvBool("SLED_ACCESS_LOG", defAccess), "Specify weather to run with or without HTTP access logs.")
	flags.StringVar(&confPort, "port", GetEnvString("SLED_SERVER_PORT", defPort), "Starting server port.")
	flags.StringVar(&confPID, "pid-file", GetEnvString("SLED_SERVER_PID", defPID), "Specify server PID file path.")
	flags.BoolVar(&enableTLS, "tls", GetEnvBool("SLED_TLS", defTLS), "Specify weather to run server in secure mode.")
	flags.StringVar(&confCert, "tls-cert", GetEnvString("SLED_TLS_CERT", defCert), "Specify TLS certificate file path.")
	flags.StringVar(&confKey, "tls-key", GetEnvString("SLED_TLS_KEY", defKey), "Specify TLS key file path.")
	flags.BoolVarP(&help, "help", "h", false, "Show this help")
	flags.BoolVar(&version, "version", false, "Display version information")
	flags.SortFlags = false
	flag.Parse()
}

// PrintHelp prints help text.
func PrintHelp() {
	fmt.Printf("Usage: sled [options] <command> [<args>]\n")
	fmt.Printf("\n")
	fmt.Printf("A simple web app template.\n")
	fmt.Printf("\n")
	fmt.Printf("Options:\n")
	flag.PrintDefaults()
	fmt.Printf("\n")
}

// PrintVersion prints version information about the binary.
func PrintVersion() {
	fmt.Printf("Made with love.\n")
	fmt.Printf("Version: %s\n", BinVersion)
	fmt.Printf("Go Version %s\n", GoVersion)
	fmt.Printf("License: MIT\n")
}

// Run is the entry point for starting the command line interface.
func Run() {
	config := server.Config{
		LogLvl: confLogLvl,
		Access: enableAccess,
		Port:   confPort,
		PID:    confPID,
		TLS:    enableTLS,
		Cert:   confCert,
		Key:    confKey,
	}

	if version {
		PrintVersion()
		return
	}

	if help {
		PrintHelp()
		return
	}

	server.Start(config)
}
