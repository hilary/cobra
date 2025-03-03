package cobra

import (
	"fmt"
)

type PositionalArgs func(cmd *Command, args []string) error

// validateArgs returns an error if there are any positional args that are not in
// the `ValidArgs` field of `Command`
func validateArgs(cmd *Command, args []string) error {
	if len(cmd.ValidArgs) > 0 {
		for _, v := range args {
			if !stringInSlice(v, cmd.ValidArgs) {
				return fmt.Errorf("invalid argument %q for %q%s", v, cmd.CommandPath(), cmd.findSuggestions(args[0]))
			}
		}
	}
	return nil
}

// NoArgs returns an error if any args are included.
func NoArgs(cmd *Command, args []string) error {
	if len(args) > 0 {
		if cmd.HasAvailableSubCommands() {
			return fmt.Errorf("unknown command %q for %q", args[0], cmd.CommandPath())
		}
		return fmt.Errorf("\"%s\" rejected; %q does not accept args", args[0], cmd.CommandPath())
	}
	return nil
}

// ArbitraryArgs never returns an error.
func ArbitraryArgs(cmd *Command, args []string) error {
	return nil
}

// MinimumNArgs returns an error if there is not at least N args.
func MinimumNArgs(n int) PositionalArgs {
	return func(cmd *Command, args []string) error {
		if len(args) < n {
			return fmt.Errorf("requires at least %d arg(s), only received %d", n, len(args))
		}
		return nil
	}
}

// MaximumNArgs returns an error if there are more than N args.
func MaximumNArgs(n int) PositionalArgs {
	return func(cmd *Command, args []string) error {
		if len(args) > n {
			return fmt.Errorf("accepts at most %d arg(s), received %d", n, len(args))
		}
		return nil
	}
}

// ExactArgs returns an error if there are not exactly n args.
func ExactArgs(n int) PositionalArgs {
	return func(cmd *Command, args []string) error {
		if len(args) != n {
			return fmt.Errorf("accepts %d arg(s), received %d", n, len(args))
		}
		return nil
	}
}

// RangeArgs returns an error if the number of args is not within the expected range.
func RangeArgs(min int, max int) PositionalArgs {
	return func(cmd *Command, args []string) error {
		if len(args) < min || len(args) > max {
			return fmt.Errorf("accepts between %d and %d arg(s), received %d", min, max, len(args))
		}
		return nil
	}
}

// ExactValidArgs returns an error if there are not exactly N positional args OR
// there are any positional args that are not in the `ValidArgs` field of `Command`
//
// Deprecated: now `ExactArgs` honors `ValidArgs`, when defined and not empty
func ExactValidArgs(n int) PositionalArgs {
	return ExactArgs(n)
}

// OnlyValidArgs returns an error if any args are not in the list of `ValidArgs`.
//
// Deprecated: now `ArbitraryArgs` honors `ValidArgs`, when defined and not empty
func OnlyValidArgs(cmd *Command, args []string) error {
	return ArbitraryArgs(cmd, args)
}
