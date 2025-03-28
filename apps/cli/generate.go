package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "starter-kit",
	Short: "A CLI tool to generate framework starter kits",
	Long:  `A CLI tool that helps you generate starter kits for various frameworks like React, Next.js, and Vue.js.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type StarterKit struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Structure   []string `json:"structure"`
	Pros        []string `json:"pros"`
	Cons        []string `json:"cons"`
}

type StarterKitsData struct {
	StarterKits []StarterKit `json:"starterKits"`
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new starter kit",
	Run:   runGenerate,
}

func runGenerate(cmd *cobra.Command, args []string) {
	// First, ask what they want to build
	buildType := ""
	buildPrompt := &survey.Select{
		Message: "What do you want to build?",
		Options: []string{
			"Website",
			"Mobile App",
			"API",
			"Desktop App",
			"CLI Tool",
		},
	}
	survey.AskOne(buildPrompt, &buildType)

	// Then ask for development type
	developmentType := ""
	typePrompt := &survey.Select{
		Message: "What would you like to develop?",
		Options: []string{"Frontend", "Backend"},
	}
	survey.AskOne(typePrompt, &developmentType)

	// Define frameworks based on selection
	var frameworks []string
	if developmentType == "Frontend" {
		frameworks = []string{"React", "Vue.js", "Next.js"}
	} else {
		frameworks = []string{"Node.js", "Go", "Python/FastAPI"}
	}

	// Prompt for framework selection
	var selectedFramework string
	prompt := &survey.Select{
		Message: "Choose a framework:",
		Options: frameworks,
	}
	survey.AskOne(prompt, &selectedFramework)

	// Prompt for TypeScript
	useTS := false
	tsPrompt := &survey.Confirm{
		Message: "Would you like to use TypeScript?",
		Default: true,
	}
	survey.AskOne(tsPrompt, &useTS)

	// Create output directory
	outputDir := "output"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		return
	}

	fmt.Printf("\nGenerating %s starter kit with TypeScript: %v\n", selectedFramework, useTS)
	fmt.Printf("Your starter kit will be created in the '%s' directory\n", outputDir)
}

func readStarterKitsData() (*StarterKitsData, error) {
	// Read from the shared package
	filePath := filepath.Join("..", "..", "packages", "shared", "starter-kits.json")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var starterKits StarterKitsData
	if err := json.Unmarshal(data, &starterKits); err != nil {
		return nil, err
	}

	return &starterKits, nil
}
