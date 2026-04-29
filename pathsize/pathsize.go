package pathsize

import (
	"os"
	"fmt"
//"strconv"
)


func GetPathSize(path string, recursive, human, all bool) (string, error) {
	var result string
	var err error

switch isDirCheck(path){
	case false:
	result, err = getFileSize(path)
	case true:
	result, err = getDirSize(path)	
}
	return result, err
}

func isDirCheck(dir string) bool {
	fileInfo, err := os.Stat(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return fileInfo.IsDir()
}

func getFileSize(dir string) (string, error){
	info, err := os.Lstat(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	r := fmt.Sprintf("%d", info.Size())
	return r, nil
}

func getDirSize(dir string) (string, error){
var sumSize int64 = 0
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	for _, file := range files {

		tempDir := dir + "/" + file.Name()


	fileInfo2, err := os.Stat(tempDir)

    if (!fileInfo2.IsDir()) {
			info, err := os.Lstat(tempDir)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	sumSize += info.Size()
	}
	_, err = os.Lstat(tempDir)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	}
	//fmt.Printf("%dB\t%s\n", sumSize, dir)
	r := fmt.Sprintf("%d", sumSize)

	return r, nil
}

func printFileSize(dir string) {
	info, err := os.Lstat(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("%dB\t%s\n", info.Size(), dir)
}

func printDirSize(dir string) {
var sumSize int64 = 0
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return 
	}
	for _, file := range files {

		tempDir := dir + "/" + file.Name()


	fileInfo2, err := os.Stat(tempDir)

    if (!fileInfo2.IsDir()) {
			info, err := os.Lstat(tempDir)
	if err != nil {
		fmt.Println("Error:", err)
		return 
	}
	sumSize += info.Size()
	}
	_, err = os.Lstat(tempDir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	}
	fmt.Printf("%dB\t%s\n", sumSize, dir)
}