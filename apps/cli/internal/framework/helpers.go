package framework

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/latreon/starter-kits-monorepo/apps/cli/internal/utils"
)

func DisplayFrameworks() {
	fmt.Println("\n🚀 Welcome to Starter Kits!")
	fmt.Println("A modern CLI tool to bootstrap your projects with popular frameworks and tools.")
	fmt.Println("\nSupports:")
	fmt.Println("• Frontend & Backend frameworks")
	fmt.Println("• Vite, RSBuild")
	fmt.Println("• TypeScript")
	fmt.Println("• npm, yarn, pnpm, bun")
	fmt.Println("\n--------------------------------------------")

	fmt.Println("\n📦 Available Frameworks:")
	fmt.Println("Frontend:          Backend:")
	fmt.Println("- React            - Node.js")
	fmt.Println("- Preact           - Nest.js")
	fmt.Println("- Solid")
	fmt.Println("- Vue.js")
	fmt.Println("- Svelte")
	fmt.Println("- Next.js")
	fmt.Println()
}

func SelectFramework() string {
	frameworks := []string{
		"React",
		"Preact",
		"Solid",
		"Vue.js",
		"Svelte",
		"Next.js",
		"Node.js",
		"Nest.js",
	}

	var selected string
	prompt := &survey.Select{
		Message:  "Choose a framework:",
		Options:  frameworks,
		PageSize: 10,
	}
	if err := survey.AskOne(prompt, &selected); err != nil {
		if err == terminal.InterruptErr {
			utils.HandleInterrupt()
		}
	}
	return selected
}

func SelectBuildTool(framework string) string {
	var buildTool string
	switch framework {
	case "React", "Preact", "Solid", "Vue.js", "Svelte":
		prompt := &survey.Select{
			Message: "Choose a build tool:",
			Options: []string{"Vite", "RSBuild"},
		}
		if err := survey.AskOne(prompt, &buildTool); err != nil {
			if err == terminal.InterruptErr {
				utils.HandleInterrupt()
			}
		}
	case "Node.js":
		prompt := &survey.Select{
			Message: "Choose a Node.js template:",
			Options: []string{"Express", "Fastify", "Koa"},
		}
		if err := survey.AskOne(prompt, &buildTool); err != nil {
			if err == terminal.InterruptErr {
				utils.HandleInterrupt()
			}
		}
	}
	return buildTool
}

func SelectPackageManager() string {
	packageManagers := []string{"npm", "yarn", "pnpm", "bun"}
	var selected string
	prompt := &survey.Select{
		Message: "Choose a package manager:",
		Options: packageManagers,
	}
	if err := survey.AskOne(prompt, &selected); err != nil {
		if err == terminal.InterruptErr {
			utils.HandleInterrupt()
		}
	}
	return selected
}

func SelectTypeScript() bool {
	var useTS bool
	prompt := &survey.Confirm{
		Message: "Use TypeScript?",
	}
	if err := survey.AskOne(prompt, &useTS); err != nil {
		if err == terminal.InterruptErr {
			utils.HandleInterrupt()
		}
	}
	return useTS
}

func IsFrontend(framework string) bool {
	switch framework {
	case "React", "Preact", "Solid", "Vue.js", "Svelte", "Next.js":
		return true
	default:
		return false
	}
} 