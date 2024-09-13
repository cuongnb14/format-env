package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

// readConfigFile reads the key-value pairs from a .env file
func readConfigFile(filePath string) (map[string]string, error) {
	config := make(map[string]string)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		config[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}

// formatEnv generates a formatted .env file using a template
func formatEnv(templatePath string, stage string) error {
	log.Printf("format env for stage %s using template %s", stage, templatePath)

	configPath := fmt.Sprintf("%s/%s.env", templatePath, stage)
	outputPath := configPath

	templateContent, err := os.ReadFile(templatePath + "/template.env")
	if err != nil {
		return fmt.Errorf("failed to read template file: %v", err)
	}

	config, err := readConfigFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	tmpl, err := template.New("env").Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outputFile.Close()

	err = tmpl.Execute(outputFile, config)
	if err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	return nil
}

// FormatEnv handles multiple stages and template path
func FormatEnv(templatePath string, stage string) error {
	stages := []string{"dev", "testing", "unstable", "staging"}
	if stage == "all" {
		for _, s := range stages {
			err := formatEnv(templatePath, s)
			if err != nil {
				log.Printf("failed to format env for stage %s: %v", s, err)
			}
		}
	} else {
		err := formatEnv(templatePath, stage)
		if err != nil {
			log.Printf("failed to format env for stage %s: %v", stage, err)
		}
	}
	return nil
}

func main() {
	if len(os.Args) < 3 {
		log.Println("Usage: format_env <template_path> <stage>")
		return
	}

	templatePath := os.Args[1]
	stage := os.Args[2]

	err := FormatEnv(templatePath, stage)
	if err != nil {
		log.Fatalf("Error formatting env: %v", err)
	}
}
