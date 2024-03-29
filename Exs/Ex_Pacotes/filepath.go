package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk("./src/ProblemaNumérico", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path, info.Name(), info.Size(), info.IsDir())
		return nil
	})
}
