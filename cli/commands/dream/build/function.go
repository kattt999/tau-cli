package build

import (
	"path"

	commonSpec "github.com/taubyte/go-specs/common"
	functionSpec "github.com/taubyte/go-specs/function"
	dreamI18n "github.com/taubyte/tau-cli/i18n/dream"
	applicationLib "github.com/taubyte/tau-cli/lib/application"
	dreamLib "github.com/taubyte/tau-cli/lib/dream"
	functionPrompts "github.com/taubyte/tau-cli/prompts/function"
	"github.com/urfave/cli/v2"
)

// build a specific function after checking if Dreamland system is running
func buildFunction(ctx *cli.Context) error {
	// ensure dreamland system is running
	if !dreamLib.IsRunning() {
		dreamI18n.Help().IsDreamlandRunning()
		return dreamI18n.ErrorDreamlandNotStarted
	}

	// select a function based on the context
	function, err := functionPrompts.GetOrSelect(ctx)
	if err != nil {
		return err
	}

	// initiliaze the build helper to gather essential build details
	builder, err := initBuild()
	if err != nil {
		return err
	}

	// prepare data for compling the function
	compileFor := &dreamLib.CompileForDFunc{
		ProjectId:  builder.project.Get().Id(),
		ResourceId: function.Id,
		Branch:     builder.currentBranch,
		Call:       function.Call,
	}

	// check if an application has been selected
	if len(builder.selectedApp) > 0 {
		app, err := applicationLib.Get(builder.selectedApp)
		if err != nil {
			return err
		}

		compileFor.ApplicationId = app.Id
		compileFor.Path = path.Join(builder.projectConfig.CodeLoc(), commonSpec.ApplicationPathVariable.String(), builder.selectedApp, functionSpec.PathVariable.String(), function.Name)
	} else {
		compileFor.Path = path.Join(builder.projectConfig.CodeLoc(), functionSpec.PathVariable.String(), function.Name)
	}

	return compileFor.Execute()
}
