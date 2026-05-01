//go:build frida_compiler_backend_cli

package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag/v2"
)

func main() {
	// ①+② skip "define your own enum flag type" and enumeration values, as we
	// already have a 3rd party one.

	// ③ Map 3rd party enumeration values to their textual representations
	var LoglevelIds = map[log.Level][]string{
		log.TraceLevel: {"trace"},
		log.DebugLevel: {"debug"},
		log.InfoLevel:  {"info"},
		log.WarnLevel:  {"warning", "warn"},
		log.ErrorLevel: {"error"},
		log.FatalLevel: {"fatal"},
		log.PanicLevel: {"panic"},
	}

	// ④ Define your enum flag value and set the your logging default value.
	var loglevel log.Level = log.WarnLevel

	rootCmd := &cobra.Command{
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Printf("logging level is: %d=%q\n",
				loglevel,
				cmd.PersistentFlags().Lookup("log").Value.String())
		},
	}

	// ⑤ Define the CLI flag parameters for your wrapped enum flag.
	rootCmd.PersistentFlags().Var(
		enumflag.New(&loglevel, "log", LoglevelIds, enumflag.EnumCaseInsensitive),
		"log",
		"sets logging level; can be 'trace', 'debug', 'info', 'warn', 'error', 'fatal', 'panic'")

	// Defaults to what we set above: warn level.
	_ = rootCmd.Execute()

	// User specifies a specific level, such as log level.
	rootCmd.SetArgs([]string{"--log", "debug"})
	_ = rootCmd.Execute()
}
