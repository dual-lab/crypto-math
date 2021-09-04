package cmd

import (
	"strconv"

	"com.github/dual-lab/crypto-math/m/cmd/err"
	"github.com/spf13/cobra"
)

func argsIntValidators(command *cobra.Command, args []string) error {
	for _, arg := range args {
		if _, e := strconv.Atoi(arg); e != nil {
			return err.ErrInvaliTypedArg
		}
	}

	return nil

}

func agrsLenValidators(command *cobra.Command, args []string, exactlen int) error {
	if len(args) != exactlen {
		return err.ErrInvalidArg
	}
	return nil
}
