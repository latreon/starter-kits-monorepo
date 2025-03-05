package framework

import (
	"fmt"
	"os"

	"github.com/latreon/starter-kits-monorepo/apps/cli/internal/utils"
)

func GenerateBackend(framework, buildTool, packageManager string, useTS bool) {
	outputDir := "output"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		return
	}

	if err := os.Chdir(outputDir); err != nil {
		fmt.Printf("Error changing to output directory: %v\n", err)
		return
	}

	switch framework {
	case "Node.js":
		generateNodeProject(buildTool, packageManager, useTS)
	case "Nest.js":
		generateNestProject(packageManager, useTS)
	}
}

func generateNodeProject(buildTool, packageManager string, useTS bool) {
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

	// Install dependencies based on template and package manager
	installCmd := utils.GetPackageManagerInstallCmd(packageManager)
	switch buildTool {
	case "Express":
		utils.ExecuteCommand(installCmd[0], append(installCmd[1:], "express")...)
	case "Fastify":
		utils.ExecuteCommand(installCmd[0], append(installCmd[1:], "fastify")...)
	case "Koa":
		utils.ExecuteCommand(installCmd[0], append(installCmd[1:], "koa")...)
	}

	if useTS {
		utils.ExecuteCommand(installCmd[0], append(installCmd[1:], "typescript", "@types/node", "ts-node", "-D")...)
	}
}

func generateNestProject(packageManager string, useTS bool) {
	createCmd := []string{"create", "nest", "."}
	switch packageManager {
	case "npm":
		createCmd = append([]string{"npx", "@nestjs/cli"}, createCmd...)
	case "yarn":
		createCmd = append([]string{"yarn", "@nestjs/cli"}, createCmd...)
	case "pnpm":
		createCmd = append([]string{"pnpm", "@nestjs/cli"}, createCmd...)
	case "bun":
		createCmd = append([]string{"bunx", "@nestjs/cli"}, createCmd...)
	}

	if err := utils.ExecuteCommand(createCmd[0], createCmd[1:]...); err != nil {
		fmt.Printf("Error creating Nest.js project: %v\n", err)
		return
	}
}

// Add other backend-related functions... 