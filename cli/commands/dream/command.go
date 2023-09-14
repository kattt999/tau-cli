package dream

import (
	"fmt"

	"github.com/taubyte/tau-cli/cli/commands/dream/build"
	dreamLib "github.com/taubyte/tau-cli/lib/dream"
	projectLib "github.com/taubyte/tau-cli/lib/project"
	"github.com/urfave/cli/v2"
)

const (
	defaultBind        = "node@1/verbose,seer@2/copies,node@2/copies"
	dreamCacheLocation = "~/.cache/dreamland/universe-tau"
)

// config for caching a dream
var (
	cacheDream = []string{"--id", "tau", "--keep"}
)

// defines the main CLI command for interfacing with the local taubyte network
var Command = &cli.Command{
	Name:  "dream",
	Usage: "Starts and interfaces with a local taubyte network.  All leading arguments to `tau dream ...` are passed to dreamland",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "cache",
			Usage: fmt.Sprintf("caches the universe in `%s` keeping data for subsequent restarts", dreamCacheLocation),
		},
	},
	Action: func(c *cli.Context) error {
		// fetch the selected project
		project, err := projectLib.SelectedProjectInterface()
		if err != nil {
			return err
		}

		// get the repository for the selected project
		h := projectLib.Repository(project.Get().Name())
		projectRepositories, err := h.Open()
		if err != nil {
			return err
		}

		// fetch the current branch of the project
		branch, err := projectRepositories.CurrentBranch()
		if err != nil {
			return err
		}

		// build the base arguments
		baseStartDream := []string{"new", "multiverse", "--bind", defaultBind, "--branch", branch}
		if c.IsSet("cache") {
			return dreamLib.Execute(append(baseStartDream, cacheDream...)...)
		} else {
			return dreamLib.Execute(baseStartDream...)
		}
	},

	Subcommands: []*cli.Command{
		injectCommand,
		attachCommand,
		build.Command,
	},
}
