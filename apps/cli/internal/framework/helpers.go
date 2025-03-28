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
	fmt.Println("• Websites, Mobile Apps, APIs, Desktop Apps, and CLI Tools")
	fmt.Println("• npm, yarn, pnpm, bun")
	fmt.Println("\n--------------------------------------------")

	fmt.Println("\n📦 Available Frameworks:")
	fmt.Println("Website:           Mobile:            API:                Desktop:           CLI:")
	fmt.Println("- Next.js          - React Native     - Node.js/Express   - Electron         - Go/Cobra")
	fmt.Println("- React            - Flutter          - Nest.js           - Tauri            - Node.js/Commander")
	fmt.Println("- Vue.js           - Ionic           - FastAPI           - Qt               - Python/Click")
	fmt.Println("- Svelte")
	fmt.Println("- Preact")
	fmt.Println("- Solid")
	fmt.Println("- Astro")
	fmt.Println(" ")
}

func SelectFramework() string {
	// First ask what they want to build
	buildTypes := []string{
		"Website",
		"Mobile App",
		"API",
		"Desktop App",
		"CLI Tool",
	}
	var buildType string
	buildPrompt := &survey.Select{
		Message:  "What do you want to build?",
		Options:  buildTypes,
		PageSize: 10,
	}
	if err := survey.AskOne(buildPrompt, &buildType); err != nil {
		if err == terminal.InterruptErr {
			utils.HandleInterrupt()
		}
	}

	// Show different options based on build type
	var frameworks []string
	switch buildType {
	case "Website":
		frameworks = []string{
			"React",
			"Next.js",
			"Vue.js",
			"Svelte",
			"Astro",
		}
	case "Mobile App":
		frameworks = []string{
			"React Native",
			"Flutter",
			"Ionic",
		}
	case "API":
		frameworks = []string{
			"Node.js/Express",
			"Nest.js",
			"FastAPI",
			"Go/Fiber",
		}
	case "Desktop App":
		frameworks = []string{
			"Electron",
			"Tauri",
			"Qt",
		}
	case "CLI Tool":
		frameworks = []string{
			"Go/Cobra",
			"Node.js/Commander",
			"Python/Click",
		}
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

	// Ask for additional options based on selection
	if buildType == "API" {
		// Ask for database
		var database string
		dbPrompt := &survey.Select{
			Message: "Choose a database:",
			Options: []string{
				"PostgreSQL",
				"MongoDB",
				"MySQL",
				"SQLite",
			},
		}
		if err := survey.AskOne(dbPrompt, &database); err != nil {
			if err == terminal.InterruptErr {
				utils.HandleInterrupt()
			}
		}
	}

	return selected
}

func SelectBuildTool(framework string) string {
	var buildTool string
	switch framework {
	case "React":
		prompt := &survey.Select{
			Message: "Choose a build tool:",
			Options: []string{
				"Vite (Recommended)",
				"RSBuild",
				"Parcel",
				"Webpack",
			},
		}
		if err := survey.AskOne(prompt, &buildTool); err != nil {
			if err == terminal.InterruptErr {
				utils.HandleInterrupt()
			}
		}

		// Additional configuration based on build tool
		switch buildTool {
		case "Vite (Recommended)":
			var config string
			configPrompt := &survey.Select{
				Message: "Choose Vite configuration:",
				Options: []string{
					"Default (Recommended)",
					"Default + SWC",
				},
			}
			survey.AskOne(configPrompt, &config)
		case "RSBuild":
			var config string
			configPrompt := &survey.Select{
				Message: "Choose RSBuild configuration:",
				Options: []string{
					"Modern Build (Recommended)",
					"Legacy Build",
					"Custom Build",
				},
			}
			survey.AskOne(configPrompt, &config)
		case "Parcel":
			var config string
			configPrompt := &survey.Select{
				Message: "Choose Parcel configuration:",
				Options: []string{
					"Basic Setup (Recommended)",
					"Zero Config",
					"Custom Config",
				},
			}
			survey.AskOne(configPrompt, &config)
		case "Webpack":
			var config string
			configPrompt := &survey.Select{
				Message: "Choose Webpack configuration:",
				Options: []string{
					"Basic Config (Recommended)",
					"Development Optimized",
					"Production Optimized",
					"Custom Config",
				},
			}
			survey.AskOne(configPrompt, &config)
		}
	case "Vue.js":
		prompt := &survey.Select{
			Message: "Choose a build tool:",
			Options: []string{
				"Vite (Recommended)",
				"Vue CLI",
				"Nuxt",
				"Webpack",
			},
		}
		if err := survey.AskOne(prompt, &buildTool); err != nil {
			if err == terminal.InterruptErr {
				utils.HandleInterrupt()
			}
		}
	case "Node.js":
		prompt := &survey.Select{
			Message: "Choose a Node.js template:",
			Options: []string{
				"Express (Recommended)",
				"Fastify",
				"Koa",
			},
		}
		if err := survey.AskOne(prompt, &buildTool); err != nil {
			if err == terminal.InterruptErr {
				utils.HandleInterrupt()
			}
		}
	case "Svelte":
		prompt := &survey.Select{
			Message: "Choose a build tool:",
			Options: []string{
				"Vite (Recommended)",
				"SvelteKit",
				"Rollup",
			},
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
		Message: "Would you like to use TypeScript?",
		Default: true,
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
