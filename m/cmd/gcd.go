package cmd

import (
	"fmt"
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
		// Retrive a and b
		a, _ := strconv.Atoi(args[0])
		b, _ := strconv.Atoi(args[1])
		// Retrive extended flags
		if extended != 0 && extended != 1 && extended != 2 {
			return err.ErrInvalidFlagValue
		}
		fmt.Printf("gdm(%d, %d)\n", a, b)
		return nil
	},
}

var extended uint8

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
	gcdCmd.Flags().Uint8VarP(&extended, "extended", "e", 0, "Calculate the extended Eulero algorithm coefficents. 0(disalbed) 1(enabled) 2(u min positive solution)")
	rootCmd.AddCommand(gcdCmd)
}
