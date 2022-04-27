package metadata

import (
	"fmt"
	"os"
	"strings"
)

func UpdateDateFromMetadata(path string, info os.FileInfo) {
	if strings.HasSuffix(strings.ToLower(path), ".jpg") ||
		strings.HasSuffix(strings.ToLower(path), ".jpeg") ||
		strings.HasSuffix(strings.ToLower(path), ".gif") {
		updateDateFromMetadataJpg(path, info)
	} else if strings.HasSuffix(strings.ToLower(path), ".mpg") ||
		strings.HasSuffix(strings.ToLower(path), ".mpeg") {
		updateDateFromMetadataMpg(path, info)
	}
}

func updateDateFromMetadataMpg(path string, info2 os.FileInfo) {

}
func updateDateFromMetadataJpg(filepath string, info os.FileInfo) {
	var fileDateTime = extractExifMetadataDate(filepath, info)
	if fileDateTime != nil {
		fmt.Println("DateTime", fileDateTime)
		err := os.Chtimes(filepath, *fileDateTime, *fileDateTime)
		if err != nil {
			fmt.Println(err)
		}
	}
}
