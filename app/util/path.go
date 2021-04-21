package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// 检测并创建文件夹
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
		return false, nil
	}
	return false, err
}

func CreateFile(path string) error {
	fp, err := os.Create(path) // 如果文件已存在，会将文件清空。
	if err != nil {
		return err
	}
	defer fp.Close() //关闭文件，释放资源。
	return nil
}

// 获取文件md5,两种方式 v1比v2快一倍左右
func GetMd5V1(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", md5hash.Sum(nil)), nil
}

func GetMd5V2(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	body, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", md5.Sum(body)), nil
}

// 获取前端模板
func GetTemplateFiles(dir string) ([]string, error) {
	var data []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() == false && strings.HasSuffix(info.Name(), ".html") {
			data = append(data, path)
		}
		return err
	})
	if err != nil {
		return data, err
	}
	return data, nil
}
