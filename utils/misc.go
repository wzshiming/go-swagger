package utils

import (
	"os"
	"path/filepath"
	"strings"
)

// FileExists reports whether the named file or directory exists.
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// 获得 全部能引用包的路径
func GetPaths() []string {
	gopath := []string{}
	gopath = append(gopath, "./", "./vendor/")

	for _, v := range []string{"../", "../../"} {
		gopath = append(gopath, v, filepath.Join(v, "src"))
	}

	for _, v := range strings.Split(os.Getenv("GOPATH"), ";") {
		gopath = append(gopath, filepath.Join(v, "src"))
	}

	gopath = append(gopath, filepath.Join(os.Getenv("GOROOT"), "src"))

	for i := 0; i != len(gopath); {
		gopath[i] = filepath.Clean(gopath[i])
		fi, err := os.Stat(gopath[i])
		if err != nil || !fi.IsDir() {
			gopath = append(gopath[:i], gopath[i+1:]...)
			continue
		}
		i++
	}
	return gopath
}
