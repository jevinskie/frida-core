//go:build frida_compiler_backend_cli

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag/v2"
)

var module string

var output string

var compress bool

var verbose bool

type BundleFormatKind enumflag.Flag

const (
	Esm BundleFormatKind = iota
	Iife
)

var BundleFormatKindIds = map[BundleFormatKind][]string{
	Esm:  {"esm"},
	Iife: {"iife"},
}
var bundleformat BundleFormatKind = Esm

type TypeCheckKind enumflag.Flag

const (
	Full TypeCheckKind = iota
	None
)

var TypeCheckKindIds = map[TypeCheckKind][]string{
	Full: {"full"},
	None: {"none"},
}
var typecheck TypeCheckKind = Full

type PlatformKind enumflag.Flag

const (
	Gum PlatformKind = iota
	Browser
	Neutral
)

var PlatformKindIds = map[PlatformKind][]string{
	Gum:     {"gum"},
	Browser: {"browser"},
	Neutral: {"neutral"},
}
var platform PlatformKind = Gum

var external = []string{}

var rootCmd = &cobra.Command{
	Use:  "frida-core-compiler MODULE",
	Run:  run,
	Args: cobra.ExactArgs(1),
}

func init() {
	// rootCmd.PersistentFlags().String(&module, "TypeScript/JavaScript module to compile")
	rootCmd.PersistentFlags().BoolVarP(&compress, "compress", "c", false, "minify code")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "be verbose")
	rootCmd.PersistentFlags().VarP(
		enumflag.New(&bundleformat, "bundle-format", BundleFormatKindIds, enumflag.EnumCaseSensitive),
		"bundle-format", "B",
		"desired bundle format")
	rootCmd.PersistentFlags().VarP(
		enumflag.New(&typecheck, "type-check", TypeCheckKindIds, enumflag.EnumCaseSensitive),
		"type-check", "T",
		"desired type-checking mode")
	rootCmd.PersistentFlags().VarP(
		enumflag.New(&platform, "platform", PlatformKindIds, enumflag.EnumCaseSensitive),
		"platform", "",
		"JavaScript runtime platform")
	rootCmd.Flags().StringArrayVarP(&external, "external", "E", []string{}, "mark MODULE as external (may be specified multiple times)")
}

func run(cmd *cobra.Command, args []string) {
	module = args[0]
	fmt.Printf("cmd: %v\n", cmd)
	fmt.Printf("args: %v\n", args)
	fmt.Printf("module is: %v\n",
		module)
	fmt.Printf("compress is: %v\n",
		compress)
	fmt.Printf("verbose is: %v\n",
		verbose)
	fmt.Printf("bundle-format is: %v=%q\n",
		bundleformat,
		cmd.PersistentFlags().Lookup("bundle-format").Value.String())
	fmt.Printf("type-check is: %v=%q\n",
		typecheck,
		cmd.PersistentFlags().Lookup("type-check").Value.String())
	fmt.Printf("platform is: %v=%q\n",
		platform,
		cmd.PersistentFlags().Lookup("platform").Value.String())
	fmt.Printf("external is: %v\n",
		external)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
