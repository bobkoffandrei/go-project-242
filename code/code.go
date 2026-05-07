package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	var result int64
	var resultStr string
	var err error
	dr, err := isDirCheck(path)
		if err != nil {
		return "", fmt.Errorf("ошибка проверки пути к файлу или папке: %w", err)
	}

	switch dr {
	case false:
		result, err = getFileSize(path)
		resultStr = humanSize(result, human)
	case true:
		result, err = getDirSize(path, all, recursive)
		resultStr = humanSize(result, human)
	}

	return resultStr, err
}

func humanSize(s int64, b bool) string {

	hs := ""

	switch {
	case s < 1000 || !b:
		hs = fmt.Sprintf("%.0fB", float64(s))
	case s > 999 && s < 1000000:
		hs = fmt.Sprintf("%.2fKB", float64(s)/1024)
	case s > 999999 && s < 1000000000:
		hs = fmt.Sprintf("%.2fMB", float64(s)/(1024*1024))
	case s > 999999999 && s < 1000000000000:
		hs = fmt.Sprintf("%.2fGB", float64(s)/(1024*1024*1024))
	case s > 999999999999 && s < 1000000000000000:
		hs = fmt.Sprintf("%.2fTB", float64(s)/(1024*1024*1024*1024))
	case s > 999999999999999 && s < 1000000000000000000:
		hs = fmt.Sprintf("%.2fPB", float64(s)/(1024*1024*1024*1024*1024))
	case s > 999999999999999999:
		hs = fmt.Sprintf("%.2fEB", float64(s)/(1024*1024*1024*1024*1024*1024))
	}
	return hs
}

// проверка на папку
func isDirCheck(dir string) (bool, error) {
	fileInfo, err := os.Stat(dir)
	if err != nil {
		return false, fmt.Errorf("ошибка на этапе проверки директории: %w", err)
	}
	return fileInfo.IsDir(), nil
}

// подсчет файла
func getFileSize(dir string) (int64, error) {
	info, err := os.Lstat(dir)
	if err != nil {
		return 0, fmt.Errorf("ошибка подсчета размера файла: %w", err)
	}
	return info.Size(), nil
}

// подсчет директории
func getDirSize(dir string, hidden bool, rsv bool) (int64, error) {
	var sumSize int64 = 0
	files, err := os.ReadDir(dir)
	if err != nil {
		return 0, fmt.Errorf("ошибка подсчета размера папки: %w", err)
	}
	for _, file := range files {

		tempDir := filepath.Join(dir, file.Name())

		fileInfo, _ := os.Stat(tempDir)

		if !fileInfo.IsDir() {
			info, err := os.Lstat(tempDir)
			if err != nil {
				return 0, err
			}
			if hidden {
				sumSize += info.Size()
			}
			if !hidden && !isHiddenFile(file.Name()) {
				sumSize += info.Size()
			}

		} else {
			//Исправил баг "при -r без -a рекурсия идёт в скрытые директории" 
			if rsv && hidden{
				dirSize, err := getDirSize(tempDir, hidden, rsv)
				if err != nil {
				return 0, err
			}
				sumSize += dirSize
			}

			if rsv && !hidden{
				if !isHiddenFile(tempDir){
				dirSize, err := getDirSize(tempDir, hidden, rsv)
				if err != nil {
				return 0, err
			}
				sumSize += dirSize
				}
			}

		}
		_, err = os.Lstat(tempDir)
		if err != nil {
			return 0, fmt.Errorf("ошибка подсчета размера папки: %w", err)
		}
	}

	return sumSize, nil
}

func isHiddenFile(s string) bool {
	return strings.HasPrefix(s, ".")
}
