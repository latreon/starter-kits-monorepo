package utils

import (
	"fmt"
	"os"
)

func HandleInterrupt() {
	fmt.Printf("\n\n")
	fmt.Println("👋 Process terminated by user")
	fmt.Println("🧹 Cleaning up...")
	if _, err := os.Stat("output"); err == nil {
		os.RemoveAll("output")
		fmt.Println("✨ Cleanup completed")
	}
	os.Exit(0)
} 