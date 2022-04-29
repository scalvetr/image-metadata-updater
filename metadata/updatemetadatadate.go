package metadata

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func UpdateMetadataDate(path string, info os.FileInfo, fileDateTime *time.Time, override bool, replaces map[string]string) {
	if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".jpg") ||
		strings.HasSuffix(strings.ToLower(path), ".jpeg") ||
		strings.HasSuffix(strings.ToLower(path), ".gif") {
		updateMetadataDateJpg(path, fileDateTime, override, replaces)
	} else if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".mpg") ||
		strings.HasSuffix(strings.ToLower(path), ".mpeg") {
		updateMetadataDateMpeg(path, fileDateTime, override, replaces)
	}
}
func updateMetadataDateMpeg(filepath string, fileDateTime *time.Time, override bool, replaces map[string]string) {

}

func updateMetadataDateJpg(filepath string, fileDateTime *time.Time, override bool, replaces map[string]string) {
	fmt.Println("    - file: ", filepath)
	var existingFileDateTime = extractExifMetadataDate(filepath)
	if existingFileDateTime != nil {
		fmt.Println("      existingDateTime: ", existingFileDateTime)
		fileDateTime, override = processReplaces(existingFileDateTime, replaces)
	}
	if existingFileDateTime == nil || override {
		fmt.Println("      set - newDateTime: ", fileDateTime)
		setExifMetadataDate(filepath, *fileDateTime)
	} else {
		fmt.Println("      keep - existingDateTime")
	}
}

func processReplaces(dateTime *time.Time, replaces map[string]string) (*time.Time, bool) {
	layout := "2006-01-02T15:04:05Z07:00"

	key := dateTime.Format("2006-01-02")
	if val, ok := replaces[key]; ok {
		orig := dateTime.Format(layout)
		newVal := val + orig[10:]

		k, err := time.Parse(layout, newVal)
		if err != nil {
			log.Fatal(err)
		}
		return &k, true
	}
	return dateTime, false
}
