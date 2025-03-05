package utils

import (
	"os"
	"os/exec"
)

func ExecuteCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func GetPackageManagerInstallCmd(pm string) []string {
	switch pm {
	case "yarn":
		return []string{"yarn", "add"}
	case "pnpm":
		return []string{"pnpm", "add"}
	case "bun":
		return []string{"bun", "add"}
	default:
		return []string{"npm", "install"}
	}
} 