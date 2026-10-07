package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/forPelevin/gomoji"
)

func processFiles(targetDir string) error {
	acceptedTypes := []string{".md", ".txt", ".odt", ".doc", ".docx", ".rtf"}
	dirTree, err := os.ReadDir(targetDir)
	if err != nil {
		return fmt.Errorf("Error reading directory: %w", err)
	}

	for _, entry := range dirTree {

		if entry.IsDir() {
			new_target := filepath.Join(targetDir, entry.Name())
			fmt.Printf("%s is a directory; going recursive\n", entry.Name())
			processFiles(new_target)
			continue
		}

		fullFilePath := filepath.Join(targetDir, entry.Name())
		fileType := filepath.Ext(fullFilePath)

		if !slices.Contains(acceptedTypes, fileType) {
			fmt.Printf("%s isn't a text file; skipped\n", entry.Name())
			continue
		}

		firstLine, err := getFirstLine(fullFilePath)
		if err != nil {
			return err
		}

		newTitle := cleanFirstLine(firstLine) + fileType

		if err := renameCopy(fullFilePath, newTitle); err != nil {
			return err
		}

	}
	return nil
}

func renameCopy(fullFilePath, newTitle string) error {

	copyFileDir := filepath.Join(filepath.Dir(fullFilePath), "renamed_copies")

	_, err := os.Stat(copyFileDir)
	if errors.Is(err, fs.ErrNotExist) {
		if err := os.MkdirAll(copyFileDir, 0755); err != nil {
			return fmt.Errorf("Error creating new directory: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("Error checking directory's existence: %w", err)
	}

	newPath := filepath.Join(copyFileDir, newTitle)

	sourceData, err := os.ReadFile(fullFilePath)
	if err != nil {
		return fmt.Errorf("File cannot be read for copying: %w", err)
	}

	fmt.Printf("Copying '%s' to '%s'\n", fullFilePath, newTitle)
	if err := os.WriteFile(newPath, sourceData, 0666); err != nil {
		return fmt.Errorf("Error writing new file: %w", err)
	}
	fmt.Println("Success!")
	fmt.Println("")

	return nil
}

func getFirstLine(fullFilePath string) (string, error) {
	fileContents, err := os.Open(fullFilePath) // this does not read the file into memory

	if err != nil {
		return "", fmt.Errorf("File cannot be opened: %w", err)
	}

	r := bufio.NewReader(fileContents)
	firstLine, err := r.ReadString('\n') // only reading up to the first newline character
	if err != nil {
		return "", fmt.Errorf("Error reading first line of file: %w", err)
	}

	if err := fileContents.Close(); err != nil {
		return "", fmt.Errorf("Error closing file: %w", err)
	}

	return firstLine, nil
}

func cleanFirstLine(firstLine string) string {

	replacer := strings.NewReplacer("/", "_", "^", "_", ",", "_", "+", "plus", "!", "", "...", "_",
		".", "dot_", "'", "", "**", "double_star_", "*", "star_", "(", "", ")", "",
		"[", "_", "]", "_", ":", "_", " ", "_", "…", "")
	removeSymbols := replacer.Replace(firstLine)

	if gomoji.ContainsEmoji(removeSymbols) {
		removeSymbols = gomoji.RemoveEmojis(removeSymbols)
	}

	cleanFileName := strings.TrimSpace(strings.Trim(removeSymbols, "#!_"))

	return cleanFileName

}
