package nfile

import (
	"fmt"
	"io"
	"os"
)

// FileExists 检测文件是否存在
func FileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	if info.IsDir() {
		return false, fmt.Errorf("the supplied path %s is a dir", path)
	}
	return true, nil
}

// DirExists 检测目录是否存在
func DirExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	if !info.IsDir() {
		return false, fmt.Errorf("the supplied path %s exists but is not a dir", path)
	}
	return true, nil
}

// DirEmpty 检测目录是否为空
func DirEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.ReadDir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

// CreateDirIfNotExist 如果目录不存在则创建
func CreateDirIfNotExist(path string) (bool, error) {
	if exists, _ := DirExists(path); exists {
		return true, nil
	}

	if err := os.MkdirAll(path, 0755); err != nil {
		return false, err
	}

	return DirExists(path)
}
