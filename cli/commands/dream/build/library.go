package build

import (
	"fmt"
	"os"
	"path"
	"strings"

	dreamI18n "github.com/taubyte/tau-cli/i18n/dream"
	libraryI18n "github.com/taubyte/tau-cli/i18n/library"
	applicationLib "github.com/taubyte/tau-cli/lib/application"
	dreamLib "github.com/taubyte/tau-cli/lib/dream"
	libraryPrompts "github.com/taubyte/tau-cli/prompts/library"
	"github.com/urfave/cli/v2"
)

// initiate the build process for a given library
func buildLibrary(ctx *cli.Context) error {
	// make sure Dreamland system is running
	if !dreamLib.IsRunning() {
		dreamI18n.Help().IsDreamlandRunning()
		return dreamI18n.ErrorDreamlandNotStarted
	}

	// fetch a library from libraryPrompts
	library, err := libraryPrompts.GetOrSelect(ctx)
	if err != nil {
		return err
	}

	// initialize build environment
	builder, err := initBuild()
	if err != nil {
		return err
	}

	// set up compilation parameters for the repository
	compileFor := &dreamLib.CompileForRepository{
		ProjectId:  builder.project.Get().Id(),
		ResourceId: library.Id,
		Branch:     builder.currentBranch,
	}

	// if an app has been chosen, fetch app details and update the compilation parameters
	if len(builder.selectedApp) > 0 {
		app, err := applicationLib.Get(builder.selectedApp)
		if err != nil {
			return err
		}

		compileFor.ApplicationId = app.Id
	}

	// make sure the library repository name is in the correct format
	splitName := strings.Split(library.RepoName, "/")
	if len(splitName) != 2 {
		return fmt.Errorf("invalid repository name `%s` expected `user/repo`", library.RepoName)
	}

	// set the path
	compileFor.Path = path.Join(builder.projectConfig.LibraryLoc(), splitName[1])

	// check if the library exists at the given path
	_, err = os.Stat(compileFor.Path)
	if err != nil {
		libraryI18n.Help().BeSureToCloneLibrary()
		return err
	}

	// start the compilation process
	return compileFor.Execute()
}
