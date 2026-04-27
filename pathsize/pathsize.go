package pathsize

import (
	"os"
	"fmt"
//	"strconv"
)


func GetPathSize(path string, recursive, human, all bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		fmt.Println("Error:", err)
		
		return "", err

	}
	fmt.Println("Размер:", info.Size())
	//result := strconv.FormatInt(info.Size, 10) 
	return "", nil
}