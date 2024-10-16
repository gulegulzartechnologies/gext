package filesystem

import (
	"fmt"
	"os"
)

func CheckFolder(path string) error {

	folder, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("folder does not exist: %v", err)
		}
		return fmt.Errorf("error opening folder: %v", err)
	}
	defer folder.Close()

	return nil

}
