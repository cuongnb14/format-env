package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

func defaultFunc(value, fallback any) string {
	if value == nil {
		return fallback.(string)
	}
	if value == "" {
		return fallback.(string)
	}
	return value.(string)
}

// readConfigFile reads the key-value pairs from a .env file
func readConfigFile(filePath string) (map[string]string, error) {
	config := make(map[string]string)

	file, err := os.OpenFile(filePath, os.O_CREATE, 0644)
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
func formatEnv(envDir string, stage string) error {
	log.Printf("using template %s/_template.env format %s/%s.env", envDir, envDir, stage)

	configPath := fmt.Sprintf("%s/%s.env", envDir, stage)
	outputPath := configPath

	templateContent, err := os.ReadFile(envDir + "/_template.env")
	if err != nil {
		return fmt.Errorf("failed to read template file: %v", err)
	}

	config, err := readConfigFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	tmpl := template.Must(template.New("env").Option("missingkey=zero").Funcs(template.FuncMap{
		"df": defaultFunc,
	}).Parse(string(templateContent)))

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

func main() {
	if len(os.Args) < 3 {
		log.Println("Usage: fenv <env_dir> <stages>")
		return
	}

	envDir := os.Args[1]
	stageStr := os.Args[2]
	result := strings.Split(stageStr, ",")
	for _, v := range result {
		err := formatEnv(envDir, v)
		if err != nil {
			log.Fatalf("Error formatting env: %v", err)
		}
	}
}
