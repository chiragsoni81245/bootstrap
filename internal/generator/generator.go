package generator

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"github.com/chiragsoni81245/bootstrap/internal/config"
)


func CreateNewProject(templateName string, config config.Config) error{
    templatePath := fmt.Sprintf("./skeletons/%s", templateName)
    project, ok := config["project"].(map[string]any)
    if !ok {
        return fmt.Errorf("project parameter is required in config")
    }
    destinationDir, ok := project["folder_name"].(string)
    if !ok {
        return fmt.Errorf("folder_name parameter is required in config under project section")
    }
    destinationDir = fmt.Sprintf("./%s", destinationDir)
    
    return filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing path %q: %w", path, err)
		}

		// Determine the relative path from the templatePath
		relPath, err := filepath.Rel(templatePath, path)
		if err != nil {
			return fmt.Errorf("failed to compute relative path: %w", err)
		}

		destPath := filepath.Join(destinationDir, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, os.ModePerm)
		}

        filePathParts := strings.Split(path, ".")
        fileExtension := filePathParts[len(filePathParts)-1]
        extentionForTemplating := []string{"go", "js", "css", "html", "md", "mod", "sum", "txt", "yaml", "json"}
        filesForTemplating := []string{"Makefile"}
        
        if slices.Contains(extentionForTemplating, strings.ToLower(fileExtension)) || slices.Contains(filesForTemplating, info.Name()) {
            // Read file content
            inputBytes, err := os.ReadFile(path)
            if err != nil {
                return fmt.Errorf("failed to read file %q: %w", path, err)
            }

            // Parse as Go template
            tmpl, err := template.New(relPath).Parse(string(inputBytes))
            if err != nil {
                return fmt.Errorf("error parsing template %q: %w", path, err)
            }

            // Create destination file
            outFile, err := os.Create(destPath)
            if err != nil {
                return fmt.Errorf("failed to create output file %q: %w", destPath, err)
            }
            defer outFile.Close()

            // Execute template and write to destination file
            err = tmpl.Execute(outFile, config)
            if err != nil {
                return fmt.Errorf("error executing template for file %q: %w", path, err)
            }
        } else {
            sourceFile, err := os.Open(path)
            if err != nil {
                return err
            }
            defer sourceFile.Close()

            // Create the destination file
            destFile, err := os.Create(destPath)
            if err != nil {
                return err
            }
            defer destFile.Close()

            // Copy the contents
            _, err = io.Copy(destFile, sourceFile)
            if err != nil {
                return err
            }

            // Optionally copy file permissions
            info, err := os.Stat(path)
            if err != nil {
                return err
            }
            return os.Chmod(destPath, info.Mode())
        }

		return nil
	})
}
