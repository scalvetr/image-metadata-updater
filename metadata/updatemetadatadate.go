package metadata

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func UpdateMetadataDate(path string, info os.FileInfo, fileDateTime *time.Time, override bool) {
	if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".jpg") ||
		strings.HasSuffix(strings.ToLower(path), ".jpeg") ||
		strings.HasSuffix(strings.ToLower(path), ".gif") {
		updateMetadataDateJpg(path, fileDateTime, override)
	}
}

func updateMetadataDateJpg(filepath string, fileDateTime *time.Time, override bool) {
	fmt.Println("    - file: ", filepath)
	var existingFileDateTime = extractExifMetadataDate(filepath)
	if existingFileDateTime != nil {
		fmt.Println("      existingDateTime: ", existingFileDateTime)
	}
	if existingFileDateTime == nil || override {
		fmt.Println("      set - newDateTime: ", fileDateTime)
		setExifMetadataDate(filepath, *fileDateTime)
	} else {
		fmt.Println("      keep - existingDateTime")
	}
}
