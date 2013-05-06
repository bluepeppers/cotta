package data

import (
	"log"
	"os"
	"path/filepath"
)

func GetDataDir() string {
	exeName := os.Args[0]
	abs, _ := filepath.Abs(exeName)
	exeDir := filepath.Dir(abs)
	start := []string{exeDir}
	paths := append(start, dataPaths...)
	for _, pth := range paths {
		_, err := os.Open(filepath.Clean(pth + "/COTTADATA"))
		if !os.IsNotExist(err) {
			return pth
		}
	}
	log.Printf("Could not find data directory, defaulting to current directory")
	log.Printf("DataDir=%q", exeDir)
	return exeDir
}
