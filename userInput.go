package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func userInput() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Which directory contains the files you want to rename? > ")
		if !scanner.Scan() {
			break
		}

		submittedRequest := scanner.Text()
		targetDirectory := ""

		if strings.ToLower(submittedRequest) == "exit" {
			fmt.Println("Exiting the converter")
			os.Exit(0)
		}

		homeDirectory, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error parsing home directory: %v", err)
			continue
		}

		if submittedRequest == "." {
			currentDir, err := os.Getwd()
			if err != nil {
				fmt.Printf("Error getting the current working directory: %s\n", err)
				continue
			}
			targetDirectory = currentDir
		} else {
			targetDirectory = filepath.Join(homeDirectory, submittedRequest)
		}

		fileStats, err := os.Stat(targetDirectory)
		if err != nil {
			fmt.Printf("Error getting stats for target dir: %v\n", err)
			continue
		}

		if fileStats.IsDir() {
			fmt.Printf("The target directory is: %s\n", targetDirectory)
			err := processFiles(targetDirectory)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("%s isn't a directory. Check the path again.", targetDirectory)
			continue
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error encountered: %v\n", err)
	}
}
