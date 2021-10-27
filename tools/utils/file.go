package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	ErrFileName = errors.New("err file name")
)

func WriteFile(filename string, data []byte, perm os.FileMode) error {
	// 判断文件为目录
	if filename[len(filename)-1] == os.PathSeparator {
		return ErrFileName
	}
	path := GetPath(filename)
	if _, err := os.Stat(path); err != nil {
		os.MkdirAll(path, os.ModeDir)
	}
	return ioutil.WriteFile(filename, data, perm)
}

func ReadFile(filename string) ([]byte, error) {
	//
	_, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	return os.ReadFile(filename)
}

func GetPath(filename string) string {
	return filepath.Dir(filename)
}

// 获取可执行文件目录路径
func GetCurrfilePath() string {
	return filepath.Dir(os.Args[0])
}

// 获取上级目录路径
func GetParentfilePath(path string) string {
	return filepath.Dir(path)
}
