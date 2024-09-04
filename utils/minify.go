package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/js"
)

var m *minify.M

func init() {
	m = minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/js", js.Minify)
}

func SetStatic(appendNumber int64, server *fiber.App) {
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

	// log.Println("Copying static files...")
	err := copyStaticFiles(stylesDir, destDir, appendNumber)
	if err != nil {
		log.Fatal().Err(err).Msg("Error copying static style files")
	}

	err = copyStaticFiles(scriptsDir, destDir, appendNumber)
	if err != nil {
		log.Fatal().Err(err).Msg("Error copying static script files")
	}

	server.Static("/assets", "./assets")
	server.Static("/images", "./static/assets/images")
}

func copyStaticFiles(sourceDir, destDir string, appendNumber int64) error {
	files, err := os.ReadDir(sourceDir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		oldPath := filepath.Join(sourceDir, file.Name())

		ext := filepath.Ext(file.Name())
		baseName := strings.TrimSuffix(file.Name(), ext)
		newFileName := fmt.Sprintf("%s_%d.min%s", baseName, appendNumber, ext)
		newPath := filepath.Join(destDir, newFileName)

		if err := copyAndMinifyFile(oldPath, newPath); err != nil {
			log.Error().
				Err(err).
				Str("filename", file.Name()).
				Msg("Error copying and minifying file")
			continue
		}
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
