package util

import (
	"context"
	"os"

	"github.com/adrianriobo/pupoc/pkg/util/logging"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PluginInfo struct {
	name    string
	version string
}

type DeployableTarget struct {
	projectName string
	stackName   string
	deployFunc  pulumi.RunFunc
	plugin      PluginInfo
}

// this function gets our stack ready for update/destroy by prepping the workspace, init/selecting the stack
// and doing a refresh to make sure state and cloud resources are in sync
func CreateOrSelectStack(ctx context.Context, target DeployableTarget) auto.Stack {
	// create or select a stack with an inline Pulumi program
	s, err := auto.UpsertStackInlineSource(ctx, target.stackName, target.projectName, target.deployFunc)
	if err != nil {
		logging.Errorf("Failed to create or select stack: %v", err)
		os.Exit(1)
	}
	logging.Debugf("Created/Selected stack %q\n", target.stackName)

	w := s.Workspace()

	// for inline source programs, we must manage plugins ourselves
	err = w.InstallPlugin(ctx, target.plugin.name, target.plugin.version)
	if err != nil {
		logging.Errorf("Failed to install program plugins: %v", err)
		os.Exit(1)
	}

	if err := s.SetConfig(ctx, "aws:region", auto.ConfigValue{Value: "us-west-2"}); err != nil {
		logging.Errorf("Failed setting the region: %v", err)
		os.Exit(1)
	}
	_, err = s.Refresh(ctx)
	if err != nil {
		logging.Errorf("Failed to refresh stack: %v\n", err)
		os.Exit(1)
	}
	return s
}
