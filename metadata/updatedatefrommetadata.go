package metadata

import (
	"fmt"
	"os"
	"strings"
)

func UpdateDateFromMetadata(path string, info os.FileInfo) {
	if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".jpg") ||
		strings.HasSuffix(strings.ToLower(path), ".jpeg") ||
		strings.HasSuffix(strings.ToLower(path), ".gif") {
		updateDateFromMetadataJpg(path)
	} else if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".mpg") ||
		strings.HasSuffix(strings.ToLower(path), ".mpeg") {
		updateDateFromMetadataMpg(path)
	}
}

func updateDateFromMetadataMpg(filepath string) {
}
func updateDateFromMetadataJpg(filepath string) {
	fmt.Println("    - file: ", filepath)
	var fileDateTime = extractExifMetadataDate(filepath)
	if fileDateTime != nil {
		fmt.Println("      dateTime: ", fileDateTime)
		err := os.Chtimes(filepath, *fileDateTime, *fileDateTime)
		if err != nil {
			fmt.Println(err)
		}
	}
}
