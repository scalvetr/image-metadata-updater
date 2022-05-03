package action

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	a "photo-manager-cli/album"
	c "photo-manager-cli/config"
	"photo-manager-cli/metadata"
	"time"
)

func CheckAlbumDateMismatch(config c.Config) {
	fmt.Println("[INIT] CheckAlbumDateMismatch")
	fmt.Println("path: ", config.Path)
	fmt.Println("albumNamePattern: ", config.AlbumInfoConfig.GetAlbumNamePattern())
	fmt.Println("getFolderRegexp: ", config.AlbumInfoConfig.GetFolderRegexp())

	files, err := ioutil.ReadDir(config.Path)
	if err != nil {
		log.Fatal(err)
	}

	var directories []fs.FileInfo
	for _, file := range files {
		if file.IsDir() {
			directories = append(directories, file)
		}
	}

	var w *bufio.Writer
	if config.CheckAlbumDateMismatchConfig.ReportFile != "" {
		var f *os.File
		f, err = os.Create(config.CheckAlbumDateMismatchConfig.ReportFile)
		if err != nil {
			log.Fatal(err)
		}
		w = bufio.NewWriter(f)
	}

	for _, directory := range directories {
		var info = a.ExtractAlbumInfo(directory, config.AlbumInfoConfig)
		fmt.Println("[CheckAlbumDateMismatch] - ", info.GetName(config.AlbumInfoConfig))
		fixDate(config.Path, directory, info, w)
	}
	fmt.Println("[Finish] CheckAlbumDateMismatch")
}

func fixDate(basePath string, directory fs.FileInfo, albumInfo a.AlbumInfo, w *bufio.Writer) {
	if w != nil {
		w.WriteString(fmt.Sprintf("- name: %s\n  year: %04d\n  month: %02d\n  mismatches:\n", albumInfo.Name, albumInfo.Year, albumInfo.Month))
	}
	filepath.Walk(filepath.Join(basePath, directory.Name()),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				metadata.FixDate(path, albumInfo.Year, albumInfo.Month,
					func(filePath string, fileDate *time.Time) {
						if w != nil {
							w.WriteString(fmt.Sprintf("  - file: %s\n    date: %v\n", filePath, fileDate))
						}
					})
			}
			return nil
		})
}
