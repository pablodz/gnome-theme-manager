package utils

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func ExtractZip(pathZip string, extractPath string) error {
	// Open the zip file
	r, err := zip.OpenReader(pathZip)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Find the only folder in the zip file
	var folderToExtract string
	for _, f := range r.File {
		if f.FileInfo().IsDir() {
			folderToExtract = f.Name
			break
		}
	}

	for _, f := range r.File {
		if filepath.Dir(f.Name) == folderToExtract {
			if f.FileInfo().IsDir() {
				os.MkdirAll(filepath.Join(extractPath, f.Name), f.Mode())
			} else {
				os.MkdirAll(filepath.Join(extractPath, filepath.Dir(f.Name)), f.Mode())
				rc, err := f.Open()
				if err != nil {
					log.Fatal(err)
					return err
				}
				defer rc.Close()

				outFile := filepath.Join(extractPath, f.Name)
				log.Println("Extracting: " + outFile)
				f, err := os.OpenFile(outFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
				if err != nil {
					log.Fatal(err)
					return err
				}
				defer f.Close()

				_, err = io.Copy(f, rc)
				if err != nil {
					log.Fatal(err)
					return err
				}
			}
		}
	}
	return nil
}

func ContainsShortSHA(input string) bool {
	// regular expression to match short SHA hashes with exactly 7 characters
	shortSHA := regexp.MustCompile(`\-\b[a-z0-9]{7}\b`)

	// find all short SHA hashes in the input string
	matches := shortSHA.FindAllString(input, -1)

	// return true if any matches were found, false otherwise
	return len(matches) > 0
}
