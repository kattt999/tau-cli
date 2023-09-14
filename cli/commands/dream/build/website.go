package build

import (
	"fmt"
	"os"
	"path"
	"strings"

	dreamI18n "github.com/taubyte/tau-cli/i18n/dream"
	websiteI18n "github.com/taubyte/tau-cli/i18n/website"
	applicationLib "github.com/taubyte/tau-cli/lib/application"
	dreamLib "github.com/taubyte/tau-cli/lib/dream"
	websitePrompts "github.com/taubyte/tau-cli/prompts/website"
	"github.com/urfave/cli/v2"
)

// initiate the building process for the website
func buildWebsite(ctx *cli.Context) error {
	// make sure DreamLand system is running
	if !dreamLib.IsRunning() {
		dreamI18n.Help().IsDreamlandRunning()
		return dreamI18n.ErrorDreamlandNotStarted
	}

	// fetch the website from websitePrompts
	website, err := websitePrompts.GetOrSelect(ctx)
	if err != nil {
		return err
	}

	// initialize build environment
	builder, err := initBuild()
	if err != nil {
		return err
	}

	// set up compilation parameters for the website
	compileFor := &dreamLib.CompileForRepository{
		ProjectId:  builder.project.Get().Id(),
		ResourceId: website.Id,
		Branch:     builder.currentBranch,
	}

	// if a website has been selected, fetch website details and update compilation parameters
	if len(builder.selectedApp) > 0 {
		app, err := applicationLib.Get(builder.selectedApp)
		if err != nil {
			return err
		}

		compileFor.ApplicationId = app.Id
	}

	// ensure the website name is in the correct format
	splitName := strings.Split(website.RepoName, "/")
	if len(splitName) != 2 {
		return fmt.Errorf("invalid repository name `%s` expected `user/repo`", website.RepoName)
	}

	// check if the website exists at the given path
	compileFor.Path = path.Join(builder.projectConfig.WebsiteLoc(), splitName[1])
	_, err = os.Stat(compileFor.Path)
	if err != nil {
		websiteI18n.Help().BeSureToCloneWebsite()
		return err
	}

	// start the compilation
	return compileFor.Execute()
}
