package fileDir

import (
	"os"
	"path/filepath"
)

func ExecuteDirectory() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}
