package cmd

import "github.com/spf13/cobra"

var inverModMCmd = &cobra.Command{
	Use:   "invm a m",
	Short: "Calculate the inverse of a mod m.",
	Long:  "Use the ecludian extended algorithm to calculate b | ab = 1 (mod m) if gcd(a,m)=1 ",
	Args: func(command *cobra.Command, args []string) error {
		if e := agrsLenValidators(command, args, 2); e != nil {
			return e
		}
		if e := argsIntValidators(command, args); e != nil {
			return e
		}

		return nil
	},
	RunE: func(commadn *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(inverModMCmd)
}
