package action

import (
	"fmt"
	"os"
	"path/filepath"
	"photo-manager-cli/config"
	"photo-manager-cli/metadata"
	"regexp"
)

func UpdateDateFromMetadata(config config.Config) {
	fmt.Println("[INIT] UpdateDateFromMetadata")
	fmt.Println("path: ", config.Path)
	fmt.Println("regexp: ", config.Regexp)

	processMetadata(config.Path, config.Regexp)
	fmt.Println("[Finish] UpdateDateFromMetadata")
}

func processMetadata(basePath string, regexpStr string) {
	filepath.Walk(filepath.Join(basePath),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			process := true
			if regexpStr != "" {
				process, _ = regexp.MatchString(regexpStr, info.Name())
			}

			if process {
				metadata.UpdateDateFromMetadata(path, info)
			}
			return nil
		})
}
