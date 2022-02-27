package commands

import (
	"fmt"
	"os"

	"github.com/enenumxela/nmapx/internal/configuration"
	"github.com/spf13/cobra"
)

var (
	RootCMD = &cobra.Command{
		Use:   "nmapx",
		Short: "A CLI utility to merge and/or parse nmap XML output",
		Long:  configuration.BANNER,
	}
)

func init() {
	RootCMD.AddCommand(Merge())
}

func Execute() {
	if err := RootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
