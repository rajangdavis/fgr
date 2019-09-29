package cmd

import (
	// "fmt"
	"github.com/spf13/cobra"
	// "gopkg.in/yaml.v2"
)

var configFile string
var csvFromFile string
var yAxis string
var xAxis string
var agg string

var chart = &cobra.Command{
	Use:   "chart",
	Short: "chart - build charts from configuration",
	Long: `chart - build charts from configuration - will add configurations as needed`,
}

func init() {
	chart.Flags().StringVarP(&configFile, "file", "f", "", "Points to the file location of an existing config file")
	chart.Flags().StringVarP(&csvFromFile, "csvfile", "c", "", "Hint")
	chart.Flags().StringVarP(&yAxis, "yaxis", "y", "", "Hint")
	chart.Flags().StringVarP(&xAxis, "xaxis", "x", "", "Hint")
	chart.Flags().StringVarP(&agg, "agg", "a", "", "Hint")
	fgr.AddCommand(chart)
}