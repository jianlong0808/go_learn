package pathutils

import (
	"os"
	"path/filepath"
)

//ProjectAbsPath 获取可执行文件所在的绝对路径
func ProjectAbsPath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return dir, nil
}
