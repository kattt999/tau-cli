package build

import (
	"github.com/taubyte/tau-cli/cli/common/options"
	"github.com/taubyte/tau-cli/flags"
	"github.com/taubyte/tau-cli/i18n"
	"github.com/urfave/cli/v2"
)

// attachName0 appends the 'Name' flag to commands and updates ArgsUsage and Before properties.
func attachName0(commands []*cli.Command) []*cli.Command {
	for _, cmd := range commands {
		cmd.Flags = append(cmd.Flags, flags.Name)
		cmd.ArgsUsage = i18n.ArgsUsageName
		cmd.Before = options.SetNameAsArgs0
	}

	return commands
}

// Command represents the primary 'build' command for the CLI with its subcommands.
var Command = &cli.Command{
	Name: "build",
	// Subcommands are related to different build process
	Subcommands: attachName0([]*cli.Command{
		// build for configurations
		{
			Name:   "config",
			Action: buildConfig,
		},
		// initiates the code building process
		{
			Name:   "code",
			Action: buildCode,
		},
		// builds a specific function
		{
			Name:   "function",
			Action: buildFunction,
		},
		// build a website
		{
			Name:   "website",
			Action: buildWebsite,
		},
		// build a library
		{
			Name:   "library",
			Action: buildLibrary,
		},
	}),
}
