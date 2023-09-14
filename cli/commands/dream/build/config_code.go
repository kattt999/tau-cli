package build

import (
	dreamI18n "github.com/taubyte/tau-cli/i18n/dream"
	dreamLib "github.com/taubyte/tau-cli/lib/dream"
	"github.com/urfave/cli/v2"
)

// checks if the DreamLand system is running and initiates the build for configurations or code based on input flags
func executeConfigCode(config bool, code bool) error {
	// ensure dreamLib is running
	if !dreamLib.IsRunning() {
		dreamI18n.Help().IsDreamlandRunning()
		return dreamI18n.ErrorDreamlandNotStarted
	}

	// initialize the build helper to gather essential build details
	builder, err := initBuild()
	if err != nil {
		return err
	}

	// use the local build library to build either configuration or code
	return dreamLib.BuildLocalConfigCode{
		Config:      config,
		Code:        code,
		Branch:      builder.currentBranch,
		ProjectPath: builder.projectConfig.Location,
		ProjectID:   builder.project.Get().Id(),
	}.Execute()
}

// triggers the build for configuarations
func buildConfig(ctx *cli.Context) error {
	return executeConfigCode(true, false)
}

// trigers the build for code
func buildCode(ctx *cli.Context) error {
	return executeConfigCode(false, true)
}
