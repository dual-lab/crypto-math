package cmd

import (
	"strconv"

	"com.github/dual-lab/crypto-math/m/cmd/err"
	"github.com/spf13/cobra"
)

var gcdCmd = &cobra.Command{
	Use:   "gcd a b",
	Short: "Calculate the greatest common divisor",
	Long:  "Use the ecludian extended algorithm to calculate the gcd, u, v",
	Args:  argsValidators,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var extended bool

func argsValidators(command *cobra.Command, args []string) error {
	if len(args) != 2 {
		return err.ErrInvalidArg
	}
	if _, e := strconv.Atoi(args[0]); e != nil {
		return err.ErrInvaliTypedArg
	}

	if _, e := strconv.Atoi(args[1]); e != nil {
		return err.ErrInvaliTypedArg
	}

	return nil

}

func init() {
	gcdCmd.Flags().BoolVarP(&extended, "extended", "e", false, "Calculate the extended Eulero algorithm coefficent")
	rootCmd.AddCommand(gcdCmd)
}
