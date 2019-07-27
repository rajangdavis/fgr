package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "fgr",
	Short: "fgr - build charts from configuration",
	Long: `fgr - a Command Line Interface application to build charts from configuration files

fgr  Copyright (C) 2018 Rajan G. Davis
	
This program comes with ABSOLUTELY NO WARRANTY.
This is free software, and you are welcome to redistribute it
under certain conditions`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "path to file")
}