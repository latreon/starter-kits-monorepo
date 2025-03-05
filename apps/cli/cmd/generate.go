package cmd

import (
	"github.com/spf13/cobra"
	"github.com/latreon/starter-kits-monorepo/apps/cli/internal/framework"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new starter kit",
	Run:   runGenerate,
}

func runGenerate(cmd *cobra.Command, args []string) {
	// Show available frameworks
	framework.DisplayFrameworks()

	// Get user selections
	selectedFramework := framework.SelectFramework()
	buildTool := framework.SelectBuildTool(selectedFramework)
	packageManager := framework.SelectPackageManager()
	useTS := framework.SelectTypeScript()

	// Generate the project
	if framework.IsFrontend(selectedFramework) {
		framework.GenerateFrontend(selectedFramework, buildTool, packageManager, useTS)
	} else {
		framework.GenerateBackend(selectedFramework, buildTool, packageManager, useTS)
	}
} 