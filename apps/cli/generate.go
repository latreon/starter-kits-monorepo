package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

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

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new starter kit",
	Run:   runGenerate,
}

func runGenerate(cmd *cobra.Command, args []string) {
	// Read starter kits data
	data, err := readStarterKitsData()
	if err != nil {
		fmt.Printf("Error reading starter kits data: %v\n", err)
		return
	}

	// Prepare framework choices
	var frameworks []string
	for _, kit := range data.StarterKits {
		frameworks = append(frameworks, kit.Name)
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
		Message: "Use TypeScript?",
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