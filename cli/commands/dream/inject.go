package dream

import (
	dreamI18n "github.com/taubyte/tau-cli/i18n/dream"
	dreamLib "github.com/taubyte/tau-cli/lib/dream"
	loginLib "github.com/taubyte/tau-cli/lib/login"
	projectLib "github.com/taubyte/tau-cli/lib/project"
	"github.com/urfave/cli/v2"
)

// provides a CLI sub command to inject a given project to the dreamland
var injectCommand = &cli.Command{
	Name: "inject",
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

				// retrieve the details of the currently logged-in profile.
				profile, err := loginLib.GetSelectedProfile()
				if err != nil {
					return err
				}

				// set up the project to import in dreamland.
				prodProject := &dreamLib.ProdProject{
					Project: project,
					Profile: profile,
				}

				return prodProject.Import()
			},
		},
	},
}
