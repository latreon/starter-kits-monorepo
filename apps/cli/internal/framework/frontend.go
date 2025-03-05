package framework

import (
	"fmt"
	"os"

	"github.com/latreon/starter-kits-monorepo/apps/cli/internal/utils"
)

func GenerateFrontend(framework, buildTool, packageManager string, useTS bool) {
	outputDir := "output"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		return
	}

	if err := os.Chdir(outputDir); err != nil {
		fmt.Printf("Error changing to output directory: %v\n", err)
		return
	}

	switch buildTool {
	case "Vite":
		generateViteProject(framework, packageManager, useTS)
	case "RSBuild":
		generateRSBuildProject(framework, packageManager, useTS)
	}
}

func generateViteProject(framework, packageManager string, useTS bool) {
	// Set template based on framework and TypeScript choice
	var template string
	switch framework {
	case "React":
		template = "react"
		if useTS {
			template = "react-ts"
		}
	case "Preact":
		template = "preact"
		if useTS {
			template = "preact-ts"
		}
	case "Solid":
		template = "solid"
		if useTS {
			template = "solid-ts"
		}
	case "Vue.js":
		template = "vue"
		if useTS {
			template = "vue-ts"
		}
	case "Svelte":
		template = "svelte"
		if useTS {
			template = "svelte-ts"
		}
	}

	fmt.Printf("\nConfiguring Vite for %s...\n", framework)

	// Create Vite project with the correct template
	switch packageManager {
	case "npm":
		if err := utils.ExecuteCommand("npm", "create", "vite@latest", ".", "--", "--template", template); err != nil {
			fmt.Printf("Error configuring Vite: %v\n", err)
			return
		}
		utils.ExecuteCommand("npm", "install")
	case "yarn":
		if err := utils.ExecuteCommand("yarn", "create", "vite", ".", "--template", template); err != nil {
			fmt.Printf("Error configuring Vite: %v\n", err)
			return
		}
		utils.ExecuteCommand("yarn")
	case "pnpm":
		if err := utils.ExecuteCommand("pnpm", "create", "vite", ".", "--template", template); err != nil {
			fmt.Printf("Error configuring Vite: %v\n", err)
			return
		}
		utils.ExecuteCommand("pnpm", "install")
	case "bun":
		if err := utils.ExecuteCommand("bunx", "create-vite", ".", "--template", template); err != nil {
			fmt.Printf("Error configuring Vite: %v\n", err)
			return
		}
		utils.ExecuteCommand("bun", "install")
	}

	fmt.Printf("\n✨ %s project created successfully!\n", framework)
	fmt.Printf("To get started:\n")
	fmt.Printf("  cd output\n")
	fmt.Printf("  %s dev\n", packageManager)
}

func generateRSBuildProject(framework, packageManager string, useTS bool) {
	fmt.Println("\nConfiguring RSBuild...")

	// Initialize project based on package manager
	switch packageManager {
	case "npm":
		utils.ExecuteCommand("npm", "init", "-y")
	case "yarn":
		utils.ExecuteCommand("yarn", "init", "-y")
	case "pnpm":
		utils.ExecuteCommand("pnpm", "init")
	case "bun":
		utils.ExecuteCommand("bun", "init")
	}

	// Install dependencies based on package manager
	installCmd := utils.GetPackageManagerInstallCmd(packageManager)
	utils.ExecuteCommand(installCmd[0], append(installCmd[1:], "@rsbuild/core", "@rsbuild/plugin-react", "-D")...)
	utils.ExecuteCommand(installCmd[0], append(installCmd[1:], "react", "react-dom")...)

	if useTS {
		utils.ExecuteCommand(installCmd[0], append(installCmd[1:], "typescript", "@types/react", "@types/react-dom", "-D")...)
	}
}

// Add other frontend-related functions... 