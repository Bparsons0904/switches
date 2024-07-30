package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func setStatic(appendNumber int64, server *fiber.App) {
	stylesDir := "./static/styles"
	scriptsDir := "./static/scripts"
	destDir := "./assets"

	if err := os.RemoveAll(destDir); err != nil {
		fmt.Printf("Error removing existing destination directory: %v\n", err)
		return
	}

	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating destination directory: %v\n", err)
		return
	}

	log.Println("Copying static files...")
	copyStaticFiles(stylesDir, destDir, appendNumber)
	copyStaticFiles(scriptsDir, destDir, appendNumber)

	server.Static("/assets", "./assets")
	server.Static("/images", "./static/assets/images")
}

func copyStaticFiles(sourceDir, destDir string, appendNumber int64) error {
	log.Printf("Reading directory %s\n", sourceDir)
	files, err := os.ReadDir(sourceDir)
	log.Println("files", files)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		oldPath := filepath.Join(sourceDir, file.Name())
		log.Println("oldPath", oldPath)

		ext := filepath.Ext(file.Name())
		baseName := strings.TrimSuffix(file.Name(), ext)
		newFileName := fmt.Sprintf("%s_%d.min%s", baseName, appendNumber, ext)
		newPath := filepath.Join(destDir, newFileName)

		if err := copyAndMinifyFile(oldPath, newPath); err != nil {
			fmt.Printf("Error copying and minifying file %s: %v\n", file.Name(), err)
			continue
		}

		fmt.Printf("Copied, minified, and renamed %s to %s\n", file.Name(), newFileName)
	}
	return nil
}

func copyAndMinifyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if strings.HasSuffix(strings.ToLower(src), ".css") {
		err = m.Minify("text/css", destFile, sourceFile)
		if err != nil {
			return fmt.Errorf("minification error: %v", err)
		}
	} else if strings.HasSuffix(strings.ToLower(src), ".js") {
		err = m.Minify("text/js", destFile, sourceFile)
		if err != nil {
			return fmt.Errorf("minification error: %v", err)
		}
	} else {
		_, err = io.Copy(destFile, sourceFile)
		if err != nil {
			return err
		}
	}

	return nil
}
