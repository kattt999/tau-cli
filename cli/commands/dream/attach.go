package dream

import (
	dreamI18n "github.com/taubyte/tau-cli/i18n/dream"
	dreamLib "github.com/taubyte/tau-cli/lib/dream"
	loginLib "github.com/taubyte/tau-cli/lib/login"
	projectLib "github.com/taubyte/tau-cli/lib/project"
	"github.com/urfave/cli/v2"
)

// define a CLI command to attach a given project to DreamLand
var attachCommand = &cli.Command{
	Name: "attach",
	Subcommands: []*cli.Command{
		{
			Name: "project",
			Action: func(ctx *cli.Context) error {
				// make sure dreamland is running
				if !dreamLib.IsRunning() {
					dreamI18n.Help().IsDreamlandRunning()
					return dreamI18n.ErrorDreamlandNotStarted
				}

				// fetch the selected project details
				project, err := projectLib.SelectedProjectInterface()
				if err != nil {
					return err
				}

				// fetch details of the profile currently logged in
				profile, err := loginLib.GetSelectedProfile()
				if err != nil {
					return err
				}

				// attach the selected project to the profile
				prodProject := &dreamLib.ProdProject{
					Project: project,
					Profile: profile,
				}

				return prodProject.Attach()
			},
		},
	},
}
