package cmd

import (
	"fmt"
	"strconv"

	"com.github/dual-lab/crypto-math/m/algebra"
	"com.github/dual-lab/crypto-math/m/cmd/err"
	"github.com/spf13/cobra"
)

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
		// Retrive a and b
		a, _ := strconv.Atoi(args[0])
		m, _ := strconv.Atoi(args[1])
		inv, ok := algebra.InverModM(a, m)
		if ok == false {
			return err.ErrInverseNotExists
		}
		fmt.Printf("%d\n", inv)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(inverModMCmd)
}
