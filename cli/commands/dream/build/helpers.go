package build

import (
	"github.com/taubyte/go-project-schema/project"
	"github.com/taubyte/tau-cli/env"
	projectLib "github.com/taubyte/tau-cli/lib/project"
	"github.com/taubyte/tau-cli/singletons/config"
)

// a utility structure to assist with the build process
type buildHelper struct {
	project       project.Project
	projectConfig config.Project
	currentBranch string
	selectedApp   string
}

// initialize the buildHelper structure by gathering the project and repository details for the building process
func initBuild() (*buildHelper, error) {
	var err error
	helper := &buildHelper{}

	// fectch selected project's interface
	helper.project, err = projectLib.SelectedProjectInterface()
	if err != nil {
		return nil, err
	}

	// get configuration details for the selected project
	helper.projectConfig, err = projectLib.SelectedProjectConfig()
	if err != nil {
		return nil, err
	}

	// open the project repository based on the project name
	h := projectLib.Repository(helper.project.Get().Name())
	projectRepositories, err := h.Open()
	if err != nil {
		return nil, err
	}

	// fetch the current branch for the opened repository
	helper.currentBranch, err = projectRepositories.CurrentBranch()
	if err != nil {
		return nil, err
	}

	// fetch the selected application's name
	helper.selectedApp, _ = env.GetSelectedApplication()
	return helper, nil
}
