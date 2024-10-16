package filesystem

import "os"

func CreateFolder(path string) error {

	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
