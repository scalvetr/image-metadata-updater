package action

import (
	"fmt"
	"os"
	"path/filepath"
	"photo-manager-cli/config"
	"photo-manager-cli/metadata"
)

func UpdateDateFromMetadata(config config.Config) {
	fmt.Println("[INIT] UpdateDateFromMetadata")
	fmt.Println("path: ", config.Path)

	processMetadata(config.Path)
	fmt.Println("[Finish] UpdateDateFromMetadata")
}

func processMetadata(basePath string) {
	filepath.Walk(filepath.Join(basePath),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			metadata.UpdateDateFromMetadata(path, info)
			return nil
		})
}
