package pathsize

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	var result string
	var err error

	switch isDirCheck(path) {
	case false:
		result, err = getFileSize(path)
		result = humanSize(result, human)
	case true:
		result, err = getDirSize(path, all, recursive)
		result = humanSize(result, human)
	}
	return result, err
}

func humanSize(s string, b bool) string {
	size, _ := strconv.Atoi(s)
	lenSize := len(s)
	hs := s

	switch {
	case lenSize < 4 || !b:
		hs = fmt.Sprintf("%.0fB", float64(size))

	case lenSize > 3 && lenSize < 7:
		hs = fmt.Sprintf("%.2fKB", float64(size)/1024)

	case lenSize > 6 && lenSize < 10:
		hs = fmt.Sprintf("%.2fMB", float64(size)/(1024*1024))

	case lenSize > 9 && lenSize < 13:
		hs = fmt.Sprintf("%.2fGB", float64(size)/(1024*1024*1024))

	case lenSize > 12 && lenSize < 17:
		hs = fmt.Sprintf("%.2fTB", float64(size)/(1024*1024*1024*1024))

	case lenSize > 16 && lenSize < 21:
		hs = fmt.Sprintf("%.2fPB", float64(size)/(1024*1024*1024*1024*1024*1024))

	case lenSize > 20 && lenSize < 25:
		hs = fmt.Sprintf("%.2fEB", float64(size)/(1024*1024*1024*1024*1024*1024*1024))

	}
	return hs
}

func isDirCheck(dir string) bool {
	fileInfo, err := os.Stat(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return fileInfo.IsDir()
}

func getFileSize(dir string) (string, error) {
	info, err := os.Lstat(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	r := fmt.Sprintf("%d", info.Size())
	return r, nil
}

func getDirSize(dir string, hidden bool, rsv bool) (string, error) {
	var sumSize int64 = 0
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	for _, file := range files {

		tempDir := dir + "/" + file.Name()

		fileInfo2, _ := os.Stat(tempDir)

		if !fileInfo2.IsDir() {
			info, err := os.Lstat(tempDir)
			if err != nil {
				fmt.Println("Error:", err)
				return "", err
			}
			if hidden {
				sumSize += info.Size()
			}
			if !hidden && !isHiddenFile(file.Name()) {
				sumSize += info.Size()
			}

		}else{

			if(rsv){
			dirSize, _ := getDirSize(tempDir, hidden, rsv)

			dirRes, _ := strconv.Atoi(dirSize)

			sumSize += int64(dirRes)
			}


		}
		_, err = os.Lstat(tempDir)
		if err != nil {
			fmt.Println("Error:", err)
			return "", err
		}
	}

	r := fmt.Sprintf("%d", sumSize)

	return r, nil
}

func isHiddenFile(s string) bool {
	return strings.HasPrefix(s, ".")
}
