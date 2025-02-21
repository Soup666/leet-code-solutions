package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go <project_name>")
		os.Exit(1)
	}

	projectName := os.Args[1]
	projectDir := filepath.Join("./", projectName)

	// Create project directory
	if err := os.MkdirAll(projectDir, 0755); err != nil {
		fmt.Println("Error creating project directory:", err)
		os.Exit(1)
	}

	// Create main.go file
	mainGoContent := fmt.Sprintf(`package main

import "fmt"

func main() {
	fmt.Println("Hello, %s!")
}`, projectName)

	if err := os.WriteFile(filepath.Join(projectDir, "main.go"), []byte(mainGoContent), 0644); err != nil {
		fmt.Println("Error creating main.go:", err)
		os.Exit(1)
	}

	// Initialize go.mod file
	goModContent := fmt.Sprintf("module %s\n\n go 1.20", strings.ToLower(projectName))

	if err := os.WriteFile(filepath.Join(projectDir, "go.mod"), []byte(goModContent), 0644); err != nil {
		fmt.Println("Error creating go.mod:", err)
		os.Exit(1)
	}

	fmt.Println("Go project", projectName, "created successfully!")
}
