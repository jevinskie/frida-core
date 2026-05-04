//go:build frida_compiler_backend_cli

package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag/v2"
)

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

func main() {
	rootCmd := &cobra.Command{
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Printf("bundle-format is: %d=%q\n",
				bundleformat,
				cmd.PersistentFlags().Lookup("bundle-format").Value.String())
			fmt.Printf("type-check is: %d=%q\n",
				typecheck,
				cmd.PersistentFlags().Lookup("type-check").Value.String())
			fmt.Printf("platform is: %d=%q\n",
				platform,
				cmd.PersistentFlags().Lookup("platform").Value.String())
			fmt.Printf("external is: %d\n",
				external)
		},
	}
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

	_ = rootCmd.Execute()
}
