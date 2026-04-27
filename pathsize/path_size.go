package pathsize

import (
	"os"
)


func GetPathSize(path string, recursive, human, all bool) (string, error) {
	info, err := os.Lstat("string")
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return info, nil
}